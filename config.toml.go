package core

import (
	"github.com/hezof/core/internal/toml"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

/********************************************
 * 配置及插件
 ********************************************/

// ConfigPlugin 配置处理器(模板,属性)
type ConfigPlugin interface {
	Exec(data []byte) ([]byte, error)
}

// ConfigContext 配置结构
// 支持多配置文件, 各个配置文件的数据独立存储,不做合并. 所以values是个slice!
type ConfigContext struct {
	sync.RWMutex
	Values []map[string]any
}

func (c *ConfigContext) SetData(datas ...[]byte) (err error) {
	c.Lock()
	defer c.Unlock()

	for _, data := range datas {
		value := make(map[string]any)
		if err = toml.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Values = append(c.Values, value)
	}
	return nil
}

func (c *ConfigContext) GetFirst(path string) (any, bool) {
	c.RLock()
	defer c.RUnlock()

	/*
		查找流程: 遍历多配置文件(按设置顺序), 返回第一个符合path的配置
	*/
	for _, value := range c.Values {
		val, ok := ExtractConfig(value, path)
		if ok {
			return val, true
		}
	}
	return nil, false
}

func (c *ConfigContext) GetLast(path string) (any, bool) {
	c.RLock()
	defer c.RUnlock()

	/*
		查找流程: 遍历多配置文件(按设置逆序), 返回最后一个符合path的配置.
	*/
	for i := len(c.Values) - 1; i >= 0; i-- {
		val, ok := ExtractConfig(c.Values[i], path)
		if ok {
			return val, true
		}
	}
	return nil, false
}

func (c *ConfigContext) GetAll(path string) []any {
	c.RLock()
	defer c.RUnlock()

	/*
		查找流程: 遍历多配置文件(按设置顺序), 返回所有符合path的配置.
	*/
	var all []any
	for _, value := range c.Values {
		val, ok := ExtractConfig(value, path)
		if ok {
			all = append(all, val)
		}
	}
	return all
}

/********************************************
 * 配置文件及辅助函数
 ********************************************/

// CONFIG_TOML_FILE 配置文件
const CONFIG_TOML_FILE = "CONFIG_TOML_FILE"

// CONFIG_TOML_NAME 配置名称
const CONFIG_TOML_NAME = "config.toml"

func ConfigTomlFile() ([]string, error) {

	/********************************************
	 * 默认配置加载方式:
	 * 1. 环境变量指定: COMPONENT_TOML_FILE. 该方式支持逗号数组与glob模式通配符:
	 * - 星号（*）：匹配任意数量的字符
	 * - 问号（?）：匹配任意单个字符
	 * - 方括号（[]）：匹配括号内的任意一个字符
	 * 2. 默认配置扫描
	 * - {启动目录}/component.toml
	 * - {工作目录}/component.toml
	 ********************************************/

	if env := os.Getenv(CONFIG_TOML_FILE); env != "" {
		val := make([]string, 0, 4)
		for _, v := range strings.Split(env, ",") {
			v = strings.TrimSpace(v)
			if v != "" {
				if strings.IndexAny(v, "*?[") != -1 {
					matches, err := filepath.Glob(v)
					if err != nil {
						return nil, err
					}
					for _, match := range matches {
						val = append(val, match)
					}
				} else {
					val = append(val, v)
				}
			}
		}
		return Uniq(val), nil
	} else {
		path, err := LocatePath(CONFIG_TOML_NAME)
		if err != nil {
			return nil, err
		}
		if path != "" {
			return []string{path}, nil
		}
		return nil, nil
	}
}

// MarshalConfig 编码配置
func MarshalConfig(val map[string]any) ([]byte, error) {
	return toml.Marshal(val)
}

// UnmarshalConfig 解码配置
func UnmarshalConfig(data []byte) (map[string]any, error) {
	cfg := make(map[string]any)
	err := toml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// ExtractConfig 根据path查询节点,并返回结果. path格式同toml定义
func ExtractConfig(val map[string]any, path string) (any, bool) {
	for {
		pos := strings.IndexByte(path, '.')
		if pos == -1 {
			rt, ok := val[path]
			return rt, ok
		}
		val, _ = val[path[:pos]].(map[string]any)
		if val == nil {
			return nil, false
		}
		path = path[pos+1:]
	}
}

// 默认配置对象
var (
	_configContext         = new(ConfigContext)
	_environConfigTemplate = new(EnvironConfigTemplate)
)
