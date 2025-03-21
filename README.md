# framework

框架库. 提供:

- 类型转换, 包括MapStruct
- 统一错误(结果)
- 配置管理
- 组件托管
- 钩子机制
- jsonrpc

## 类型转换

- func As{Type} ({Type}, error)
- func To{Type} {Type}
- MapStruct(m map[string]any, s any, tag string) error

## 统一错误(结果)

```
/*
StatusResult统一结果与错误的数据结构, 并实现与Grpc Error的转换.
因为Grpc Error Status只有Code字段, 约定StatusResult的Status/Code分别存储在高10位与低22位!

约定StatusResult Code取值范围:
- [0,17)         表示保留错误码! grpc内置错误码, 参考codes._maxCode
- [17,4194393)   表示业务错误码! 最大值(2^22 - 1)! 因为Grpc Code的前10位用于表示StatusResult Status

约定StatusResult Status取值范围:
- (0,511]
*/
const (
	ErrorCodeBits   = 22 // 由于grpc的问题,  int32需要保留一个符号位
	ErrorCodeMask   = 1<<ErrorCodeBits - 1
	ErrorStatusBits = 9
	ErrorStatusMask = 1<<ErrorStatusBits - 1
)

// Error 带状态码的结果
type Error interface {
	error
	GetCode() uint32
	GetStatus() uint32
	SetStatus(status uint32)
	GetName() string
	SetName(name string)
	GetMessage() string
	SetMessage(message string)
	GetDetails() []string
}

// StatusResult 带状态的结果. 必须注意status与code的约定取值范围!
type StatusResult struct {
	Status  uint32   `json:"status,omitempty"`  // 状态代码(http).
	Code    uint32   `json:"code"`              // 错误代码. 0表示成功
	Name    string   `json:"name,omitempty"`    // 错误名称. OK表示成功
	Message string   `json:"message,omitempty"` // 错误消息.
	Details []string `json:"-"`                 // 错误参数.
	Data    any      `json:"-"`                 // 结果数据
}
```

- func StatusError(status uint32, code uint32, message string, details ...string) *StatusResult
- func StatusErrorFrom(err error) *StatusResult

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
- func Init()
- func InitTomlData(datas ...[]byte)
- func Reload(reloadPolicy func(base string, config *ManagedConfig, newValues map[string]any) bool) error
- func ReloadTomlData(reloadPolicy func(base string, config *ManagedConfig, newValues map[string]any) bool, datas ...[]
  byte) error
- func Exit(hints ...func(base string, config *ManagedConfig, err error))

## 钩子机制(TODO)

- type JoinPoint uint8

```
const (
	BeforeInit   JoinPoint = 0x01
	AfterInit    JoinPoint = 0x02
	BeforeReload JoinPoint = 0x04
	AfterReload  JoinPoint = 0x08
	BeforeExit   JoinPoint = 0x10
	AfterExit    JoinPoint = 0x20
)

```

- func JoinHook(join JoinPoint, name string, call func(config *ConfigContext, managed *ManagedContext))
- func ExecHook(join JoinPoint, config *ConfigContext, managed *ManagedContext) {
