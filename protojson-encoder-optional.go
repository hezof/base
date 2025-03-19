package base

/*************************************
	bool类型: EncodeBytes_<empty>
 *************************************/

func EncodeBoolOptional(w *JsonEncoder, value *bool) {
	switch {
	case value == nil:
		encodeNull(w)
	case *value == false:
		encodeFalse(w)
	default:
		encodeTrue(w)
	}
}

func EncodeBoolOptional_OmitEmpty(w *JsonEncoder, name string, value *bool) {
	if value != nil && *value != false {
		encodeTrueMember(w, name)
	}
}

func EncodeBoolOptional_WithEmpty(w *JsonEncoder, name string, value *bool) {

	switch {
	case value == nil:
		encodeNullMember(w, name)
	case *value == false:
		encodeFalseMember(w, name)
	default:
		encodeTrueMember(w, name)
	}
}

func EncodeBoolOptional_ConvEmpty(w *JsonEncoder, name string, value *bool) {
	w.ensure(9 + len(name))
	switch {
	case value == nil || *value == false:
		encodeFalseMember(w, name)
	default:
		encodeTrueMember(w, name)
	}
}

/*************************************
	int32类型: WriteInt32Optional_<empty>
 *************************************/

func EncodeInt32Optional(w *JsonEncoder, value *int32) {
	switch {
	case value == nil:
		encodeNull(w)
	case *value == 0:
		encodeZero(w)
	default:
		encodeInt32(w, *value)
	}
}

func EncodeInt32Optional_OmitEmpty(w *JsonEncoder, name string, value *int32) {
	if value != nil && *value != 0 {
		encodeInt32Member(w, name, *value)
	}
}

func EncodeInt32Optional_WithEmpty(w *JsonEncoder, name string, value *int32) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case *value == 0:
		encodeZeroMember(w, name)
	default:
		encodeInt32Member(w, name, *value)
	}
}

func EncodeInt32Optional_ConvEmpty(w *JsonEncoder, name string, value *int32) {
	switch {
	case value == nil || *value == 0:
		encodeZeroMember(w, name)
	default:
		encodeInt32Member(w, name, *value)
	}
}

/*************************************
	int64类型: WriteInt64Optional_<empty>
 *************************************/

func EncodeInt64Optional(w *JsonEncoder, value *int64) {
	switch {
	case value == nil:
		encodeNull(w)
	case *value == 0:
		encodeZero(w)
	default:
		encodeInt64(w, *value)
	}
}

func EncodeInt64Optional_OmitEmpty(w *JsonEncoder, name string, value *int64) {
	if value != nil && *value != 0 {
		encodeInt64Member(w, name, *value)
	}
}

func EncodeInt64Optional_WithEmpty(w *JsonEncoder, name string, value *int64) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case *value == 0:
		encodeZeroMember(w, name)
	default:
		encodeInt64Member(w, name, *value)
	}
}

func EncodeInt64Optional_ConvEmpty(w *JsonEncoder, name string, value *int64) {
	switch {
	case value == nil || *value == 0:
		encodeZeroMember(w, name)
	default:
		encodeInt64Member(w, name, *value)
	}
}

/*************************************
	uint32类型: WriteUint32Optional_<empty>
 *************************************/

func EncodeUint32Optional(w *JsonEncoder, value *uint32) {
	switch {
	case value == nil:
		encodeNull(w)
	case *value == 0:
		encodeZero(w)
	default:
		encodeUint32(w, *value)
	}
}

func EncodeUint32Optional_OmitEmpty(w *JsonEncoder, name string, value *uint32) {
	if value != nil && *value != 0 {
		encodeUint32Member(w, name, *value)
	}
}

func EncodeUint32Optional_WithEmpty(w *JsonEncoder, name string, value *uint32) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case *value == 0:
		encodeZeroMember(w, name)
	default:
		encodeUint32Member(w, name, *value)
	}
}

