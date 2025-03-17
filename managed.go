package core

import (
	"fmt"
	"sync"
)

// ManagedTarget 托管目标接口
type ManagedTarget = any

// ManagedComponent 托管组件接口
type ManagedComponent interface {
	Bind(target ManagedTarget)
}

// ManagedConfig 组件配置
type ManagedConfig struct {
	Key   string         // 配置中的引用key,用于Reload更新
	Names []string       // 支持逗号分隔的别名机制
	Value map[string]any // 配置值
}

// ManagedFactory 组件工厂接口
type ManagedFactory interface {
	// Manage 创建托管组件
	Manage() ManagedComponent
	// Create 创建托管目标
	Create(c ManagedConfig) (ManagedTarget, error)
	// Destroy 销毁托管目标
	Destroy(c ManagedConfig, v ManagedTarget) error
}

// TargetWrapper 托管目标包裹器
type TargetWrapper struct {
	Config ManagedConfig // 配置
	Target ManagedTarget // 目标
}

// FactoryWrapper 托管工厂包裹器
type FactoryWrapper struct {
	sync.RWMutex                             // 读写锁
	Factory      ManagedFactory              // 工厂
	Manage       map[string]ManagedComponent // 托管组件. 键为Resource(name)声明的name
	Target       []*TargetWrapper            // 具体实例. 键为配置中[kind.name]定义的name
}

// ManagedContext 托管容器
type ManagedContext struct {
	sync.RWMutex                            // 读写锁
	Factory      map[string]*FactoryWrapper // 注册工厂. 键是kind, 值是托管工厂
	Indexes      []string                   // 注册顺序, 决定Init()/Exit()的组件顺序
}

func (fc *ManagedContext) MustRegisterFactory(kind string, factory ManagedFactory) {
	fc.Lock()
	defer fc.Unlock()

	ft, ok := fc.Factory[kind]
	if ok {
		panic(fmt.Errorf("factory existed %v(%T)", kind, ft))
	}
	fc.Indexes = append(fc.Indexes, kind)
	fc.Factory[kind] = &FactoryWrapper{
		Factory: factory,
		Manage:  make(map[string]ManagedComponent),
		Target:  make([]*TargetWrapper, 0, 4),
	}
}

func (fc *ManagedContext) MustRetrieveFactory(kind string) *FactoryWrapper {
	fc.RLock()
	defer fc.RUnlock()

	ft, ok := fc.Factory[kind]
	if !ok {
		panic(fmt.Errorf("factory unknown %v", kind))
	}
	return ft
}

func (fw *FactoryWrapper) GetOrNewComponent(name string) ManagedComponent {
	fw.RLock()
	mc, ok := fw.Manage[name]
	fw.RUnlock()
	if !ok {
		fw.Lock()
		// 二次检测
		if mc, ok = fw.Manage[name]; !ok {
			mc = fw.Factory.Manage()
			fw.Manage[name] = mc
		}
		fw.Unlock()
	}
	return mc
}

func (mc *ManagedContainer) MustInit(root ...ConfigToml) {

}

func (mc *ManagedContainer) Reload(root ...ConfigToml) error {
	return nil
}

func (mc *ManagedContainer) Exit(hints ...func(kind, name string, err error)) {

}

// _managed 内部托管容器
var _managed = &ManagedContext{
	Factory: make(map[string]*FactoryWrapper),
	Indexes: make([]string, 0, 16),
}
