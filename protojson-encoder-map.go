package framework

/*************************************
	bool类型: MapBoolMap_<empty>
 *************************************/

func EncodeBoolMap(w *JsonEncoder, value map[string]bool) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeObjectEmpty(w)
	default:
		encodeObjectWith(w, value, EncodeBool)
	}
}

func EncodeBoolMap_OmitEmpty(w *JsonEncoder, name string, value map[string]bool) {
	if len(value) != 0 {
		encodeObjectMemberWith(w, name, value, EncodeBool)
	}
}

func EncodeBoolMap_WithEmpty(w *JsonEncoder, name string, value map[string]bool) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		encodeObjectMemberWith(w, name, value, EncodeBool)
	}
}

func EncodeBoolMap_ConvEmpty(w *JsonEncoder, name string, value map[string]bool) {
	switch {
	case value == nil || len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		encodeObjectMemberWith(w, name, value, EncodeBool)
	}
}

/*************************************
	int32类型: MapInt32Map_<empty>
 *************************************/

func EncodeInt32Map(w *JsonEncoder, value map[string]int32) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeObjectEmpty(w)
	default:
		encodeObjectWith(w, value, EncodeInt32)
	}
}

func EncodeInt32Map_OmitEmpty(w *JsonEncoder, name string, value map[string]int32) {
	if len(value) != 0 {
		encodeObjectMemberWith(w, name, value, EncodeInt32)
	}
}

func EncodeInt32Map_WithEmpty(w *JsonEncoder, name string, value map[string]int32) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		encodeObjectMemberWith(w, name, value, EncodeInt32)
	}
}

func EncodeInt32Map_ConvEmpty(w *JsonEncoder, name string, value map[string]int32) {
	switch {
	case value == nil || len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		encodeObjectMemberWith(w, name, value, EncodeInt32)
	}
}

/*************************************
	int64类型: MapInt64Map_<empty>
 *************************************/

func EncodeInt64Map(w *JsonEncoder, value map[string]int64) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeObjectEmpty(w)
	default:
		encodeObjectWith(w, value, EncodeInt64)
	}
}

func EncodeInt64Map_OmitEmpty(w *JsonEncoder, name string, value map[string]int64) {
	if len(value) != 0 {
		encodeObjectMemberWith(w, name, value, EncodeInt64)
	}
}

func EncodeInt64Map_WithEmpty(w *JsonEncoder, name string, value map[string]int64) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		encodeObjectMemberWith(w, name, value, EncodeInt64)
	}
}

func EncodeInt64Map_ConvEmpty(w *JsonEncoder, name string, value map[string]int64) {
	switch {
	case value == nil || len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		encodeObjectMemberWith(w, name, value, EncodeInt64)
	}
}

/*************************************
	uint32类型: MapUint32Map_<empty>
 *************************************/

func EncodeUint32Map(w *JsonEncoder, value map[string]uint32) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeObjectEmpty(w)
	default:
		encodeObjectWith(w, value, EncodeUint32)
	}
}

func EncodeUint32Map_OmitEmpty(w *JsonEncoder, name string, value map[string]uint32) {
	if len(value) != 0 {
		encodeObjectMemberWith(w, name, value, EncodeUint32)
	}
}

func EncodeUint32Map_WithEmpty(w *JsonEncoder, name string, value map[string]uint32) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		encodeObjectMemberWith(w, name, value, EncodeUint32)
	}
}

func EncodeUint32Map_ConvEmpty(w *JsonEncoder, name string, value map[string]uint32) {
	switch {
	case value == nil || len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		encodeObjectMemberWith(w, name, value, EncodeUint32)
	}
}

/*************************************
	uin64类型: MapUint64Map_<empty>
 *************************************/

func EncodeUint64Map(w *JsonEncoder, value map[string]uint64) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeObjectEmpty(w)
	default:
		encodeObjectWith(w, value, EncodeUint64)
	}
}

func EncodeUint64Map_OmitEmpty(w *JsonEncoder, name string, value map[string]uint64) {
	if len(value) != 0 {
		encodeObjectMemberWith(w, name, value, EncodeUint64)
	}
}

func EncodeUint64Map_WithEmpty(w *JsonEncoder, name string, value map[string]uint64) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		encodeObjectMemberWith(w, name, value, EncodeUint64)
	}
}

func EncodeUint64Map_ConvEmpty(w *JsonEncoder, name string, value map[string]uint64) {
	switch {
	case value == nil || len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		encodeObjectMemberWith(w, name, value, EncodeUint64)
	}
}

/*************************************
	float32类型: MapFloatMap_<empty>
 *************************************/

func EncodeFloatMap(w *JsonEncoder, value map[string]float32) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeObjectEmpty(w)
	default:
		encodeObjectWith(w, value, EncodeFloat)
	}
}

