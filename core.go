package core

import "fmt"

/*************************************************
 * 容器功能
 *************************************************/

// Register 注册组件工厂, 重复注册会panic!
func Register(kind string, factory ManagedFactory) {
	err := _managedContext.RegisterFactory(kind, factory)
	if err != nil {
		panic(fmt.Errorf("register factory error: %v, %v", kind, err))
	}
}

// Resource 断言组件实例, 若无kind工厂会panic!
func Resource[T any](kind, name string) T {
	factory := _managedContext.RetrieveFactory(kind)
	if factory == nil {
		panic(fmt.Errorf("retrieve factory empty: %v", kind))
	}
	return factory.GetOrNewComponent(name).(T)
}

/*************************************************
 * 核心功能
 *************************************************/

func Init() {
	datas, err := ReadFile(ConfigTomlFile())
	if err != nil {
		panic(fmt.Errorf("init config context error: %v", err))
	}
	InitTomlData(datas...)
}

func InitTomlData(datas ...[]byte) {
	// 0. 初始配置
	if err := _configContext.SetTomlData(datas...); err != nil {
		panic(fmt.Errorf("init config context error: %v", err))
	}

	// 1. 回调钩子
	ExecHook(BeforeInit, _configContext, _managedContext)

	// 2. 初始托管
	if err := _managedContext.Init(_configContext); err != nil {
		panic(fmt.Errorf("init managed context error: %v", err))
	}

	// 最后回调钩子
	ExecHook(AfterInit, _configContext, _managedContext)
}

func Reload() error {
	datas, err := ReadFile(ConfigTomlFile())
	if err != nil {
		return fmt.Errorf("reload config context error: %v", err)
	}
	return ReloadTomlData(datas...)
}

func ReloadTomlData(datas ...[]byte) error {
	// 0. 重载配置
	if err := _configContext.SetTomlData(datas...); err != nil {
		return fmt.Errorf("reload config context error: %v", err)
	}

	// 1. 回调钩子
	ExecHook(BeforeReload, _configContext, _managedContext)

	// 2. 重载托管
	if err := _managedContext.Reload(_configContext); err != nil {
		return fmt.Errorf("reload managed context error: %v", err)
	}

	// 最后回调钩子
	ExecHook(AfterReload, _configContext, _managedContext)

	return nil
}

func Exit(hints ...func(kind, name string, err error)) {
	// 1. 回调钩子
	ExecHook(BeforeReload, _configContext, _managedContext)

	// 2. 退出托管
	_managedContext.Exit(hints...)

	// 最后回调钩子
	ExecHook(AfterReload, _configContext, _managedContext)
}
