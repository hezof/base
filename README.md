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

- 

## 组件托管(TODO)

## 钩子机制(TODO)