func EncodeUint32Optional_ConvEmpty(w *JsonEncoder, name string, value *uint32) {
	switch {
	case value == nil || *value == 0:
		encodeZeroMember(w, name)
	default:
		encodeUint32Member(w, name, *value)
	}
}

/*************************************
	uin64类型: WriteUint64Optional_<empty>
 *************************************/

func EncodeUint64Optional(w *JsonEncoder, value *uint64) {
	switch {
	case value == nil:
		encodeNull(w)
	case *value == 0:
		encodeZero(w)
	default:
		encodeUint64(w, *value)
	}
}

func EncodeUint64Optional_OmitEmpty(w *JsonEncoder, name string, value *uint64) {
	if value != nil && *value != 0 {
		encodeUint64Member(w, name, *value)
	}
}

func EncodeUint64Optional_WithEmpty(w *JsonEncoder, name string, value *uint64) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case *value == 0:
		encodeZeroMember(w, name)
	default:
		encodeUint64Member(w, name, *value)
	}
}

func EncodeUint64Optional_ConvEmpty(w *JsonEncoder, name string, value *uint64) {
	switch {
	case value == nil || *value == 0:
		encodeZeroMember(w, name)
	default:
		encodeUint64Member(w, name, *value)
	}
}

/*************************************
	float32类型: WriteFloatOptional_<empty>
 *************************************/

func EncodeFloatOptional(w *JsonEncoder, value *float32) {
	switch {
	case value == nil:
		encodeNull(w)
	case *value == 0:
		encodeZero(w)
	default:
		encodeFloat(w, *value)
	}
}

func EncodeFloatOptional_OmitEmpty(w *JsonEncoder, name string, value *float32) {
	if value != nil && *value != 0 {
		encodeFloatMember(w, name, *value)
	}
}

func EncodeFloatOptional_WithEmpty(w *JsonEncoder, name string, value *float32) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case *value == 0:
		encodeZeroMember(w, name)
	default:
		encodeFloatMember(w, name, *value)
	}
}

func EncodeFloatOptional_ConvEmpty(w *JsonEncoder, name string, value *float32) {
	switch {
	case value == nil || *value == 0:
		encodeZeroMember(w, name)
	default:
		encodeFloatMember(w, name, *value)
	}
}

/*************************************
	float64类型: WriteDoubleOptional_<empty>
 *************************************/

func EncodeDoubleOptional(w *JsonEncoder, value *float64) {
	switch {
	case value == nil:
		encodeNull(w)
	case *value == 0:
		encodeZero(w)
	default:
		encodeDouble(w, *value)
	}
}

func EncodeDoubleOptional_OmitEmpty(w *JsonEncoder, name string, value *float64) {
	if value != nil && *value != 0 {
		encodeDoubleMember(w, name, *value)
	}
}

func EncodeDoubleOptional_WithEmpty(w *JsonEncoder, name string, value *float64) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case *value == 0:
		encodeZeroMember(w, name)
	default:
		encodeDoubleMember(w, name, *value)
	}
}

func EncodeDoubleOptional_ConvEmpty(w *JsonEncoder, name string, value *float64) {
	switch {
	case value == nil || *value == 0:
		encodeZeroMember(w, name)
	default:
		encodeDoubleMember(w, name, *value)
	}
}

/*************************************
	string类型: WriteStringOptional_<escape_html>_<empty>
 *************************************/

func EncodeStringOptional(w *JsonEncoder, value *string) {
	switch {
	case value == nil:
		encodeNull(w)
	case *value == "":
		encodeStringEmpty(w)
	default:
		encodeString(w, *value)
	}
}

func EncodeStringOptional_OmitEmpty(w *JsonEncoder, name string, value *string) {
	if value != nil && *value != "" {
		encodeStringMember(w, name, *value)
	}
}