func EncodeFloatMap_OmitEmpty(w *JsonEncoder, name string, value map[string]float32) {
	if len(value) != 0 {
		encodeObjectMemberWith(w, name, value, EncodeFloat)
	}
}

func EncodeFloatMap_WithEmpty(w *JsonEncoder, name string, value map[string]float32) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		encodeObjectMemberWith(w, name, value, EncodeFloat)
	}
}

func EncodeFloatMap_ConvEmpty(w *JsonEncoder, name string, value map[string]float32) {
	switch {
	case value == nil || len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		encodeObjectMemberWith(w, name, value, EncodeFloat)
	}
}

/*************************************
	float64类型: MapDoubleMap_<empty>
 *************************************/

func EncodeDoubleMap(w *JsonEncoder, value map[string]float64) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeObjectEmpty(w)
	default:
		encodeObjectWith(w, value, EncodeDouble)
	}
}

func EncodeDoubleMap_OmitEmpty(w *JsonEncoder, name string, value map[string]float64) {
	if len(value) != 0 {
		encodeObjectMemberWith(w, name, value, EncodeDouble)
	}
}

func EncodeDoubleMap_WithEmpty(w *JsonEncoder, name string, value map[string]float64) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		encodeObjectMemberWith(w, name, value, EncodeDouble)
	}
}

func EncodeDoubleMap_ConvEmpty(w *JsonEncoder, name string, value map[string]float64) {
	switch {
	case value == nil || len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		encodeObjectMemberWith(w, name, value, EncodeDouble)
	}
}

/*************************************
	string类型: MapStringMap_<escap_html>_<empty>
 *************************************/

func EncodeStringMap(w *JsonEncoder, value map[string]string) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeObjectEmpty(w)
	default:
		encodeObjectWith(w, value, EncodeString)
	}
}

func EncodeStringMap_OmitEmpty(w *JsonEncoder, name string, value map[string]string) {
	if len(value) != 0 {
		encodeObjectMemberWith(w, name, value, EncodeString)
	}
}

func EncodeStringMap_WithEmpty(w *JsonEncoder, name string, value map[string]string) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		encodeObjectMemberWith(w, name, value, EncodeString)
	}
}

func EncodeStringMap_ConvEmpty(w *JsonEncoder, name string, value map[string]string) {
	switch {
	case value == nil || len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		encodeObjectMemberWith(w, name, value, EncodeString)
	}
}

/*************************************
	bytes类型: MapBytesMap_<empty>
 *************************************/

func EncodeBytesMap(w *JsonEncoder, value map[string][]byte) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeObjectEmpty(w)
	default:
		encodeObjectWith(w, value, EncodeBytes)
	}
}

func EncodeBytesMap_OmitEmpty(w *JsonEncoder, name string, value map[string][]byte) {
	if len(value) != 0 {
		encodeObjectMemberWith(w, name, value, EncodeBytes)
	}
}

func EncodeBytesMap_WithEmpty(w *JsonEncoder, name string, value map[string][]byte) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		encodeObjectMemberWith(w, name, value, EncodeBytes)
	}
}

func EncodeBytesMap_ConvEmpty(w *JsonEncoder, name string, value map[string][]byte) {
	switch {
	case value == nil || len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		encodeObjectMemberWith(w, name, value, EncodeBytes)
	}
}

/*************************************
	enum类型: MapEnumMap_<enum_as_int>_<empty>
 *************************************/

func EncodeEnumNameMap[Enum ~int32](w *JsonEncoder, value map[string]Enum, names map[int32]string) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeObjectEmpty(w)
	default:
		w.ensure(2)
		w.buff = append(w.buff, leftBrace)
		for k, v := range value {
			encodeString(w, k)
			w.buff = append(w.buff, colon)
			encodeString(w, names[int32(v)])
			w.buff = append(w.buff, comma)
		}
		w.buff[len(w.buff)-1] = rightBrace
	}
}

func EncodeEnumNameMap_OmitEmpty[Enum ~int32](w *JsonEncoder, name string, value map[string]Enum, names map[int32]string) {
	if len(value) != 0 {
		w.ensure(4 + len(name))
		w.buff = append(w.buff, quotes)
		w.buff = append(w.buff, name...)
		w.buff = append(w.buff, quotes, colon, leftBrace)
		for k, v := range value {
			encodeString(w, k)
			w.buff = append(w.buff, colon)
			encodeString(w, names[int32(v)])
			w.buff = append(w.buff, comma)
		}
		w.buff[len(w.buff)-1] = rightBrace
		w.buff = append(w.buff, comma)
	}
}

