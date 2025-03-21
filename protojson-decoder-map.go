package framework

func DecodeBoolMap(r *JsonDecoder, p *map[string]bool) {
	switch r.token {
	case ObjectBegin:
		if *p == nil {
			*p = make(map[string]bool)
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		t := r.next()
		if t == ObjectEnd {
			return
		}
		for {
			if t != String {
				r.expectedTokenError(String)
				return
			}
			k := r.readString()
			if r.next() != Colon {
				r.expectedTokenError(Colon)
				return
			}
			r.next()

			var v bool
			DecodeBool(r, &v)
			(*p)[k] = v

			t = r.next()
			switch t {
			case Comma:
				t = r.next()
				if t == ObjectEnd {
					r.invalidCharacterError()
					return
				}
			case ObjectEnd:
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
		r.expectedTokenError(ObjectBegin)
	}
}

func DecodeInt32Map(r *JsonDecoder, p *map[string]int32) {
	switch r.token {
	case ObjectBegin:
		if *p == nil {
			*p = make(map[string]int32)
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		t := r.next()
		if t == ObjectEnd {
			return
		}
		for {
			if t != String {
				r.expectedTokenError(String)
				return
			}
			k := r.readString()
			if r.next() != Colon {
				r.expectedTokenError(Colon)
				return
			}
			r.next()

			var v int32
			DecodeInt32(r, &v)
			(*p)[k] = v

			t = r.next()
			switch t {
			case Comma:
				t = r.next()
				if t == ObjectEnd {
					r.invalidCharacterError()
					return
				}
			case ObjectEnd:
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
		r.expectedTokenError(ObjectBegin)
	}
}

func DecodeInt64Map(r *JsonDecoder, p *map[string]int64) {
	switch r.token {
	case ObjectBegin:
		if *p == nil {
			*p = make(map[string]int64)
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		t := r.next()
		if t == ObjectEnd {
			return
		}
		for {
			if t != String {
				r.expectedTokenError(String)
				return
			}
			k := r.readString()
			if r.next() != Colon {
				r.expectedTokenError(Colon)
				return
			}
			r.next()

			var v int64
			DecodeInt64(r, &v)
			(*p)[k] = v

			t = r.next()
			switch t {
			case Comma:
				t = r.next()
				if t == ObjectEnd {
					r.invalidCharacterError()
					return
				}
			case ObjectEnd:
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
		r.expectedTokenError(ObjectBegin)
	}
}

func DecodeUint32Map(r *JsonDecoder, p *map[string]uint32) {
	switch r.token {
	case ObjectBegin:
		if *p == nil {
			*p = make(map[string]uint32)
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		t := r.next()
		if t == ObjectEnd {
			return
		}
		for {
			if t != String {
				r.expectedTokenError(String)
				return
			}
			k := r.readString()
			if r.next() != Colon {
				r.expectedTokenError(Colon)
				return
			}
			r.next()

			var v uint32
			DecodeUint32(r, &v)
			(*p)[k] = v

			t = r.next()
			switch t {
			case Comma:
				t = r.next()
				if t == ObjectEnd {
					r.invalidCharacterError()
					return
				}
			case ObjectEnd:
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
		r.expectedTokenError(ObjectBegin)
	}
}

func DecodeUint64Map(r *JsonDecoder, p *map[string]uint64) {
	switch r.token {
	case ObjectBegin:
		if *p == nil {
			*p = make(map[string]uint64)
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		t := r.next()
		if t == ObjectEnd {
			return
		}
		for {
			if t != String {
				r.expectedTokenError(String)
				return
			}
			k := r.readString()
			if r.next() != Colon {
				r.expectedTokenError(Colon)
				return
			}
			r.next()

			var v uint64
			DecodeUint64(r, &v)
			(*p)[k] = v

			t = r.next()
			switch t {
			case Comma:
				t = r.next()
				if t == ObjectEnd {
					r.invalidCharacterError()
					return
				}
			case ObjectEnd:
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
		r.expectedTokenError(ObjectBegin)
	}
}

func DecodeFloatMap(r *JsonDecoder, p *map[string]float32) {
	switch r.token {
	case ObjectBegin:
		if *p == nil {
			*p = make(map[string]float32)
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		t := r.next()
		if t == ObjectEnd {
			return
		}
		for {
			if t != String {
				r.expectedTokenError(String)
				return
			}
			k := r.readString()
			if r.next() != Colon {
				r.expectedTokenError(Colon)
				return
			}
			r.next()

			var v float32
			DecodeFloat(r, &v)
			(*p)[k] = v

			t = r.next()
			switch t {
			case Comma:
				t = r.next()
				if t == ObjectEnd {
					r.invalidCharacterError()
					return
				}
			case ObjectEnd:
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
		r.expectedTokenError(ObjectBegin)
	}
}

func DecodeDoubleMap(r *JsonDecoder, p *map[string]float64) {
	switch r.token {
	case ObjectBegin:
		if *p == nil {
			*p = make(map[string]float64)
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		t := r.next()
		if t == ObjectEnd {
			return
		}
		for {
			if t != String {
				r.expectedTokenError(String)
				return
			}
			k := r.readString()
			if r.next() != Colon {
				r.expectedTokenError(Colon)
				return
			}
			r.next()

			var v float64
			DecodeDouble(r, &v)
			(*p)[k] = v

			t = r.next()
			switch t {
			case Comma:
				t = r.next()
				if t == ObjectEnd {
					r.invalidCharacterError()
					return
				}
			case ObjectEnd:
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
		r.expectedTokenError(ObjectBegin)
	}
}

func DecodeStringMap(r *JsonDecoder, p *map[string]string) {
	switch r.token {
	case ObjectBegin:
		if *p == nil {
			*p = make(map[string]string)
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		t := r.next()
		if t == ObjectEnd {
			return
		}
		for {
			if t != String {
				r.expectedTokenError(String)
				return
			}
			k := r.readString()
			if r.next() != Colon {
				r.expectedTokenError(Colon)
				return
			}
			r.next()

			var v string
			DecodeString(r, &v)
			(*p)[k] = v

			t = r.next()
			switch t {
			case Comma:
				t = r.next()
				if t == ObjectEnd {
					r.invalidCharacterError()
					return
				}
			case ObjectEnd:
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
		r.expectedTokenError(ObjectBegin)
	}
}

func DecodeBytesMap(r *JsonDecoder, p *map[string][]byte) {
	switch r.token {
	case ObjectBegin:
		if *p == nil {
			*p = make(map[string][]byte)
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		t := r.next()
		if t == ObjectEnd {
			return
		}
		for {
			if t != String {
				r.expectedTokenError(String)
				return
			}
			k := r.readString()
			if r.next() != Colon {
				r.expectedTokenError(Colon)
				return
			}
			r.next()

			var v []byte
			DecodeBytes(r, &v)
			(*p)[k] = v

			t = r.next()
			switch t {
			case Comma:
				t = r.next()
				if t == ObjectEnd {
					r.invalidCharacterError()
					return
				}
			case ObjectEnd:
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
		r.expectedTokenError(ObjectBegin)
	}
}

func DecodeEnumNameMap[Enum ~int32](r *JsonDecoder, p *map[string]Enum, values map[string]int32) {
	switch r.token {
	case ObjectBegin:
		if *p == nil {
			*p = make(map[string]Enum)
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		t := r.next()
		if t == ObjectEnd {
			return
		}
		for {
			if t != String {
				r.expectedTokenError(String)
				return
			}
			k := r.readString()
			if r.next() != Colon {
				r.expectedTokenError(Colon)
				return
			}
			r.next()

			var v Enum
			DecodeEnumName(r, &v, values)
			(*p)[k] = v

			t = r.next()
			switch t {
			case Comma:
				t = r.next()
				if t == ObjectEnd {
					r.invalidCharacterError()
					return
				}
			case ObjectEnd:
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
		r.expectedTokenError(ObjectBegin)
	}
}

func DecodeEnumMap[Enum ~int32](r *JsonDecoder, p *map[string]Enum) {
	switch r.token {
	case ObjectBegin:
		if *p == nil {
			*p = make(map[string]Enum)
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		t := r.next()
		if t == ObjectEnd {
			return
		}
		for {
			if t != String {
				r.expectedTokenError(String)
				return
			}
			k := r.readString()
			if r.next() != Colon {
				r.expectedTokenError(Colon)
				return
			}
			r.next()

			var v Enum
			DecodeEnum(r, &v)
			(*p)[k] = v

			t = r.next()
			switch t {
			case Comma:
				t = r.next()
				if t == ObjectEnd {
					r.invalidCharacterError()
					return
				}
			case ObjectEnd:
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
		r.expectedTokenError(ObjectBegin)
	}
}

func DecodeMessageMap[Message any](r *JsonDecoder, p *map[string]*Message) {
	switch r.token {
	case ObjectBegin:
		if *p == nil {
			*p = make(map[string]*Message)
		}

		r.token = 0 // 指示next()执行"step info"而不是"step over"
		t := r.next()
		if t == ObjectEnd {
			return
		}
		for {
			if t != String {
				r.expectedTokenError(String)
				return
			}
			k := r.readString()
			if r.next() != Colon {
				r.expectedTokenError(Colon)
				return
			}
			r.next()

			var v *Message
			DecodeMessage(r, &v)
			(*p)[k] = v

			t = r.next()
			switch t {
			case Comma:
				t = r.next()
				if t == ObjectEnd {
					r.invalidCharacterError()
					return
				}
			case ObjectEnd:
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
		r.expectedTokenError(ObjectBegin)
	}
}
