package base

import (
	"fmt"
	"math"
)

/*************************************
	bool类型: EncodeBool_<empty>
 *************************************/

func EncodeBool(w *JsonEncoder, value bool) {
	if value {
		encodeTrue(w)
	} else {
		encodeFalse(w)
	}
}

func EncodeBool_OmitEmpty(w *JsonEncoder, name string, value bool) {
	if value {
		encodeTrueMember(w, name)
	}
}

func EncodeBool_WithEmpty(w *JsonEncoder, name string, value bool) {
	if value {
		encodeTrueMember(w, name)
	} else {
		encodeFalseMember(w, name)
	}
}

func EncodeBool_ConvEmpty(w *JsonEncoder, name string, value bool) {
	if value {
		encodeTrueMember(w, name)
	} else {
		encodeFalseMember(w, name)
	}
}

/*************************************
	int32类型: EncodeInt32_<empty>
 *************************************/

func EncodeInt32(w *JsonEncoder, value int32) {
	if value != 0 {
		encodeInt32(w, value)
	} else {
		encodeZero(w)
	}
}

func EncodeInt32_OmitEmpty(w *JsonEncoder, name string, value int32) {
	if value != 0 {
		encodeInt32Member(w, name, value)
	}
}

func EncodeInt32_WithEmpty(w *JsonEncoder, name string, value int32) {
	if value != 0 {
		encodeInt32Member(w, name, value)
	} else {
		encodeZeroMember(w, name)
	}
}

func EncodeInt32_ConvEmpty(w *JsonEncoder, name string, value int32) {
	if value != 0 {
		encodeInt32Member(w, name, value)
	} else {
		encodeZeroMember(w, name)
	}
}

/*************************************
	int64类型: EncodeInt64_<empty>
 *************************************/

func EncodeInt64(w *JsonEncoder, value int64) {
	if value != 0 {
		encodeInt64(w, value)
	} else {
		encodeZero(w)
	}
}

func EncodeInt64_OmitEmpty(w *JsonEncoder, name string, value int64) {
	if value != 0 {
		encodeInt64Member(w, name, value)
	}
}

func EncodeInt64_WithEmpty(w *JsonEncoder, name string, value int64) {
	if value != 0 {
		encodeInt64Member(w, name, value)
	} else {
		encodeZeroMember(w, name)
	}
}

func EncodeInt64_ConvEmpty(w *JsonEncoder, name string, value int64) {
	if value != 0 {
		encodeInt64Member(w, name, value)
	} else {
		encodeZeroMember(w, name)
	}
}

/*************************************
	uint32类型: EncodeUint32_<empty>
 *************************************/

func EncodeUint32(w *JsonEncoder, value uint32) {
	if value != 0 {
		encodeUint32(w, value)
	} else {
		encodeZero(w)
	}
}

func EncodeUint32_OmitEmpty(w *JsonEncoder, name string, value uint32) {
	if value != 0 {
		encodeUint32Member(w, name, value)
	}
}

func EncodeUint32_WithEmpty(w *JsonEncoder, name string, value uint32) {
	if value != 0 {
		encodeUint32Member(w, name, value)
	} else {
		encodeZeroMember(w, name)
	}
}

func EncodeUint32_ConvEmpty(w *JsonEncoder, name string, value uint32) {
	if value != 0 {
		encodeUint32Member(w, name, value)
	} else {
		encodeZeroMember(w, name)
	}
}

/*************************************
	uin64类型: EncodeUint64_<empty>
 *************************************/

func EncodeUint64(w *JsonEncoder, value uint64) {
	if value != 0 {
		encodeUint64(w, value)
	} else {
		encodeZero(w)
	}
}

func EncodeUint64_OmitEmpty(w *JsonEncoder, name string, value uint64) {
	if value != 0 {
		encodeUint64Member(w, name, value)
	}
}

func EncodeUint64_WithEmpty(w *JsonEncoder, name string, value uint64) {
	if value != 0 {
		encodeUint64Member(w, name, value)
	} else {
		encodeZeroMember(w, name)
	}
}

