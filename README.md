# package ioc

ioc库. 提供

- 基本工具
- 类型转换, 包括MapStruct
- 配置管理
- 组件托管

## 基本工具

- Nvl{T}
- Min
- Max
- Keys
- Vals
- Uniq
- ...

## 类型转换

- func As{Type} ({Type}, error)
- func To{Type} {Type}
- MapStruct(m map[string]any, s any, tag string) error

## 配置管理(TODO)

- type ConfigContext struct

```
// ConfigContext 配置结构
// 支持多配置文件, 各个配置文件的数据独立存储,不做合并. 所以values是个slice!
type ConfigContext struct {
	sync.RWMutex
	Plugin ConfigPlugin
	Values []map[string]any
}
```

- func ConfigBool(path string) (bool, error)
- func ConfigString(path string) (string, error)
- func ConfigTime(path string) (time.Time, error)
- func ConfigDuration(path string) (time.Duration, error)
- func ConfigInt(path string) (int, error)
- func ConfigInt8(path string) (int8, error)
- func ConfigInt16(path string) (int16, error)
- func ConfigInt32(path string) (int32, error)
- func ConfigInt64(path string) (int64, error)
- func ConfigUint(path string) (uint, error)
- func ConfigUint8(path string) (uint8, error)
- func ConfigUint16(path string) (uint16, error)
- func ConfigUint32(path string) (uint32, error)
- func ConfigUint64(path string) (uint64, error)
- func ConfigFloat32(path string) (float32, error)
- func ConfigFloat64(path string) (float64, error)
- func MustConfigBool(path string) bool
- func MustConfigString(path string) string
- func MustConfigTime(path string) time.Time
- func MustConfigDuration(path string) time.Duration
- func MustConfigInt(path string) int
- func MustConfigInt8(path string) int8
- func MustConfigInt16(path string) int16
- func MustConfigInt32(path string) int32
- func MustConfigInt64(path string) int64
- func MustConfigUint(path string) uint
- func MustConfigUint8(path string) uint8
- func MustConfigUint16(path string) uint16
- func MustConfigUint32(path string) uint32
- func MustConfigUint64(path string) uint64
- func MustConfigFloat32(path string) float32
- func MustConfigFloat64(path string) float64
- func OptiConfigBool(path string, def bool) bool
- func OptiConfigString(path string, def string) string
- func OptiConfigTime(path string, def time.Time) time.Time
- func OptiConfigDuration(path string, def time.Duration) time.Duration
- func OptiConfigInt(path string, def int) int
- func OptiConfigInt8(path string, def int8) int8
- func OptiConfigInt16(path string, def int16) int16
- func OptiConfigInt32(path string, def int32) int32
- func OptiConfigInt64(path string, def int64) int64
- func OptiConfigUint(path string, def uint) uint
- func OptiConfigUint8(path string, def uint8) uint8
- func OptiConfigUint16(path string, def uint16) uint16
- func OptiConfigUint32(path string, def uint32) uint32
- func OptiConfigUint64(path string, def uint64) uint64
- func OptiConfigFloat32(path string, def float32) float32
- func OptiConfigFloat64(path string, def float64) float64
- func ConfigStruct(path string, structPointer any, tag string) (bool, error)

## 组件托管(TODO)

- type ManagedContext struct

```
// ManagedContext 托管容器
type ManagedContext struct {
	sync.RWMutex                                 // 只在上下文设置读写锁
	Proxies      map[string]*ManagedFactoryProxy // 注册工厂. 键是kind, 值是托管工厂
	Indexes      []string                        // 注册顺序, 决定Init()/Exit()的组件顺序
}
```

- func Register(base string, factory ManagedFactory)
- func Component[T any](base string, name string) T