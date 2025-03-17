package core

import "github.com/hezof/log"

var (
	OnInitHook   []*Hook
	OnReloadHook []*Hook
	OnExitHook   []*Hook
)

type Hook struct {
	Name string
	Call func(config *ConfigToml)
}

func (h Hook) Exec(config *ConfigToml) {
	defer func() {
		if prr := recover(); prr != nil {
			log.Error("hook exec panic: %v, %v", h.Name, prr)
		}
	}()
	h.Call(config)
}
