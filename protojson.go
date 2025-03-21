package framework

import (
	"fmt"
	"github.com/hezof/base/internal/json"
	"io"
	"reflect"
	"sync"
)

var (
	ProtoJsonDecoderBufferSize = 8 * 1024       // 默认8K
	ProtoJsonEncoderBufferSize = 8 * 1024       // 默认8K
	EncodeJson                 = json.Marshal   // 标准编码
	DecodeJson                 = json.Unmarshal // 标准解码
)

func DecodeJsonReader(in io.Reader, val any) error {
	data, err := io.ReadAll(in)
	if err != nil {
		return err
	}
	return DecodeJson(data, val)
}

/*
FieldCodec 核心接口, 实现Message的解码与编码.
该接口用于加速proto.Message的JSON解码/编码速度.
*/
type FieldCodec interface {
	DecodeField(r *JsonDecoder, f string)
	EncodeField(w *JsonEncoder)
}

func DecodeAny(r *JsonDecoder, val any) {
	switch r.token {
	case ObjectBegin:
		if jc, ok := val.(FieldCodec); ok {
			r.readObject(jc)
		} else {
			err := DecodeJson(r.dumpObjectOrArray(ObjectBegin), val)
			if err != nil {
				r.reportError(err)
			}
		}
	case ObjectEnd:
		r.invalidCharacterError()
	case ArrayBegin:
		err := DecodeJson(r.dumpObjectOrArray(ArrayBegin), val)
		if err != nil {
			r.reportError(err)
		}
	case ArrayEnd:
		r.invalidCharacterError()
	case Comma:
		r.invalidCharacterError()
	case Colon:
		r.invalidCharacterError()
	case String:
		rv := reflect.ValueOf(val)
		rt := rv.Type()
		for rt.Kind() == reflect.Pointer {
			rt = rt.Elem()
			if rv.IsNil() {
				if rv.CanSet() {
					rv.Set(reflect.New(rt))
				}
			}
			rv = rv.Elem()
		}
		switch rt.Kind() {
		case reflect.String:
			if rv.CanSet() {
				rv.SetString(r.readString())
			}
		default:
			r.reportError(fmt.Errorf("cannot unmarshal string into type %v", rv.Kind()))
		}
	case Number:
		rv := reflect.ValueOf(val)
		rt := rv.Type()
		for rt.Kind() == reflect.Pointer {
			rt = rt.Elem()
			if rv.IsNil() {
				if rv.CanSet() {
					rv.Set(reflect.New(rt))
				}
			}
			rv = rv.Elem()
		}
		switch rv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if rv.CanSet() {
				rv.SetInt(r.readInt64())
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if rv.CanSet() {
				rv.SetUint(r.readUint64())
			}
		case reflect.Float32, reflect.Float64:
			if rv.CanSet() {
				rv.SetFloat(r.readFloat64())
			}
		default:
			r.reportError(fmt.Errorf("cannot unmarshal number into type %v", rv.Kind()))
		}
	case True:
		r.skipTrue()
		rv := reflect.ValueOf(val)
		rt := rv.Type()
		for rt.Kind() == reflect.Pointer {
			rt = rt.Elem()
			if rv.IsNil() {
				if rv.CanSet() {
					rv.Set(reflect.New(rt))
				}
			}
			rv = rv.Elem()
		}
		switch rv.Kind() {
		case reflect.Bool:
			if rv.CanSet() {
				rv.SetBool(true)
			}
		default:
			r.reportError(fmt.Errorf("cannot unmarshal true into type %v", rv.Kind()))
		}
	case False:
		r.skipFalse()
		rv := reflect.ValueOf(val)
		rt := rv.Type()
		for rt.Kind() == reflect.Pointer {
			rt = rt.Elem()
			if rv.IsNil() {
				if rv.CanSet() {
					rv.Set(reflect.New(rt))
				}
			}
			rv = rv.Elem()
		}
		switch rv.Kind() {
		case reflect.Bool:
			if rv.CanSet() {
				rv.SetBool(false)
			}
		default:
			r.reportError(fmt.Errorf("cannot unmarshal false into type %v", rv.Kind()))
		}
	case Null:
		r.skipNull()
	case 0:
		r.unexpectedEndError()
	case -1:
	}

}

// EncodeAny 与EncodeMessage很相似, 但后者是泛型方法.
func EncodeAny(w *JsonEncoder, val any) {
	if jc, ok := val.(FieldCodec); ok {
		w.ensure(2)
		w.buff = append(w.buff, leftBrace)
		jc.EncodeField(w)
		if last := len(w.buff) - 1; w.buff[last] == comma {
			w.buff[last] = rightBrace
		} else {
			w.buff = append(w.buff, rightBrace)
		}
	} else {
		bs, err := EncodeJson(val)
		if err != nil {
			w.reportError(err)
		}
		_, _ = w.Write(bs)
	}
}

func EncodeAny_OmitEmpty(w *JsonEncoder, name string, val any) {
	if val != nil {
		w.ensure(4 + len(name))
		w.buff = append(w.buff, quotes)
		w.buff = append(w.buff, name...)
		w.buff = append(w.buff, quotes, colon)
		EncodeAny(w, val)
		w.buff = append(w.buff, comma)
	}
}

func EncodeAny_WithEmpty(w *JsonEncoder, name string, val any) {
	if val != nil {
		w.ensure(4 + len(name))
		w.buff = append(w.buff, quotes)
		w.buff = append(w.buff, name...)
		w.buff = append(w.buff, quotes, colon)
		EncodeAny(w, val)
		w.buff = append(w.buff, comma)
	} else {
		w.ensure(8 + len(name))
		w.buff = append(w.buff, quotes)
		w.buff = append(w.buff, name...)
		w.buff = append(w.buff, quotes, colon, 'n', 'u', 'l', 'l', comma)
	}
}

var decoders = sync.Pool{
	New: func() interface{} {
		return NewJsonDecoder(nil, ProtoJsonDecoderBufferSize)
	},
}

var encoders = sync.Pool{
	New: func() interface{} {
		return NewJsonEncoder(nil, ProtoJsonEncoderBufferSize)
	},
}

func GetDecoder(in io.Reader) *JsonDecoder {
	return decoders.Get().(*JsonDecoder).Reset(in)
}

func PutDecoder(dec *JsonDecoder) {
	decoders.Put(dec.Clean())
}

func GetEncoder(out io.Writer) *JsonEncoder {
	return encoders.Get().(*JsonEncoder).Reset(out)
}

func PutEncoder(enc *JsonEncoder) {
	encoders.Put(enc.Clean())
}

func DecodeProtoJsonReader(in io.Reader, val any) error {
	// 加速实现JsonCodec
	d := GetDecoder(in)
	defer PutDecoder(d)

	DecodeAny(d, val)
	return d.Close()
}

func DecodeProtoJsonData(data []byte, val any) error {
	// 加速实现JsonCodec
	d := NewJsonBuffer(data)

	DecodeAny(d, val)
	return d.Close()
}

func EncodeProtoJsonWriter(out io.Writer, val any) error {
	e := GetEncoder(out)
	defer PutEncoder(e)

	EncodeAny(e, val)
	return e.Close()
}

func EncodeProtoJsonData(v any) ([]byte, error) {
	w := NewJsonEncoder(nil, 1024)
	EncodeAny(w, v)
	if err := w.Close(); err != nil {
		return nil, err
	}
	return w.buff, nil
}
