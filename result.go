package clients

import (
	"fmt"
	"github.com/hezof/core"
	"github.com/hezof/protojson"
)

type StatusResult core.StatusResult

func (sr *StatusResult) Error() string {
	// 不能使用ToJson()会在EncodeField过程形成死循环
	return core.ToJson(sr)
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

var _ core.Error = (*StatusResult)(nil)

func (sr *StatusResult) DecodeField(r *protojson.JsonDecoder, f string) {
	switch f {
	case core.ResultCodeField:
		protojson.DecodeUint32(r, &sr.Code)
	case core.ResultNameField:
		protojson.DecodeString(r, &sr.Name)
	case core.ResultMessageField:
		protojson.DecodeString(r, &sr.Message)
	case core.ResultDataField:
		protojson.DecodeAny(r, sr.Data)
	}
}

func (sr *StatusResult) EncodeField(w *protojson.JsonEncoder) {
	protojson.EncodeUint32_WithEmpty(w, core.ResultCodeField, sr.Code)
	protojson.EncodeString_OmitEmpty(w, core.ResultNameField, sr.Name)
	protojson.EncodeString_OmitEmpty(w, core.ResultMessageField, sr.Message)
	protojson.EncodeAny_OmitEmpty(w, core.ResultDataField, sr.Data)
}

var _ protojson.FieldCodec = (*StatusResult)(nil)
