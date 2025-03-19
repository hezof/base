package base

/*************************************
	bool类型: EncodeBoolRepeated_<empty>
 *************************************/

func EncodeBoolRepeated(w *JsonEncoder, value []bool) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeArrayEmpty(w)
	default:
		encodeArrayWith(w, value, EncodeBool)
	}
}

func EncodeBoolRepeated_OmitEmpty(w *JsonEncoder, name string, value []bool) {
	if len(value) != 0 {
		encodeArrayMemberWith(w, name, value, EncodeBool)
	}
}

func EncodeBoolRepeated_WithEmpty(w *JsonEncoder, name string, value []bool) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		encodeArrayMemberWith(w, name, value, EncodeBool)
	}
}

func EncodeBoolRepeated_ConvEmpty(w *JsonEncoder, name string, value []bool) {
	switch {
	case value == nil || len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		encodeArrayMemberWith(w, name, value, EncodeBool)
	}
}

/*************************************
	int32类型: EncodeInt32Repeated_<empty>
 *************************************/

func EncodeInt32Repeated(w *JsonEncoder, value []int32) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeArrayEmpty(w)
	default:
		encodeArrayWith(w, value, EncodeInt32)
	}
}

func EncodeInt32Repeated_OmitEmpty(w *JsonEncoder, name string, value []int32) {
	if len(value) != 0 {
		encodeArrayMemberWith(w, name, value, EncodeInt32)
	}
}

func EncodeInt32Repeated_WithEmpty(w *JsonEncoder, name string, value []int32) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		encodeArrayMemberWith(w, name, value, EncodeInt32)
	}
}

func EncodeInt32Repeated_ConvEmpty(w *JsonEncoder, name string, value []int32) {
	switch {
	case value == nil || len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		encodeArrayMemberWith(w, name, value, EncodeInt32)
	}
}

/*************************************
	int64类型: EncodeInt64Repeated_<empty>
 *************************************/

func EncodeInt64Repeated(w *JsonEncoder, value []int64) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeArrayEmpty(w)
	default:
		encodeArrayWith(w, value, EncodeInt64)
	}
}

func EncodeInt64Repeated_OmitEmpty(w *JsonEncoder, name string, value []int64) {
	if len(value) != 0 {
		encodeArrayMemberWith(w, name, value, EncodeInt64)
	}
}

func EncodeInt64Repeated_WithEmpty(w *JsonEncoder, name string, value []int64) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		encodeArrayMemberWith(w, name, value, EncodeInt64)
	}
}

func EncodeInt64Repeated_ConvEmpty(w *JsonEncoder, name string, value []int64) {
	switch {
	case value == nil || len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		encodeArrayMemberWith(w, name, value, EncodeInt64)
	}
}

/*************************************
	uint32类型: EncodeUint32Repeated_<empty>
 *************************************/

func EncodeUint32Repeated(w *JsonEncoder, value []uint32) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeArrayEmpty(w)
	default:
		encodeArrayWith(w, value, EncodeUint32)
	}
}

func EncodeUint32Repeated_OmitEmpty(w *JsonEncoder, name string, value []uint32) {
	if len(value) != 0 {
		encodeArrayMemberWith(w, name, value, EncodeUint32)
	}
}

func EncodeUint32Repeated_WithEmpty(w *JsonEncoder, name string, value []uint32) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		encodeArrayMemberWith(w, name, value, EncodeUint32)
	}
}

func EncodeUint32Repeated_ConvEmpty(w *JsonEncoder, name string, value []uint32) {
	switch {
	case value == nil || len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		encodeArrayMemberWith(w, name, value, EncodeUint32)
	}
}

/*************************************
	uin64类型: EncodeUint64Repeated_<empty>
 *************************************/

func EncodeUint64Repeated(w *JsonEncoder, value []uint64) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeArrayEmpty(w)
	default:
		encodeArrayWith(w, value, EncodeUint64)
	}
}

func EncodeUint64Repeated_OmitEmpty(w *JsonEncoder, name string, value []uint64) {
	if len(value) != 0 {
		encodeArrayMemberWith(w, name, value, EncodeUint64)
	}
}

func EncodeUint64Repeated_WithEmpty(w *JsonEncoder, name string, value []uint64) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		encodeArrayMemberWith(w, name, value, EncodeUint64)
	}
}

func EncodeUint64Repeated_ConvEmpty(w *JsonEncoder, name string, value []uint64) {
	switch {
	case value == nil || len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		encodeArrayMemberWith(w, name, value, EncodeUint64)
	}
}

