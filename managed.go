package base

import (
	"fmt"
	"github.com/hezof/log"
	"strings"
	"sync"
)

/*************************************************
 * 容器接口
 *************************************************/

// Register 注册组件工厂, 重复注册会panic!
func Register(base string, factory ManagedFactory) {
	err := _managedContext.RegisterFactory(base, factory)
	if err != nil {
		panic(fmt.Errorf("register factory %v error: %v", base, err))
	}
}

// Component 断言组件实例, 若无base工厂会panic!
func Component[T any](base string, name string) T {
	component, err := _managedContext.RetrieveComponent(base, name)
	if err != nil {
		panic(fmt.Errorf("retrieve component %v.%v error: %v", base, name, err))
	}
	return component.(T)
}

/*************************************************
 * 容器实现
 *************************************************/

// ManagedTarget 托管目标接口
type ManagedTarget = any

// ManagedComponent 托管组件接口
type ManagedComponent interface {
	GetTarget() ManagedTarget
	SetTarget(target ManagedTarget)
}

// ManagedConfig 组件配置. 所有配置格式: [base]={key:"a,",...}
type ManagedConfig struct {
	ID    string         // 配置中的引用key,用于Reload更新
	Names []string       // 支持逗号分隔的别名机制
	Value map[string]any // 配置值
}

// [约定] _id 规则.
func _id(base string, kvs map[string]any) (string, error) {
	rt := ""
	// 全局惟一的资源实例,可能不会设置_id
	if v, ok := kvs["_id"]; ok {
		if rt, ok = v.(string); !ok {
			return rt, fmt.Errorf("invalid type for config _id %v, expected string, but found %T", base, v)
		}
	}
	return rt, nil
}

// [约定] _flat 由于多个文件(glob) + 2种格式("[core]"或"[[core]]"), 需要将配置展开为平面分片.
func _flat(base string, configGroupSlice []any) ([]map[string]any, error) {
	// 打平(展开)配置组分片
	flat := make([]map[string]any, 0, len(configGroupSlice))
	for _, configGroup := range configGroupSlice {
		switch val := configGroup.(type) {
		case []any:
			for _, v := range val {
				config, ok := v.(map[string]any)
				if !ok {
					return nil, fmt.Errorf("invalid type for config %v, expected map[string]any(or its slice), but found %T", base, v)
				}
				flat = append(flat, config)
			}
		case map[string]any:
			flat = append(flat, val)
		default:
			return nil, fmt.Errorf("invalid type for config %v, expected map[string]any(or its slice), but found %T", base, configGroup)
		}
	}
	return flat, nil
}

// [约定] _names 配置_id支持别名机制(多个名字), 可以用逗号分隔
func _names(id string) []string {
	// 全局惟一的资源实例
	if id == "" {
		return []string{""}
	}
	n := 0
	for _, c := range id {
		if c == ',' {
			n++
		}
	}
	s := make([]string, 0, n)
	for {
		pos := strings.IndexByte(id, ',')
		if pos == -1 {
			if v := strings.TrimSpace(id); v != "" {
				s = append(s, v)
			}
			return s
		}
		if v := strings.TrimSpace(id[:pos]); v != "" {
			s = append(s, v)
		}
		id = id[pos+1:]
	}
}

// AssertManagedConfig 解析多个文件base路径的配置, 可能有2种形式: "[core]"或"[[core]]"
func AssertManagedConfig(base string, vals []any) ([]*ManagedConfig, error) {

	// 打平(展开)配置组分片. 多个文件(glob) + 2种格式("[base]"或"[[base]]")
	flat, err := _flat(base, vals)
	if err != nil {
		return nil, err
	}

	// 检测名字冲突并转换配置
	exists := make(map[string]bool)
	result := make([]*ManagedConfig, 0, len(flat))

	for _, kvs := range flat {
		id, err := _id(base, kvs)
		if err != nil {
			return nil, err
		}
		names := _names(id)
		for _, name := range names {
			if exists[name] {
				return nil, fmt.Errorf("existed name for config %v.%v", base, name)
			}
			exists[name] = true
		}
		result = append(result, &ManagedConfig{
			ID:    id,
			Names: names,
			Value: kvs,
		})
	}

	return result, nil
}

