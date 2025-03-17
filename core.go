package core

import "fmt"

/*************************************************
 * 钩子机制
 *************************************************/

// OnInit 在Init时回调
func OnInit(name string, call func(config *ConfigToml)) {
	OnInitHook = append(OnInitHook, &Hook{
		Name: name,
		Call: call,
	})
}

// OnReload 在Reload时回调
func OnReload(name string, call func(config *ConfigToml)) {
	OnReloadHook = append(OnReloadHook, &Hook{
		Name: name,
		Call: call,
	})
}

// OnExit 在Exit时回调
func OnExit(name string, call func(config *ConfigToml)) {
	OnExitHook = append(OnExitHook, &Hook{
		Name: name,
		Call: call,
	})
}

/*************************************************
 * 容器功能
 *************************************************/

// Register 注册组件工厂, 重复注册会panic!
func Register(kind string, factory ManagedFactory) {
	_managed.MustRegisterFactory(kind, factory)
}

// Resource 断言组件实例, 若无kind工厂会panic!
func Resource[T any](kind, name string) T {
	return _managed.MustRetrieveFactory(kind).GetOrNewComponent(name).(T)
}

/*************************************************
 * 核心功能
 *************************************************/

func Init() {
	// 0. 初始配置
	if err := _configToml.InitTomlFile(ConfigTomlFile()); err != nil {
		panic(fmt.Errorf("init toml file error: %v", err))
	}
	// 1. 回调钩子
	for _, hook := range OnInitHook {
		hook.Exec(&_configToml)
	}
	// 2. 加载托管
	_managed.MustInit(_configToml)
}

func InitData(datas ...[]byte) {
	// 0. 初始配置
	if err := _configToml.InitTomlData(datas, nil); err != nil {
		panic(fmt.Errorf("init toml data error: %v", err))
	}
	// 1. 回调钩子
	for _, hook := range OnInitHook {
		hook.Exec(&_configToml)
	}
}

func Reload() error {
	// 0. 初始配置
	if err := _configToml.InitTomlFile(ConfigTomlFile()); err != nil {
		panic(fmt.Errorf("init toml file error: %v", err))
	}
	// 1. 回调钩子
	for _, hook := range OnReloadHook {
		hook.Exec(&_configToml)
	}

}

func ReloadData(datas ...[]byte) {
	// 0. 初始配置
	if err := _configToml.InitTomlData(datas, nil); err != nil {
		panic(fmt.Errorf("init toml data error: %v", err))
	}
	// 1. 回调钩子
	for _, hook := range OnReloadHook {
		hook.Exec(&_configToml)
	}
}

func Exit(hints ...func(kind, name string, err error)) {
	// 1. 回调钩子
	for _, hook := range OnExitHook {
		hook.Exec(&_configToml)
	}
}
