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

type hook struct {
	Join JoinPoint
	Name string
	Call func(config *ConfigContext, managed *ManagedContext)
}

func (h hook) Exec(config *ConfigContext, managed *ManagedContext) {
	defer func() {
		if prr := recover(); prr != nil {
			log.Error("exec hook %v %v panic: %v|%v", joinPointNames[h.Join], h.Name, prr, StackTrace(2, `|`))
		}
	}()
	h.Call(config, managed)
}

type Hooks struct {
	value []*hook
}

func (h *Hooks) Join(join JoinPoint, name string, call func(config *ConfigContext, managed *ManagedContext)) {
	h.value = append(h.value, &hook{
		Join: join,
		Name: name,
		Call: call,
	})
}

func (h *Hooks) Exec(join JoinPoint, config *ConfigContext, managed *ManagedContext) {
	for _, hook := range h.value {
		if hook.Join == join {
			log.Info("exec hook %v %v", joinPointNames[join], hook.Name)
			hook.Exec(config, managed)
		}
	}
}

var _hooks = new(Hooks)