func AssertManagedConfigValues(base, dstId string, vals []any) (map[string]any, error) {
	// 打平(展开)配置组分片. 多个文件(glob) + 2种格式("[base]"或"[[base]]")
	flat, err := _flat(base, vals)
	if err != nil {
		return nil, err
	}
	for _, kvs := range flat {
		id, err := _id(base, kvs)
		if err != nil {
			return nil, err
		}
		if id == dstId {
			return kvs, nil
		}
	}
	return nil, nil
}

// ManagedFactory 组件工厂接口
type ManagedFactory interface {
	// Manage 创建托管组件
	Manage(name string) ManagedComponent
	// Create 创建托管目标
	Create(c *ManagedConfig) (ManagedTarget, error)
	// Destroy 销毁托管目标
	Destroy(v ManagedTarget) error
}

// ManagedConfigTarget 托管目标包裹器
type ManagedConfigTarget struct {
	Config *ManagedConfig // 配置
	Target ManagedTarget  // 目标
}

// ManagedFactoryProxy 托管工厂包裹器
type ManagedFactoryProxy struct {
	factory    ManagedFactory              // 工厂
	Components map[string]ManagedComponent // 托管组件. 键为Resource(name)声明的name
	Targets    []*ManagedConfigTarget      // 具体实例. 键为配置中[kind.name]定义的name
}

func (fw *ManagedFactoryProxy) Manage(name string) ManagedComponent {
	mc, ok := fw.Components[name]
	if !ok {
		mc = fw.factory.Manage(name)
		fw.Components[name] = mc
	}
	return mc
}

func (fw *ManagedFactoryProxy) Create(config *ManagedConfig) (ManagedTarget, error) {
	target, err := fw.factory.Create(config)
	if err != nil {
		return nil, err
	}
	fw.Targets = append(fw.Targets, &ManagedConfigTarget{
		Config: config,
		Target: target,
	})
	return target, nil
}

func (fw *ManagedFactoryProxy) Destroy(target ManagedTarget) error {
	return fw.factory.Destroy(target)
}

var _ ManagedFactory = (*ManagedFactoryProxy)(nil)

// ManagedContext 托管容器
type ManagedContext struct {
	sync.RWMutex                                 // 只在上下文设置读写锁
	Proxies      map[string]*ManagedFactoryProxy // 注册工厂. 键是kind, 值是托管工厂
	Indexes      []string                        // 注册顺序, 决定Init()/Exit()的组件顺序
}

func (mc *ManagedContext) RegisterFactory(base string, factory ManagedFactory) error {

	log.Info("register managed factory %v", base)

	mc.Lock()
	defer mc.Unlock()

	ft, ok := mc.Proxies[base]
	if ok {
		return fmt.Errorf("existed factory %v(%T)", base, ft)
	}
	mc.Indexes = append(mc.Indexes, base)
	mc.Proxies[base] = &ManagedFactoryProxy{
		factory:    factory,
		Components: make(map[string]ManagedComponent),
		Targets:    make([]*ManagedConfigTarget, 0, 4),
	}
	return nil
}

func (mc *ManagedContext) RetrieveComponent(base string, name string) (ManagedComponent, error) {
	mc.RLock()
	defer mc.RUnlock()
	// 返回代理工厂
	ft, ok := mc.Proxies[base]
	if !ok {
		return nil, fmt.Errorf("unknown factory %v", base)
	}
	return ft.Manage(name), nil
}

func (mc *ManagedContext) Init(configContext *ConfigContext) error {

	log.Info("init managed context...")

	mc.Lock()
	defer mc.Unlock()

	/*
			初始化流程:
			1.根据factory的注册顺序
		    2.获取kind的proxy
			3.查找kind的全部配置,并处理不符合约定或名字冲突的配置
			4.遍历全部配置,执行下述:
			- proxy.Create(config)创建托管目标
			- proxy.Manage(name)获取托管组件
			- ManageComponent.SetTarget(target)绑定目标
			5.遍历proxy已有的Component,ManagedComponent.GetTarget(), 若为空则报错.
	*/
	for _, base := range mc.Indexes {
		proxy := mc.Proxies[base]
		// 处理不符约定或名字冲突的配置
		configs, err := AssertManagedConfig(base, configContext.GetAll(base))
		if err != nil {
			return err
		}
		for _, config := range configs {

			log.Info("init managed target %v.%v", base, config.ID)

			target, err := proxy.Create(config)
			if err != nil {
				return err
			}
			for _, name := range config.Names {
				component := proxy.Manage(name)
				component.SetTarget(target)
			}
		}
		// 初始化后,所有托管组件都必须绑定目标. 否则无法使用!
		for name, component := range proxy.Components {
			if component.GetTarget() == nil {
				return fmt.Errorf("init managed component %v.%v failed: target empty", base, name)
			}
		}
	}
	return nil
}

