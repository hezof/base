package core

import (
	"bytes"
	"github.com/hezof/core/internal/toml"
	"github.com/hezof/log"
	"os"
	"strings"
	"text/template"
)

var _configTemplate ConfigTemplate = new(EnvironConfigTemplate)

type ConfigTemplate interface {
	Exec(data []byte) ([]byte, error)
}

// EnvironConfigTemplate 内置配置插件(environment + text template)
type EnvironConfigTemplate struct{}

func (cp *EnvironConfigTemplate) Data() map[string]any {
	/*
		数据逻辑: 读取所有环境变量作为数据, 默认所有值都是string, 但支持toml的array([...])与table({...})
	*/
	data := make(map[string]any)
	for _, env := range os.Environ() {
		pos := strings.IndexByte(env, '=')
		if pos == -1 {
			data[env] = ""
		} else {
			key := strings.TrimSpace(env[:pos])
			val := strings.TrimSpace(env[pos+1:])
			// 支持toml的Array与Table
			switch {
			case strings.HasPrefix(val, "["):
				arr := make([]any, 0, 4)
				if err := toml.Unmarshal([]byte(val), &arr); err != nil {
					log.Error("unmarshal toml array %v error: %v", key, err)
					data[key] = val
				} else {
					data[key] = arr
				}
			case strings.HasPrefix(val, "{"):
				tbl := make(map[string]any)
				if err := toml.Unmarshal([]byte(val), &tbl); err != nil {
					log.Error("unmarshal toml table %v error: %v", key, err)
					data[key] = val
				} else {
					data[key] = tbl
				}
			default:
				data[key] = val
			}
		}
	}
	return data
}

func (cp *EnvironConfigTemplate) Exec(data []byte) ([]byte, error) {
	/*
		默认EnvironConfigPlugin插件使用text/template模板引擎, 其它引擎请请用SetConfigPlugin()替换
	*/
	tpl, err := template.New("").Parse(string(data))
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(make([]byte, 0, 2048))
	err = tpl.Execute(buf, cp.Data())
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
