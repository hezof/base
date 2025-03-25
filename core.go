package core

import (
	"fmt"
	"github.com/hezof/log"
)

/*************************************************
 * 容器功能
 *************************************************/

// Factory 注册组件工厂, 重复注册会panic!
func Factory(base string, factory ManagedFactory) {
	err := _managedContext.Register(base, factory)
	if err != nil {
		panic(fmt.Errorf("register factory %v error: %v", base, err))
	}
}

// Component 断言组件实例, 若无base工厂会panic!
func Component[T any](base string, name string) T {
	component, err := _managedContext.Retrieve(base, name)
	if err != nil {
		panic(fmt.Errorf("retrieve component %v.%v error: %v", base, name, err))
	}
	return component.(T)
}

/*************************************************
 * 核心功能
 *************************************************/

func Hook(join JoinPoint, name string, call func(config *ConfigContext, managed *ManagedContext)) {
	log.Info("join hook %v %v", joinPointNames[join], name)
	_hooks.Join(join, name, call)
}

func Template(template ConfigTemplate) {
	log.Info("install template %T ", template)
	_configTemplate = template
}

func Init() {
	datas, err := ReadFile(ConfigTomlFile())
	if err != nil {
		panic(fmt.Errorf("init config context error: %v", err))
	}
	if _configTemplate != nil {
		for i, data := range datas {
			datas[i], err = _configTemplate.Exec(data)
			if err != nil {
				panic(fmt.Errorf("init config context error: %v", err))
			}
		}
	}
	InitData(datas...)
}

func InitData(datas ...[]byte) {
	// 0. 初始配置
	if err := _configContext.SetData(datas...); err != nil {
		panic(fmt.Errorf("init config context error: %v", err))
	}

	// 1. 回调钩子
	_hooks.Exec(BeforeInit, _configContext, _managedContext)

	// 2. 初始托管(错误中止)
	if err := _managedContext.Init(_configContext); err != nil {
		panic(fmt.Errorf("init managed context error: %v", err))
	}

	// 最后回调钩子
	_hooks.Exec(AfterInit, _configContext, _managedContext)
}

func Reload(reloadPolicy func(base string, config *ManagedConfig, newValues map[string]any) bool) error {
	datas, err := ReadFile(ConfigTomlFile())
	if err != nil {
		return fmt.Errorf("reload config context error: %v", err)
	}
	if _configTemplate != nil {
		for i, data := range datas {
			datas[i], err = _configTemplate.Exec(data)
			if err != nil {
				return fmt.Errorf("reload config context error: %v", err)
			}
		}
	}
	return ReloadData(reloadPolicy, datas...)
}

func ReloadData(reloadPolicy func(base string, config *ManagedConfig, newValues map[string]any) bool, datas ...[]byte) error {
	// 0. 重载配置
	if err := _configContext.SetData(datas...); err != nil {
		return fmt.Errorf("reload config context error: %v", err)
	}

	// 1. 回调钩子
	_hooks.Exec(BeforeReload, _configContext, _managedContext)

	// 2. 重载托管(错误忽略)
	if err := _managedContext.Reload(_configContext, reloadPolicy); err != nil {
		return fmt.Errorf("reload managed context error: %v", err)
	}

	// 最后回调钩子
	_hooks.Exec(AfterReload, _configContext, _managedContext)

	return nil
}

func Exit(hints ...func(base string, config *ManagedConfig, err error)) {

	// 1. 回调钩子
	_hooks.Exec(BeforeReload, _configContext, _managedContext)

	// 2. 退出托管(容忍错误)
	_managedContext.Exit(hints...)

	// 最后回调钩子
	_hooks.Exec(AfterReload, _configContext, _managedContext)
}
