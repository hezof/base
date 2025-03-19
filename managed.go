package base

import (
	"fmt"
	"github.com/hezof/log"
	"strings"
	"sync"
)

const ID_NAME = "_id" // 同mongodb

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

// AssertManagedConfig 解析多个文件base路径的配置, 可能有2种形式: "[base]"或"[[base]]"
func AssertManagedConfig(base string, configGroupSlice []any) ([]*ManagedConfig, error) {
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
	// 检测名字冲突并转换配置
	exists := make(map[string]bool)
	result := make([]*ManagedConfig, 0, len(flat))

	for _, vals := range flat {
		var key string
		val, ok := vals[ID_NAME]
		if ok {
			if key, ok = val.(string); !ok {
				return nil, fmt.Errorf("invalid type for config _id %v, expected string, but found %T", base, val)
			}
		}
		names := strings.Split(key, ",")
		for _, name := range names {
			if exists[name] {
				return nil, fmt.Errorf("existed name for config %v.%v", base, name)
			}
			exists[name] = true
		}
		result = append(result, &ManagedConfig{
			ID:    key,
			Names: names,
			Value: vals,
		})
	}

	return result, nil
}

func AssertManagedConfigValues(vs any, ok bool) (map[string]any, bool) {
	if !ok {
		return nil, false
	}
	rt, ok := vs.(map[string]any)
	return rt, ok
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
		return fmt.Errorf("factory existed %v(%T)", base, ft)
	}
	mc.Indexes = append(mc.Indexes, base)
	mc.Proxies[base] = &ManagedFactoryProxy{
		factory:    factory,
		Components: make(map[string]ManagedComponent),
		Targets:    make([]*ManagedConfigTarget, 0, 4),
	}
	return nil
}

func (mc *ManagedContext) RetrieveFactory(base string) ManagedFactory {
	mc.RLock()
	defer mc.RUnlock()
	// 返回代理工厂
	return mc.Proxies[base]
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
				return fmt.Errorf("init managed component failed %v.%v, target empty", base, name)
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
	var err error
	for _, base := range mc.Indexes {
		proxy := mc.Proxies[base]
	_TARGET_:
		for _, ct := range proxy.Targets {
			oldValues := ct.Config.Value
			oldTarget := ct.Target
			// 重载新配置
			newValues, ok := AssertManagedConfigValues(configContext.GetFirst(base + "." + ct.Config.ID))
			if ok && reloadPolicy != nil && reloadPolicy(base, ct.Config, newValues) {

				log.Info("reload managed target %v.%v", base, ct.Config.ID)

				ct.Config.Value = newValues
				ct.Target, err = proxy.factory.Create(ct.Config)
				if err != nil {
					// 恢复旧配置
					ct.Config.Value = oldValues
					log.Error("reload managed target error %v.%v, %v", base, ct.Config.ID, err)
					continue _TARGET_
				}
				// 重置组件目标
				for _, name := range ct.Config.Names {
					proxy.Manage(name).SetTarget(ct.Target)
				}
				// 销毁旧目标,避免内存泄露
				err = proxy.Destroy(oldTarget)
				if err != nil {
					log.Error("destroy managed target error %v.%v, %v, please check memory leaks", base, ct.Config.ID, err)
				}
			}
		}
	}
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
			log.Error("exec hint panic: %v.%v, %v\n%", base, config.ID, err, StackTrace(2, `|`))
		}
	}()
	hint(base, config, err)
}

// _managed 内部托管容器
var _managedContext = &ManagedContext{
	Proxies: make(map[string]*ManagedFactoryProxy),
	Indexes: make([]string, 0, 8),
}