func EncodeUint64_ConvEmpty(w *JsonEncoder, name string, value uint64) {
	if value != 0 {
		encodeUint64Member(w, name, value)
	} else {
		encodeZeroMember(w, name)
	}
}

/*************************************
	float32类型: EncodeFloat_<empty>
 *************************************/

func EncodeFloat(w *JsonEncoder, value float32) {
	if math.IsInf(float64(value), 0) || math.IsNaN(float64(value)) {
		if w.firstError == nil {
			w.firstError = fmt.Errorf("unsupported value: %f", value)
		}
		return
	}
	if value != 0 {
		encodeFloat(w, value)
	} else {
		encodeZero(w)
	}
}

func EncodeFloat_OmitEmpty(w *JsonEncoder, name string, value float32) {
	if math.IsInf(float64(value), 0) || math.IsNaN(float64(value)) {
		if w.firstError == nil {
			w.firstError = fmt.Errorf("unsupported value: %f", value)
		}
		return
	}
	if value != 0 {
		encodeFloatMember(w, name, value)
	}
}

func EncodeFloat_WithEmpty(w *JsonEncoder, name string, value float32) {
	if math.IsInf(float64(value), 0) || math.IsNaN(float64(value)) {
		if w.firstError == nil {
			w.firstError = fmt.Errorf("unsupported value: %f", value)
		}
		return
	}
	if value != 0 {
		encodeFloatMember(w, name, value)
	} else {
		encodeZero(w)
	}
}

func EncodeFloat_ConvEmpty(w *JsonEncoder, name string, value float32) {
	if math.IsInf(float64(value), 0) || math.IsNaN(float64(value)) {
		if w.firstError == nil {
			w.firstError = fmt.Errorf("unsupported value: %f", value)
		}
		return
	}
	if value != 0 {
		encodeFloatMember(w, name, value)
	} else {
		encodeZero(w)
	}
}

/*************************************
	float64类型: EncodeDouble_<empty>
 *************************************/

func EncodeDouble(w *JsonEncoder, value float64) {
	if math.IsInf(value, 0) || math.IsNaN(value) {
		if w.firstError == nil {
			w.firstError = fmt.Errorf("unsupported value: %f", value)
		}
		return
	}
	if value != 0 {
		encodeDouble(w, value)
	} else {
		encodeZero(w)
	}
}

func EncodeDouble_OmitEmpty(w *JsonEncoder, name string, value float64) {
	if math.IsInf(value, 0) || math.IsNaN(value) {
		if w.firstError == nil {
			w.firstError = fmt.Errorf("unsupported value: %f", value)
		}
		return
	}
	if value != 0 {
		encodeDoubleMember(w, name, value)
	}
}

func EncodeDouble_WithEmpty(w *JsonEncoder, name string, value float64) {
	if math.IsInf(value, 0) || math.IsNaN(value) {
		if w.firstError == nil {
			w.firstError = fmt.Errorf("unsupported value: %f", value)
		}
		return
	}
	if value != 0 {
		encodeDoubleMember(w, name, value)
	} else {
		encodeZeroMember(w, name)
	}
}

func EncodeDouble_ConvEmpty(w *JsonEncoder, name string, value float64) {
	if math.IsInf(value, 0) || math.IsNaN(value) {
		if w.firstError == nil {
			w.firstError = fmt.Errorf("unsupported value: %f", value)
		}
		return
	}
	if value != 0 {
		encodeDoubleMember(w, name, value)
	} else {
		encodeZeroMember(w, name)
	}
}

/*************************************
	string类型: EncodeString_<escape_html>_<empty>
 *************************************/

func EncodeString(w *JsonEncoder, value string) {
	if value != "" {
		encodeString(w, value)
	} else {
		encodeStringEmpty(w)
	}
}

func EncodeString_OmitEmpty(w *JsonEncoder, name string, value string) {
	if value != "" {
		encodeStringMember(w, name, value)
	}
}

