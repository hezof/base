package base

import (
	"fmt"
)

/*************************************************
 * Call 用于框架回调机制
 *************************************************/

func Call(call func(configContext *ConfigContext, managedContext *ManagedContext) error) (err error) {
	defer func() {
		if prr := recover(); prr != nil {
			err = fmt.Errorf("call panic: %v", prr)
		}
	}()
	return call(_configContext, _managedContext)
}

/*************************************************
 * 配置管理初始化
 *************************************************/

func InitConfigContext(datas ...[]byte) error {
	return _configContext.SetData(datas...)
}

/*************************************************
 * 组件托管初始化
 *************************************************/

func OverrideFactory(base string, factory ManagedFactory) {
	_managedContext.OverrideFactory(base, factory)
}

// RegisterFactory 注册组件工厂, 重复注册会panic!
func RegisterFactory(base string, factory ManagedFactory) {
	err := _managedContext.RegisterFactory(base, factory)
	if err != nil {
		panic(fmt.Errorf("register factory %v error: %v", base, err))
	}
}

// RetrieveComponent 断言组件实例, 若无base工厂会panic!
func RetrieveComponent[T any](base string, name string) T {
	component, err := _managedContext.RetrieveComponent(base, name)
	if err != nil {
		panic(fmt.Errorf("retrieve component %v.%v error: %v", base, name, err))
	}
	return component.(T)
}

func InitManagedContext() error {
	return _managedContext.Init(_configContext)
}

func ReloadManagedContext(reloadPolicy func(base string, config *ManagedConfig, newValues map[string]any) bool) {
	_managedContext.Reload(_configContext, reloadPolicy)
}

func ExitManagedContext(hints ...func(base string, config *ManagedConfig, err error)) {
	_managedContext.Exit(hints...)
}