func EncodeStringOptional_WithEmpty(w *JsonEncoder, name string, value *string) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case *value == "":
		encodeStringEmptyMember(w, name)
	default:
		encodeStringMember(w, name, *value)
	}
}

func EncodeStringOptional_ConvEmpty(w *JsonEncoder, name string, value *string) {
	switch {
	case value == nil || *value == "":
		encodeStringEmptyMember(w, name)
	default:
		encodeStringMember(w, name, *value)
	}
}

/*************************************
	bytes类型: optional类型相同
 *************************************/

func EncodeBytesOptional(w *JsonEncoder, value []byte) {
	EncodeBytes(w, value)
}

func EncodeBytesOptional_OmitEmpty(w *JsonEncoder, name string, value []byte) {
	EncodeBytes_OmitEmpty(w, name, value)
}

func EncodeBytesOptional_WithEmpty(w *JsonEncoder, name string, value []byte) {
	EncodeBytes_WithEmpty(w, name, value)
}

func EncodeBytesOptional_ConvEmpty(w *JsonEncoder, name string, value []byte) {
	EncodeBytes_ConvEmpty(w, name, value)
}

/*************************************
	enum类型: OptionalEnum_<enum_as_int>_<empty>
 *************************************/

func EncodeEnumNameOptional[Enum ~int32](w *JsonEncoder, value *Enum, names map[int32]string) {
	if value != nil {
		encodeString(w, names[int32(*value)])
	} else {
		encodeNull(w)
	}
}

func EncodeEnumNameOptional_OmitEmpty[Enum ~int32](w *JsonEncoder, name string, value *Enum, names map[int32]string) {
	if value != nil {
		encodeStringMember(w, name, names[int32(*value)])
	}
}

func EncodeEnumNameOptional_WithEmpty[Enum ~int32](w *JsonEncoder, name string, value *Enum, names map[int32]string) {
	if value != nil {
		encodeStringMember(w, name, names[int32(*value)])
	} else {
		encodeNullMember(w, name)
	}
}

func EncodeEnumNameOptional_ConvEmpty[Enum ~int32](w *JsonEncoder, name string, value *Enum, names map[int32]string) {
	if value != nil {
		encodeStringMember(w, name, names[int32(*value)])
	} else {
		encodeStringMember(w, name, names[0])
	}
}

func EncodeEnumOptional[Enum ~int32](w *JsonEncoder, value *Enum) {
	if value != nil {
		EncodeInt32(w, int32(*value))
	} else {
		encodeNull(w)
	}
}

func EncodeEnumOptional_OmitEmpty[Enum ~int32](w *JsonEncoder, name string, value *Enum) {
	if value != nil {
		encodeInt32Member(w, name, int32(*value))
	}
}

func EncodeEnumOptional_WithEmpty[Enum ~int32](w *JsonEncoder, name string, value *Enum) {
	if value != nil {
		encodeInt32Member(w, name, int32(*value))
	} else {
		encodeNullMember(w, name)
	}
}

func EncodeEnumOptional_ConvEmpty[Enum ~int32](w *JsonEncoder, name string, value *Enum) {
	if value != nil {
		encodeInt32Member(w, name, int32(*value))
	} else {
		encodeZeroMember(w, name)
	}
}

/*************************************
	message类型
 *************************************/

func EncodeMessageOptional[Message any](w *JsonEncoder, value *Message) {
	EncodeMessage(w, value)
}

func EncodeMessageOptional_OmitEmpty[Message any](w *JsonEncoder, name string, value *Message) {
	EncodeMessage_OmitEmpty(w, name, value)
}

func EncodeMessageOptional_WithEmpty[Message any](w *JsonEncoder, name string, value *Message) {
	EncodeMessage_WithEmpty(w, name, value)
}

func EncodeMessageOptional_ConvEmpty[Message any](w *JsonEncoder, name string, value *Message) {
	EncodeMessage_ConvEmpty(w, name, value)
}
