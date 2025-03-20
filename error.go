package base

import (
	"fmt"
	"net/http"
)

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

func (sr *StatusResult) Error() string {
	// 不能使用ToJson()会在EncodeField过程形成死循环
	return ToJson(sr)
}

func (sr *StatusResult) GetCode() uint32 {
	return sr.Code
}

func (sr *StatusResult) GetStatus() uint32 {
	return sr.Status
}

func (sr *StatusResult) SetStatus(status uint32) {
	sr.Status = status
}

func (sr *StatusResult) GetName() string {
	return sr.Name
}

func (sr *StatusResult) SetName(name string) {
	sr.Name = name
}

func (sr *StatusResult) GetMessage() string {
	return sr.Message
}

func (sr *StatusResult) SetMessage(message string) {
	if len(sr.Details) > 0 {
		message = fmt.Sprintf(message, sr.Details)
	}
	sr.Message = message
}

func (sr *StatusResult) GetDetails() []string {
	return sr.Details
}

var _ Error = (*StatusResult)(nil)

func (sr *StatusResult) DecodeField(r *JsonDecoder, f string) {
	switch f {
	case "code":
		DecodeUint32(r, &sr.Code)
	case "name":
		DecodeString(r, &sr.Name)
	case "message":
		DecodeString(r, &sr.Message)
	case "data":
		DecodeAny(r, sr.Data)
	}
}

func (sr *StatusResult) EncodeField(w *JsonEncoder) {
	EncodeUint32_WithEmpty(w, "code", sr.Code)
	EncodeString_OmitEmpty(w, "name", sr.Name)
	EncodeString_OmitEmpty(w, "message", sr.Message)
	EncodeAny_OmitEmpty(w, "data", sr.Data)
}

var _ FieldCodec = (*StatusResult)(nil)

// StatusError 创建StatusResult错误实例. 必须注意status与code的取值范围:
// - Status 取值范围(0,1024)
// - Code 取值范围(0,4194304)
func StatusError(status uint32, code uint32, message string, details ...string) *StatusResult {

	status &= ErrorStatusMask
	code &= ErrorCodeMask

	if len(details) > 0 {
		message = fmt.Sprintf(message, AnySlice(details)...)
	}
	return &StatusResult{
		Status:  status,
		Code:    code,
		Message: message,
		Details: details,
	}
}

// StatusErrorFrom 定义统一的error转换为*Result规则
func StatusErrorFrom(err error) *StatusResult {

	// 内部错误
	if val, ok := err.(*StatusResult); ok {
		return val
	}

	// 接口错误
	if val, ok := err.(Error); ok {
		return &StatusResult{
			Status:  val.GetStatus(),
			Code:    val.GetCode(),
			Message: val.GetMessage(),
			Details: val.GetDetails(),
		}
	}

	// 其他错误
	return &StatusResult{
		Status:  http.StatusInternalServerError,
		Code:    2, // Grpc Unknown是2
		Message: err.Error(),
	}
}