func EncodeEnumNameMap_WithEmpty[Enum ~int32](w *JsonEncoder, name string, value map[string]Enum, names map[int32]string) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		w.ensure(4 + len(name))
		w.buff = append(w.buff, quotes)
		w.buff = append(w.buff, name...)
		w.buff = append(w.buff, quotes, colon, leftBrace)
		for k, v := range value {
			encodeString(w, k)
			w.buff = append(w.buff, colon)
			encodeString(w, names[int32(v)])
			w.buff = append(w.buff, comma)
		}
		w.buff[len(w.buff)-1] = rightBrace
		w.buff = append(w.buff, comma)
	}
}

func EncodeEnumNameMap_ConvEmpty[Enum ~int32](w *JsonEncoder, name string, value map[string]Enum, names map[int32]string) {
	switch {
	case value == nil || len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		w.ensure(4 + len(name))
		w.buff = append(w.buff, quotes)
		w.buff = append(w.buff, name...)
		w.buff = append(w.buff, quotes, colon, leftBrace)
		for k, v := range value {
			encodeString(w, k)
			w.buff = append(w.buff, colon)
			encodeString(w, names[int32(v)])
			w.buff = append(w.buff, comma)
		}
		w.buff[len(w.buff)-1] = rightBrace
		w.buff = append(w.buff, comma)
	}
}

func EncodeEnumMap[Enum ~int32](w *JsonEncoder, value map[string]Enum) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeObjectEmpty(w)
	default:
		w.ensure(2)
		w.buff = append(w.buff, leftBrace)
		for k, v := range value {
			encodeString(w, k)
			w.buff = append(w.buff, colon)
			EncodeInt32(w, int32(v))
			w.buff = append(w.buff, comma)
		}
		w.buff[len(w.buff)-1] = rightBrace
	}
}

func EncodeEnumMap_OmitEmpty[Enum ~int32](w *JsonEncoder, name string, value map[string]Enum) {
	if len(value) != 0 {
		w.ensure(4 + len(name))
		w.buff = append(w.buff, quotes)
		w.buff = append(w.buff, name...)
		w.buff = append(w.buff, quotes, colon, leftBrace)
		for k, v := range value {
			encodeString(w, k)
			w.buff = append(w.buff, colon)
			EncodeInt32(w, int32(v))
			w.buff = append(w.buff, comma)
		}
		w.buff[len(w.buff)-1] = rightBrace
		w.buff = append(w.buff, comma)
	}
}

func EncodeEnumMap_WithEmpty[Enum ~int32](w *JsonEncoder, name string, value map[string]Enum) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		w.ensure(4 + len(name))
		w.buff = append(w.buff, quotes)
		w.buff = append(w.buff, name...)
		w.buff = append(w.buff, quotes, colon, leftBrace)
		for k, v := range value {
			encodeString(w, k)
			w.buff = append(w.buff, colon)
			EncodeInt32(w, int32(v))
			w.buff = append(w.buff, comma)
		}
		w.buff[len(w.buff)-1] = rightBrace
		w.buff = append(w.buff, comma)
	}
}

func EncodeEnumMap_ConvEmpty[Enum ~int32](w *JsonEncoder, name string, value map[string]Enum) {
	switch {
	case value == nil || len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		w.ensure(4 + len(name))
		w.buff = append(w.buff, quotes)
		w.buff = append(w.buff, name...)
		w.buff = append(w.buff, quotes, colon, leftBrace)
		for k, v := range value {
			encodeString(w, k)
			w.buff = append(w.buff, colon)
			EncodeInt32(w, int32(v))
			w.buff = append(w.buff, comma)
		}
		w.buff[len(w.buff)-1] = rightBrace
		w.buff = append(w.buff, comma)
	}
}

/*************************************
	message类型: MapMessageMap_<empty>
 *************************************/

func EncodeMessageMap[Message any](w *JsonEncoder, value map[string]*Message) {
	switch {
	case value == nil:
		encodeNull(w)
	case len(value) == 0:
		encodeObjectEmpty(w)
	default:
		encodeObjectWith(w, value, EncodeMessage)
	}
}

func EncodeMessageMap_OmitEmpty[Message any](w *JsonEncoder, name string, value map[string]*Message) {
	if len(value) != 0 {
		encodeObjectMemberWith(w, name, value, EncodeMessage)
	}
}

func EncodeMessageMap_WithEmpty[Message any](w *JsonEncoder, name string, value map[string]*Message) {
	switch {
	case value == nil:
		encodeNullMember(w, name)
	case len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		encodeObjectMemberWith(w, name, value, EncodeMessage)
	}
}

func EncodeMessageMap_ConvEmpty[Message any](w *JsonEncoder, name string, value map[string]*Message) {
	switch {
	case value == nil || len(value) == 0:
		encodeObjectEmptyMember(w, name)
	default:
		encodeObjectMemberWith(w, name, value, EncodeMessage)
	}
}
