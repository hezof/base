package framework

import "fmt"

var (
	TRUE  = true
	FALSE = false
)

func DecodeBoolOptional(r *JsonDecoder, p **bool) {
	switch r.token {
	case True:
		r.skipTrue()
		*p = &TRUE
	case False:
		r.skipFalse()
		*p = &FALSE
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(True)
	}
}

func DecodeInt32Optional(r *JsonDecoder, p **int32) {
	switch r.token {
	case Number:
		v := int32(r.readInt64())
		*p = &v
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(Number)
	}
}

func DecodeInt64Optional(r *JsonDecoder, p **int64) {
	switch r.token {
	case Number:
		v := r.readInt64()
		*p = &v
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(Number)
	}
}

func DecodeUint32Optional(r *JsonDecoder, p **uint32) {
	switch r.token {
	case Number:
		v := uint32(r.readUint64())
		*p = &v
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(Number)
	}
}

func DecodeUint64Optional(r *JsonDecoder, p **uint64) {
	switch r.token {
	case Number:
		v := r.readUint64()
		*p = &v
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(Number)
	}
}

func DecodeFloatOptional(r *JsonDecoder, p **float32) {
	switch r.token {
	case Number:
		v := float32(r.readFloat64())
		*p = &v
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(Number)
	}
}

func DecodeDoubleOptional(r *JsonDecoder, p **float64) {
	switch r.token {
	case Number:
		v := r.readFloat64()
		*p = &v
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(Number)
	}
}

func DecodeStringOptional(r *JsonDecoder, p **string) {
	switch r.token {
	case String:
		v := r.readString()
		*p = &v
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(String)
	}
}

func DecodeBytesOptional(r *JsonDecoder, p *[]byte) {
	DecodeBytes(r, p)
}

func DecodeEnumNameOptional[Enum ~int32](r *JsonDecoder, p **Enum, values map[string]int32) {
	switch r.token {
	case String:
		s := r.readString()
		v, ok := values[s]
		if !ok {
			r.reportError(fmt.Errorf("invalid enum: %v", s))
			return
		}
		e := Enum(v)
		*p = &e
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(Number)
	}
}

func DecodeEnumOptional[Enum ~int32](r *JsonDecoder, p **Enum) {
	switch r.token {
	case Number:
		e := Enum(r.readInt64())
		*p = &e
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	default:
		r.expectedTokenError(Number)
	}
}

func DecodeMessageOptional[Message any](r *JsonDecoder, p **Message) {
	DecodeMessage(r, p)
}
