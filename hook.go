package core

import "github.com/hezof/log"

type JoinPoint uint8

const (
	BeforeInit   JoinPoint = 0x01
	AfterInit    JoinPoint = 0x02
	BeforeReload JoinPoint = 0x04
	AfterReload  JoinPoint = 0x08
	BeforeExit   JoinPoint = 0x10
	AfterExit    JoinPoint = 0x20
)

var joinPointNames = map[JoinPoint]string{
	BeforeInit:   "before init",
	AfterInit:    "after init",
	BeforeReload: "before reload",
	AfterReload:  "after reload",
	BeforeExit:   "before exit",
	AfterExit:    "after exit",
}

type Hook struct {
	Join JoinPoint
	Name string
	Call func(config *ConfigContext, managed *ManagedContext)
}

func (h Hook) Exec(config *ConfigContext, managed *ManagedContext) {
	defer func() {
		if prr := recover(); prr != nil {
			log.Error("hook exec panic: %v, %v", h.Name, prr)
		}
	}()
	h.Call(config, managed)
}

var _hooks = make([]*Hook, 0, 4)

func JoinHook(join JoinPoint, name string, call func(config *ConfigContext, managed *ManagedContext)) {
	log.Info("join hook %v, %v", joinPointNames[join], name)
	_hooks = append(_hooks, &Hook{
		Join: join,
		Name: name,
		Call: call,
	})
}

func ExecHook(join JoinPoint, config *ConfigContext, managed *ManagedContext) {
	for _, hook := range _hooks {
		if hook.Join == join {
			log.Info("exec hook %v, %v", joinPointNames[join], hook.Name)
			hook.Exec(config, managed)
		}
	}
}