func (mc *ManagedContext) Reload(configContext *ConfigContext, reloadPolicy func(base string, config *ManagedConfig, newValues map[string]any) bool) {

	log.Info("reload managed context...")

	mc.Lock()
	defer mc.Unlock()
	/*
		重载流程:
		1.根据factory的注册顺序
		2.获取kind的proxy
		3.遍历proxy的targets,执行下述:
		- 从configContext获取<kind>.<key>的配置, 根据policy比较新旧差异. 如果没有差异则跳过.
		- 用proxy.Create(config)创建新目标, 并遍历已有Components, 替换其目标
		- 用proxy.Destroy(config, target)销毁旧目标
	*/

	for _, base := range mc.Indexes {
		proxy := mc.Proxies[base]
	_TARGET_:
		for _, ct := range proxy.Targets {
			oldValues := ct.Config.Value
			oldTarget := ct.Target
			// 重载新配置
			newValues, err := AssertManagedConfigValues(base, ct.Config.ID, configContext.GetAll(base))
			if err != nil {
				log.Error("reload managed target %v.%v error: %v", base, ct.Config.ID, err)
				continue _TARGET_
			}
			if newValues != nil && reloadPolicy != nil && reloadPolicy(base, ct.Config, newValues) {

				log.Info("reload managed target %v.%v", base, ct.Config.ID)

				ct.Config.Value = newValues
				ct.Target, err = proxy.factory.Create(ct.Config)
				if err != nil {
					// 恢复旧配置
					ct.Config.Value = oldValues
					log.Error("reload managed target %v.%v error: %v", base, ct.Config.ID, err)
					continue _TARGET_
				}
				// 重置组件目标
				for _, name := range ct.Config.Names {
					proxy.Manage(name).SetTarget(ct.Target)
				}
				// 销毁旧目标,避免内存泄露
				err = proxy.Destroy(oldTarget)
				if err != nil {
					log.Error("destroy managed target %v.%v error: %v, please check memory leaks", base, ct.Config.ID, err)
				}
			}
		}
	}
}

func _path(base, _id string) string {
	// 全局惟一实例
	if _id == "" {
		return base
	}
	return base + "." + _id
}

func (mc *ManagedContext) Exit(hints ...func(base string, config *ManagedConfig, err error)) {

	log.Info("exit managed context...")

	mc.Lock()
	defer mc.Unlock()

	/*
		退出流程:
		1.根据factory的注册逆序. 注意, 必须是逆序!
		2.获取kind的proxy
		3.遍历proxy的targets,执行下述:
		- 用proxy.Destroy(config, target)销毁旧目标
	*/
	for i := len(mc.Indexes) - 1; i >= 0; i-- {
		base := mc.Indexes[i]
		proxy := mc.Proxies[base]
		for _, ct := range proxy.Targets {
			log.Info("destroy managed target %v.%v", base, ct.Config.ID)
			err := proxy.Destroy(ct.Target)
			if err != nil {
				log.Info("exec hint for %v.%v", base, ct.Config.ID)
				for _, hint := range hints {
					execHint(hint, base, ct.Config, err)
				}
			}
		}
	}
}

func execHint(hint func(base string, config *ManagedConfig, err error), base string, config *ManagedConfig, err error) {
	defer func() {
		if prr := recover(); prr != nil {
			log.Error("exec hint for %v.%v panic: %v|%v", base, config.ID, err, StackTrace(2, `|`))
		}
	}()
	hint(base, config, err)
}

// _managed 内部托管容器
var _managedContext = &ManagedContext{
	Proxies: make(map[string]*ManagedFactoryProxy),
	Indexes: make([]string, 0, 8),
}
