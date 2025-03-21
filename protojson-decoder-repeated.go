package framework

func DecodeBoolRepeated(r *JsonDecoder, p *[]bool) {
	switch r.token {
	case ArrayBegin:
		if *p == nil {
			*p = make([]bool, 0)
		} else {
			*p = (*p)[0:0]
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		if r.next() == ArrayEnd {
			return
		}
		for {
			var v bool
			DecodeBool(r, &v)
			*p = append(*p, v)

			switch r.next() {
			case Comma:
				if r.next() == ArrayEnd {
					r.invalidCharacterError()
					return
				}
			case ArrayEnd:
				return
			default:
				r.invalidCharacterError()
				return
			}
		}
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(ArrayBegin)
	}
}

func DecodeInt32Repeated(r *JsonDecoder, p *[]int32) {
	switch r.token {
	case ArrayBegin:
		if *p == nil {
			*p = make([]int32, 0)
		} else {
			*p = (*p)[0:0]
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		if r.next() == ArrayEnd {
			return
		}
		for {
			var v int32
			DecodeInt32(r, &v)
			*p = append(*p, v)

			switch r.next() {
			case Comma:
				if r.next() == ArrayEnd {
					r.invalidCharacterError()
					return
				}
			case ArrayEnd:
				return
			default:
				r.invalidCharacterError()
				return
			}
		}
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(ArrayBegin)
	}
}

func DecodeInt64Repeated(r *JsonDecoder, p *[]int64) {
	switch r.token {
	case ArrayBegin:
		if *p == nil {
			*p = make([]int64, 0)
		} else {
			*p = (*p)[0:0]
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		if r.next() == ArrayEnd {
			return
		}
		for {
			var v int64
			DecodeInt64(r, &v)
			*p = append(*p, v)

			switch r.next() {
			case Comma:
				if r.next() == ArrayEnd {
					r.invalidCharacterError()
					return
				}
			case ArrayEnd:
				return
			default:
				r.invalidCharacterError()
				return
			}
		}
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(ArrayBegin)
	}
}

func DecodeUint32Repeated(r *JsonDecoder, p *[]uint32) {
	switch r.token {
	case ArrayBegin:
		if *p == nil {
			*p = make([]uint32, 0)
		} else {
			*p = (*p)[0:0]
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		if r.next() == ArrayEnd {
			return
		}
		for {
			var v uint32
			DecodeUint32(r, &v)
			*p = append(*p, v)

			switch r.next() {
			case Comma:
				if r.next() == ArrayEnd {
					r.invalidCharacterError()
					return
				}
			case ArrayEnd:
				return
			default:
				r.invalidCharacterError()
				return
			}
		}
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(ArrayBegin)
	}
}

func DecodeUint64Repeated(r *JsonDecoder, p *[]uint64) {
	switch r.token {
	case ArrayBegin:
		if *p == nil {
			*p = make([]uint64, 0)
		} else {
			*p = (*p)[0:0]
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		if r.next() == ArrayEnd {
			return
		}
		for {
			var v uint64
			DecodeUint64(r, &v)
			*p = append(*p, v)

			switch r.next() {
			case Comma:
				if r.next() == ArrayEnd {
					r.invalidCharacterError()
					return
				}
			case ArrayEnd:
				return
			default:
				r.invalidCharacterError()
				return
			}
		}
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(ArrayBegin)
	}
}

func DecodeFloatRepeated(r *JsonDecoder, p *[]float32) {
	switch r.token {
	case ArrayBegin:
		if *p == nil {
			*p = make([]float32, 0)
		} else {
			*p = (*p)[0:0]
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		if r.next() == ArrayEnd {
			return
		}
		for {
			var v float32
			DecodeFloat(r, &v)
			*p = append(*p, v)

			switch r.next() {
			case Comma:
				if r.next() == ArrayEnd {
					r.invalidCharacterError()
					return
				}
			case ArrayEnd:
				return
			default:
				r.invalidCharacterError()
				return
			}
		}
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(ArrayBegin)
	}
}

func DecodeDoubleRepeated(r *JsonDecoder, p *[]float64) {
	switch r.token {
	case ArrayBegin:
		if *p == nil {
			*p = make([]float64, 0)
		} else {
			*p = (*p)[0:0]
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		if r.next() == ArrayEnd {
			return
		}
		for {
			var v float64
			DecodeDouble(r, &v)
			*p = append(*p, v)

			switch r.next() {
			case Comma:
				if r.next() == ArrayEnd {
					r.invalidCharacterError()
					return
				}
			case ArrayEnd:
				return
			default:
				r.invalidCharacterError()
				return
			}
		}
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(ArrayBegin)
	}
}

func DecodeStringRepeated(r *JsonDecoder, p *[]string) {
	switch r.token {
	case ArrayBegin:
		if *p == nil {
			*p = make([]string, 0)
		} else {
			*p = (*p)[0:0]
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		if r.next() == ArrayEnd {
			return
		}
		for {
			var v string
			DecodeString(r, &v)
			*p = append(*p, v)

			switch r.next() {
			case Comma:
				if r.next() == ArrayEnd {
					r.invalidCharacterError()
					return
				}
			case ArrayEnd:
				return
			default:
				r.invalidCharacterError()
				return
			}
		}
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(ArrayBegin)
	}
}

func DecodeBytesRepeated(r *JsonDecoder, p *[][]byte) {
	switch r.token {
	case ArrayBegin:
		if *p == nil {
			*p = make([][]byte, 0)
		} else {
			*p = (*p)[0:0]
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		if r.next() == ArrayEnd {
			return
		}
		for {
			var v []byte
			DecodeBytes(r, &v)
			*p = append(*p, v)

			switch r.next() {
			case Comma:
				if r.next() == ArrayEnd {
					r.invalidCharacterError()
					return
				}
			case ArrayEnd:
				return
			default:
				r.invalidCharacterError()
				return
			}
		}
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(ArrayBegin)
	}
}

func DecodeEnumNameRepeated[Enum ~int32](r *JsonDecoder, p *[]Enum, values map[string]int32) {
	switch r.token {
	case ArrayBegin:
		if *p == nil {
			*p = make([]Enum, 0)
		} else {
			*p = (*p)[0:0]
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		if r.next() == ArrayEnd {
			return
		}
		for {
			var v Enum
			DecodeEnumName(r, &v, values)
			*p = append(*p, v)

			switch r.next() {
			case Comma:
				if r.next() == ArrayEnd {
					r.invalidCharacterError()
					return
				}
			case ArrayEnd:
				return
			default:
				r.invalidCharacterError()
				return
			}
		}
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(ArrayBegin)
	}
}

func DecodeEnumRepeated[Enum ~int32](r *JsonDecoder, p *[]Enum) {
	switch r.token {
	case ArrayBegin:
		if *p == nil {
			*p = make([]Enum, 0)
		} else {
			*p = (*p)[0:0]
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		if r.next() == ArrayEnd {
			return
		}
		for {
			var v Enum
			DecodeEnum(r, &v)
			*p = append(*p, v)

			switch r.next() {
			case Comma:
				if r.next() == ArrayEnd {
					r.invalidCharacterError()
					return
				}
			case ArrayEnd:
				return
			default:
				r.invalidCharacterError()
				return
			}
		}
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(ArrayBegin)
	}
}

func DecodeMessageRepeated[Message any](r *JsonDecoder, p *[]*Message) {
	switch r.token {
	case ArrayBegin:
		if *p == nil {
			*p = make([]*Message, 0)
		} else {
			*p = (*p)[0:0]
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		if r.next() == ArrayEnd {
			return
		}
		for {
			var v *Message
			DecodeMessage(r, &v)
			*p = append(*p, v)

			switch r.next() {
			case Comma:
				if r.next() == ArrayEnd {
					r.invalidCharacterError()
					return
				}
			case ArrayEnd:
				return
			default:
				r.invalidCharacterError()
				return
			}
		}
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(ArrayBegin)
	}
}
