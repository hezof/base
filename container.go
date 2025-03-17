package core

import (
	"fmt"
	"sync"
)

// ComponentConfig 组件配置
type ComponentConfig struct {
	Key   string         // 配置中的引用key,用于Reload更新
	Names []string       // 支持逗号分隔的别名机制
	Value map[string]any // 配置值
}

// NativeComponent 具体组件
type NativeComponent = interface{}

// ManagedComponent 托管组件接口
type ManagedComponent interface {
	// Bind 托管实例绑定具体实例
	Bind(native NativeComponent)
}

// ComponentFactory 组件工厂接口
type ComponentFactory interface {
	// Manage 创建托管实例
	Manage() ManagedComponent
	// Native 创建具体实例
	Native(c ConfigToml) (NativeComponent, error)
	// Destroy 销毁具体实例
	Destroy(v NativeComponent) error
}

// ConfigNative 带配置的具体实例
type ConfigNative struct {
	Config ConfigToml      // 配置
	Native NativeComponent // 实例
}

// ManagedFactory 托管工厂
type ManagedFactory struct {
	sync.RWMutex                             // 读写锁
	Kind         string                      // 类型
	Factory      ComponentFactory            // 工厂
	Manage       map[string]ManagedComponent // 托管实例. 键为Resource(name)声明的name
	Native       []*ConfigNative             // 具体实例. 键为配置中[kind.name]定义的name
}

func (mf *ManagedFactory) GetOrNewManaged(name string) ManagedComponent {
	mf.RLock()
	rt, ok := mf.Manage[name]
	mf.RUnlock()
	if !ok {
		mf.Lock()
		// 二次检测
		if rt, ok = mf.Manage[name]; !ok {
			rt = mf.Factory.Manage()
			mf.Manage[name] = rt
		}
		mf.Unlock()
	}
	return rt
}

// ManagedContainer 托管容器
type ManagedContainer struct {
	sync.RWMutex                            // 读写锁
	Factory      map[string]*ManagedFactory // 注册工厂. 键是kind, 值是托管工厂
	Indexes      []string                   // 注册顺序, 决定Init()/Exit()的组件顺序
}

func (mc *ManagedContainer) MustRegisterFactory(kind string, factory ComponentFactory) {
	mc.Lock()
	defer mc.Unlock()

	ft, ok := _container.Factory[kind]
	if ok {
		panic(fmt.Errorf("factory existed %v(%T)", kind, ft))
	}
	mc.Indexes = append(_container.Indexes, kind)
	mc.Factory[kind] = &ManagedFactory{
		Factory: factory,
		Manage:  make(map[string]ManagedComponent),
		Native:  make([]*ConfigNative, 0, 4),
	}
}

func (mc *ManagedContainer) MustGetFactory(kind string) *ManagedFactory {
	_container.RLock()
	defer _container.RUnlock()

	ft, ok := _container.Factory[kind]
	if !ok {
		panic(fmt.Errorf("factory unknown %v", kind))
	}
	return ft
}

func (mc *ManagedContainer) MustInit(root ...ConfigToml) {

}

func (mc *ManagedContainer) Reload(root ...ConfigToml) error {
	return nil
}

func (mc *ManagedContainer) Exit(hints ...func(kind, name string, err error)) {

}

// _container 内部托管容器
var _container = &ManagedContainer{
	Factory: make(map[string]*ManagedFactory),
	Indexes: make([]string, 0, 16),
}