/*************************************
	float32类型: EncodeFloatRepeated_<empty>
 *************************************/

func EncodeFloatRepeated(w *JsonEncoder, value []float32) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeArrayEmpty(w)
	default:
		encodeArrayWith(w, value, EncodeFloat)
	}
}

func EncodeFloatRepeated_OmitEmpty(w *JsonEncoder, name string, value []float32) {
	if len(value) != 0 {
		encodeArrayMemberWith(w, name, value, EncodeFloat)
	}
}

func EncodeFloatRepeated_WithEmpty(w *JsonEncoder, name string, value []float32) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		encodeArrayMemberWith(w, name, value, EncodeFloat)
	}
}

func EncodeFloatRepeated_ConvEmpty(w *JsonEncoder, name string, value []float32) {
	switch {
	case value == nil || len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		encodeArrayMemberWith(w, name, value, EncodeFloat)
	}
}

/*************************************
	float64类型: EncodeDoubleRepeated_<empty>
 *************************************/

func EncodeDoubleRepeated(w *JsonEncoder, value []float64) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeArrayEmpty(w)
	default:
		encodeArrayWith(w, value, EncodeDouble)
	}
}

func EncodeDoubleRepeated_OmitEmpty(w *JsonEncoder, name string, value []float64) {
	if len(value) != 0 {
		encodeArrayMemberWith(w, name, value, EncodeDouble)
	}
}

func EncodeDoubleRepeated_WithEmpty(w *JsonEncoder, name string, value []float64) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		encodeArrayMemberWith(w, name, value, EncodeDouble)
	}
}

func EncodeDoubleRepeated_ConvEmpty(w *JsonEncoder, name string, value []float64) {
	switch {
	case value == nil || len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		encodeArrayMemberWith(w, name, value, EncodeDouble)
	}
}

/*************************************
	string类型: EncodeStringRepeated_<escape_html>_<empty>
 *************************************/

func EncodeStringRepeated(w *JsonEncoder, value []string) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeArrayEmpty(w)
	default:
		encodeArrayWith(w, value, EncodeString)
	}
}

func EncodeStringRepeated_OmitEmpty(w *JsonEncoder, name string, value []string) {
	if len(value) != 0 {
		encodeArrayMemberWith(w, name, value, EncodeString)
	}
}

func EncodeStringRepeated_WithEmpty(w *JsonEncoder, name string, value []string) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		encodeArrayMemberWith(w, name, value, EncodeString)
	}
}

func EncodeStringRepeated_ConvEmpty(w *JsonEncoder, name string, value []string) {
	switch {
	case value == nil || len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		encodeArrayMemberWith(w, name, value, EncodeString)
	}
}

/*************************************
	bytes类型: EncodeBytesRepeated_<empty>
 *************************************/

func EncodeBytesRepeated(w *JsonEncoder, value [][]byte) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeArrayEmpty(w)
	default:
		encodeArrayWith(w, value, EncodeBytes)
	}
}

func EncodeBytesRepeated_OmitEmpty(w *JsonEncoder, name string, value [][]byte) {
	if len(value) != 0 {
		encodeArrayMemberWith(w, name, value, EncodeBytes)
	}
}

func EncodeBytesRepeated_WithEmpty(w *JsonEncoder, name string, value [][]byte) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		encodeArrayMemberWith(w, name, value, EncodeBytes)
	}
}

func EncodeBytesRepeated_ConvEmpty(w *JsonEncoder, name string, value [][]byte) {
	switch {
	case value == nil || len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		encodeArrayMemberWith(w, name, value, EncodeBytes)
	}
}

/*************************************
	enum类型: OptionalEnum_<enum_as_int>_<empty>
 *************************************/

func EncodeEnumNameRepeated[E ~int32](w *JsonEncoder, value []E, names map[int32]string) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeArrayEmpty(w)
	default:
		w.ensure(2)
		w.buff = append(w.buff, leftBracket)
		for _, v := range value {
			encodeString(w, names[int32(v)])
			w.buff = append(w.buff, comma)
		}
		w.buff[len(w.buff)-1] = rightBracket
	}
}

func EncodeEnumNameRepeated_OmitEmpty[E ~int32](w *JsonEncoder, name string, value []E, names map[int32]string) {
	if len(value) != 0 {
		w.ensure(4 + len(name))
		w.buff = append(w.buff, quotes)
		w.buff = append(w.buff, name...)
		w.buff = append(w.buff, quotes, colon, leftBracket)
		for _, v := range value {
			encodeString(w, names[int32(v)])
			w.buff = append(w.buff, comma)
		}
		w.buff[len(w.buff)-1] = rightBracket
		w.buff = append(w.buff, comma)
	}
}