func EncodeString_WithEmpty(w *JsonEncoder, name string, value string) {
	if value != "" {
		encodeStringMember(w, name, value)
	} else {
		encodeStringEmptyMember(w, name)
	}
}

func EncodeString_ConvEmpty(w *JsonEncoder, name string, value string) {
	if value != "" {
		encodeStringMember(w, name, value)
	} else {
		encodeStringEmptyMember(w, name)
	}
}

/*************************************
	bytes类型: WriteBytes_<empty>
 *************************************/

func EncodeBytes(w *JsonEncoder, value []byte) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeStringEmpty(w)
	default:
		encodeBytes(w, value)
	}
}

func EncodeBytes_OmitEmpty(w *JsonEncoder, name string, value []byte) {
	if len(value) != 0 {
		encodeBytesMember(w, name, value)
	}
}

func EncodeBytes_WithEmpty(w *JsonEncoder, name string, value []byte) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeStringEmptyMember(w, name)
	default:
		encodeBytesMember(w, name, value)
	}
}

func EncodeBytes_ConvEmpty(w *JsonEncoder, name string, value []byte) {
	switch {
	case value == nil || len(value) == 0:
		encodeStringEmptyMember(w, name)
	default:
		encodeBytesMember(w, name, value)
	}
}

/*************************************
	enum类型: EncodeEnum_<enum_as_int>_<empty>
 *************************************/

func EncodeEnumName[Enum ~int32](w *JsonEncoder, value Enum, names map[int32]string) {
	encodeString(w, names[int32(value)])
}

func EncodeEnumName_OmitEmpty[Enum ~int32](w *JsonEncoder, name string, value Enum, names map[int32]string) {
	encodeStringMember(w, name, names[int32(value)])
}

func EncodeEnumName_WithEmpty[Enum ~int32](w *JsonEncoder, name string, value Enum, names map[int32]string) {
	encodeStringMember(w, name, names[int32(value)])
}

func EncodeEnumName_ConvEmpty[Enum ~int32](w *JsonEncoder, name string, value Enum, names map[int32]string) {
	encodeStringMember(w, name, names[int32(value)])
}

func EncodeEnum[Enum ~int32](w *JsonEncoder, value Enum) {
	if value != 0 {
		encodeInt32(w, int32(value))
	} else {
		encodeZero(w)
	}
}

func EncodeEnum_OmitEmpty[Enum ~int32](w *JsonEncoder, name string, value Enum) {
	if value != 0 {
		encodeInt32Member(w, name, int32(value))
	}
}

func EncodeEnum_WithEmpty[Enum ~int32](w *JsonEncoder, name string, value Enum) {
	if value != 0 {
		encodeInt32Member(w, name, int32(value))
	} else {
		encodeZeroMember(w, name)
	}
}

func EncodeEnum_ConvEmpty[Enum ~int32](w *JsonEncoder, name string, value Enum) {
	if value != 0 {
		encodeInt32Member(w, name, int32(value))
	} else {
		encodeZeroMember(w, name)
	}
}

/*************************************
	message类型: WriteMessage_<empty>
 *************************************/

func EncodeMessage[Message any](w *JsonEncoder, value *Message) {
	if value != nil {
		encodeObject(w, value)
	} else {
		encodeNull(w)
	}
}

func EncodeMessage_OmitEmpty[Message any](w *JsonEncoder, name string, value *Message) {
	if value != nil {
		encodeObjectMember(w, name, value)
	}
}

func EncodeMessage_WithEmpty[Message any](w *JsonEncoder, name string, value *Message) {
	if value != nil {
		encodeObjectMember(w, name, value)
	} else {
		encodeNullMember(w, name)
	}
}

func EncodeMessage_ConvEmpty[Message any](w *JsonEncoder, name string, value *Message) {
	if value != nil {
		encodeObjectMember(w, name, value)
	} else {
		encodeObjectEmptyMember(w, name)
	}
}
