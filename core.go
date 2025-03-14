package core

import "fmt"

/*************************************************
 * 钩子机制
 *************************************************/

// OnInit 在Init时回调
func OnInit(hook ...func(root ...map[string]any)) {

}

// OnReload 在Reload时回调
func OnReload(hook ...func(root ...map[string]any)) {

}

// OnExit 在Exit时回调
func OnExit(hook ...func(root ...map[string]any)) {

}

/*************************************************
 * 容器功能
 *************************************************/

// Register 注册组件工厂, 重复注册会panic!
func Register(kind string, factory ComponentFactory) {
	_container.MustRegisterFactory(kind, factory)
}

// Component 断言组件实例, 若无kind工厂会panic!
func Component[T any](kind, name string) T {

}

/*************************************************
 * 核心功能
 *************************************************/

func Init() {
	// 1. 初始toml配置
	if err := _configToml.InitTomlFile(ConfigTomlFile()); err != nil {
		panic(fmt.Errorf("init toml file error: %v", err))
	}
}

func InitData(datas ...[]byte) {
	// 1. 初始toml配置
	if err := _configToml.InitTomlData(datas, nil); err != nil {
		panic(fmt.Errorf("init toml data error: %v", err))
	}

}

func Reload() error {

}

func ReloadConfig(root ...map[string]any) {

}

func Exit(hints ...func(kind, name string, err error)) {

}
