package core

import (
	"fmt"
	"github.com/hezof/log"
	"sync"
)

/********************************************
 * 组件容器API
 ********************************************/

type State = int8

const (
	StateNew     State = 0  //  已创建
	StateInit    State = 1  // 已初始
	StateExit    State = 2  // 已退出
	StateInvalid State = -1 // 不可用
)

// Config 组件配置(动态信息)
type Config = map[string]any

// Component 组件接口. 实现容器Init()/Exit()回调. 调用顺序取决于注册顺序!
type Component interface {
	// Init 容器初始回调接口, 负责组件资源的初始工作. 组件需实现该回调!
	Init(config Config) error
	// Reload 重新加载配置, 负责组件资源的热更工作. 组件支持热更需实现该回调!
	Reload(config Config) error
	// Exit 容器释放回调接口, 负责组件资源的释放工作. 组件需实现该回调!
	Exit() error
}

// Factory 工厂接口. 实现组件创建(但未初始化工作).
type Factory interface {
	// Create 创建组件(但不初始化). 返回名称(别名), 实例
	Create() ([]string, Component, error)
}

// Container 容器接口. 负责组件管理(初始化/重加载/销毁)
type Container interface {
	// Register 注册组件工厂
	Register(kind string, factory Factory) error
	// Load 加载组件
	Load(kind, name string) any
	// Init 容器初始配置
	Init(root Config) error
	// Reload 容器重加载配置
	Reload(root Config) error
	// Exit 容器退出
	Exit(hints ...func(kind, name string, err error))
}

// Register 注册组件
func Register(kind string, factory Factory) error {
	return DefaultContainer.Register(kind, factory)
}

// Load 返回组件, 若无则返回null
func Load[T any](kind, name string) T {
	c := DefaultContainer.Load(kind, name)
	r, _ := c.(T)
	return r
}

// Must 返回组件, 若无则panic
func Must[T any](kind, name string) T {
	c := DefaultContainer.Load(kind, name)
	r, ok := c.(T)
	if !ok {
		panic(fmt.Sprintf("component not-found %v.%v", kind, name))
	}
	return r
}

/*
Init 容器初始化.
配置加载顺序:
1. 若指定环境变量COMPONENT_TOML_FILE, 则直接加载, 不再往下.
2. 若存在{启动目录}/component.toml, 则直接加载, 不再往下.
3. 若存在{工作目录}/component.toml, 则直接加载, 不再往下.
4. 调用InitData()逻辑继续初始化.
*/
func Init() error {

}

/*
InitData 容器初始化.
1. 调用TemplateEngine渲染配置数据(默认是text/template)
2. 调用InitRoot()逻辑继续初始化.
*/
func InitData(data []byte) error {

}

/*
InitRoot 容器初始化.
组件初始顺序: 工厂注册顺序 > 组件创建顺序.
*/
func InitRoot(root Config) error {
	return DefaultContainer.Init(root)
}

// Reload 容器重新加载. 逻辑同Init().
func Reload() error {
	return defaultComponentContainer.Init()
}

// ReloadData 容器重新加载. 逻辑同InitData()
func ReloadData(data []byte) error {

}

// ReloadRoot 容器重新加载. 逻辑同InitRoot().
func ReloadRoot(root Config) error {

}

// Exit 容器退出. hits用于处理Component Exit错误.
func Exit(hints ...func(kind, name string, err error)) {
	defaultComponentContainer.Exit(hints...)
}

/********************************************
 * 容器默认实现
 ********************************************/

type managedComponent struct {
	Name   []string  // 名称与别名
	Target Component // 托管目录
}

type managedFactory struct {
	Kind           string                       // 工厂分类
	Target         Factory                      // 托管目录
	ComponentIndex map[string]*managedComponent // 组件索引
	ComponentArray []*managedComponent          // 组件数组
}

type defaultContainer struct {
	sync.Mutex                              // 互斥锁
	State        State                      // 0-原始, 1-Init后, 2-Exit后
	FactoryIndex map[string]*managedFactory // 工厂索引
	FactoryArray []*managedFactory          // 工厂数组
}

// Register 注册组件工厂. 注册顺序决定组件的Init()/Exit()顺序!
func (cc *defaultContainer) Register(kind string, factory Factory) error {
	cc.Mutex.Lock()
	defer cc.Mutex.Unlock()

	if _, ok := cc.Index[kind]; ok {
		return fmt.Errorf("conflicts component %v", kind)
	}
	mf := &managedFactory{
		Kind:  kind,
		Proxy: factory,
		Index: make(map[string]*managedComponent),
		Slice: make([]*managedComponent, 0, 1), // 大多数组件都是singleton
	}
	cc.Slice = append(cc.Slice, mf)
	cc.Index[kind] = mf
	return nil
}

func (cc *defaultContainer) Load(kind, name string) Component {
	cc.Mutex.Lock()
	defer cc.Mutex.Unlock()

	f := cc.Index[kind]
	if f == nil {
		return nil
	}
	c := f.Index[name]
	if c == nil {

	}
}

func (cc *defaultContainer) Init() error {

}

func (cc *defaultContainer) InitRoot(root Config) error {

}

func (cc *defaultContainer) InitData(data []byte) error {

}

func (cc *defaultContainer) Reload() error {

}

func (cc *defaultContainer) ReloadRoot(root Config) error {

}

func (cc *defaultContainer) ReloadData(data []byte) error {

}

func protectHint(hint func(kind, name string, err error), kind, name string, err error) {
	defer func() {
		if prr := recover(); prr != nil {
			log.Error("panic hint %v.", prr)
		}
	}()
	hint(kind, name, err)
}

func (cc *Container) Exit(hints ...func(kind, name string, err error)) {
	cc.Mutex.Lock()
	defer cc.Mutex.Unlock()
	// 初始化仅仅执行一次
	if cc.State < StateExit {
		for _, mf := range cc.Slice {
			for _, c := range mf.Slice {
				if err := c.Exit(); err != nil {
					log.Info("start hint %v.%v %v")
					for _, hint := range hints {

					}
				}
			}
		}
	}
}

var DefaultContainer Container = &defaultContainer{
	Index: make(map[string]*managedFactory),
	Slice: make([]*managedFactory, 0, 64),
}