func EncodeEnumNameRepeated_WithEmpty[E ~int32](w *JsonEncoder, name string, value []E, names map[int32]string) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		w.ensure(4 + len(name))
		w.buff = append(w.buff, quotes)
		w.buff = append(w.buff, name...)
		w.buff = append(w.buff, quotes, colon, leftBracket)
		for _, v := range value {
			encodeString(w, names[int32(v)])
			w.buff = append(w.buff, comma)
		}
		w.buff[len(w.buff)-1] = rightBracket
		w.buff = append(w.buff, comma)
	}
}

func EncodeEnumNameRepeated_ConvEmpty[E ~int32](w *JsonEncoder, name string, value []E, names map[int32]string) {
	switch {
	case value == nil || len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		w.ensure(4 + len(name))
		w.buff = append(w.buff, quotes)
		w.buff = append(w.buff, name...)
		w.buff = append(w.buff, quotes, colon, leftBracket)
		for _, v := range value {
			encodeString(w, names[int32(v)])
			w.buff = append(w.buff, comma)
		}
		w.buff[len(w.buff)-1] = rightBracket
		w.buff = append(w.buff, comma)
	}
}

func EncodeEnumRepeated[Enum ~int32](w *JsonEncoder, value []Enum) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeArrayEmpty(w)
	default:
		w.ensure(2)
		w.buff = append(w.buff, leftBracket)
		for _, v := range value {
			EncodeEnum(w, v)
			w.buff = append(w.buff, comma)
		}
		w.buff[len(w.buff)-1] = rightBracket
		w.buff = append(w.buff, comma)
	}
}

func EncodeEnumRepeated_OmitEmpty[Enum ~int32](w *JsonEncoder, name string, value []Enum) {
	if len(value) != 0 {
		w.ensure(4 + len(name))
		w.buff = append(w.buff, quotes)
		w.buff = append(w.buff, name...)
		w.buff = append(w.buff, quotes, colon, leftBracket)
		for _, v := range value {
			EncodeEnum(w, v)
			w.buff = append(w.buff, comma)
		}
		w.buff[len(w.buff)-1] = rightBracket
		w.buff = append(w.buff, comma)
	}
}

func EncodeEnumRepeated_WithEmpty[Enum ~int32](w *JsonEncoder, name string, value []Enum) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		w.ensure(4 + len(name))
		w.buff = append(w.buff, quotes)
		w.buff = append(w.buff, name...)
		w.buff = append(w.buff, quotes, colon, leftBracket)
		for _, v := range value {
			EncodeEnum(w, v)
			w.buff = append(w.buff, comma)
		}
		w.buff[len(w.buff)-1] = rightBracket
		w.buff = append(w.buff, comma)
	}
}

func EncodeEnumRepeated_ConvEmpty[Enum ~int32](w *JsonEncoder, name string, value []Enum) {
	switch {
	case value == nil || len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		w.ensure(4 + len(name))
		w.buff = append(w.buff, quotes)
		w.buff = append(w.buff, name...)
		w.buff = append(w.buff, quotes, colon, leftBracket)
		for _, v := range value {
			EncodeEnum(w, v)
			w.buff = append(w.buff, comma)
		}
		w.buff[len(w.buff)-1] = rightBracket
		w.buff = append(w.buff, comma)
	}
}

/*************************************
	message类型: EncodeMessageRepeated_<empty>
 *************************************/

func EncodeMessageRepeated[Message any](w *JsonEncoder, value []*Message) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeArrayEmpty(w)
	default:
		encodeArrayWith(w, value, EncodeMessage)
	}
}

func EncodeMessageRepeated_OmitEmpty[Message any](w *JsonEncoder, name string, value []*Message) {
	if len(value) != 0 {
		encodeArrayMemberWith(w, name, value, EncodeMessage)
	}
}

func EncodeMessageRepeated_WithEmpty[Message any](w *JsonEncoder, name string, value []*Message) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		encodeArrayMemberWith(w, name, value, EncodeMessage)
	}
}

func EncodeMessageRepeated_ConvEmpty[Message any](w *JsonEncoder, name string, value []*Message) {
	switch {
	case value == nil || len(value) == 0:
		encodeArrayEmptyMember(w, name)
	default:
		encodeArrayMemberWith(w, name, value, EncodeMessage)
	}
}
