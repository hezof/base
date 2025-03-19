package base

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

/*
StatusResult统一结果与错误的数据结构, 并实现与Grpc Error的转换.
因为Grpc Error Status只有Code字段, 约定StatusResult的Status/Code分别存储在高9位与低22位! 由于int32需要保留一个符号位

约定StatusResult Code取值范围:
- [0,17)         表示保留错误码! grpc内置错误码, 参考codes._maxCode
- [17,4194393)   表示业务错误码! 最大值(2^22 - 1)! 因为Grpc Code的前10位用于表示StatusResult Status

约定StatusResult Status取值范围:
- (0,511]
*/
const (
	CodeBits   = 22 // 由于grpc的问题,  int32需要保留一个符号位
	CodeMask   = 1<<CodeBits - 1
	StatusBits = 9
	StatusMask = 1<<StatusBits - 1
)

// Error 带状态码的错误
type Error interface {
	error
	GetCode() uint32
	GetStatus() uint32
	GetName() string
	GetMessage() string
	GetDetails() []string
}

// ErrorModel 带状态的结果. 必须注意status与code的约定取值范围!
type ErrorModel struct {
	Status  uint32   // 状态代码(http).
	Code    uint32   // 错误代码. 0表示成功
	Name    string   // 错误名称. OK表示成功
	Message string   // 错误消息.
	Details []string `json:"-"` // 错误参数.
}

func (sr *ErrorModel) Error() string {
	return ToJson(sr)
}

func (sr *ErrorModel) GetCode() uint32 {
	return sr.Code
}

func (sr *ErrorModel) GetStatus() uint32 {
	return sr.Status
}

func (sr *ErrorModel) GetName() string {
	return sr.Name
}

func (sr *ErrorModel) GetMessage() string {
	return sr.Message
}

func (sr *ErrorModel) GetDetails() []string {
	return sr.Details
}

var _ Error = (*ErrorModel)(nil)

// StatusError 创建StatusResult错误实例. 必须注意status与code的取值范围:
// - Status 取值范围(0,1024)
// - Code 取值范围(0,4194304)
func StatusError(status uint32, code uint32, message string, details ...string) Error {

	status &= StatusMask
	code &= CodeMask

	if len(details) > 0 {
		message = fmt.Sprintf(message, AnySlice(details)...)
	}
	return &ErrorModel{
		Status:  status,
		Code:    code,
		Message: message,
		Details: details,
	}
}

// StackTrace 打印堆栈追踪信息,如果是"/src/runtime/"自动跳过!
func StackTrace(skip int, sep string) string {
	var sb strings.Builder
	for i := 1; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			return sb.String()
		}
		// 过滤runtime的行项,避免错误日志过多!
		if strings.Index(file, "/src/runtime/") == -1 {
			if skip > 0 {
				skip--
			} else {
				if sb.Len() > 0 {
					sb.WriteString(sep)
				}
				sb.WriteString(file)
				sb.WriteByte(':')
				sb.WriteString(strconv.Itoa(line))
			}
		}
	}
}
