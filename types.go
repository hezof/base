package base

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

/************************************************
 * 类型断言函数
 ************************************************/

func AsString(val any) (string, error) {
	switch val := val.(type) {
	case string:
		return val, nil
	case bool:
		return strconv.FormatBool(val), nil
	case int:
		return strconv.FormatInt(int64(val), 10), nil
	case int8:
		return strconv.FormatInt(int64(val), 10), nil
	case int16:
		return strconv.FormatInt(int64(val), 10), nil
	case int32:
		return strconv.FormatInt(int64(val), 10), nil
	case int64:
		return strconv.FormatInt(val, 10), nil
	case uint:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint64:
		return strconv.FormatUint(val, 10), nil
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64), nil
	case time.Duration:
		return val.String(), nil
	case time.Time:
		return val.Format(time.DateTime), nil
	case nil:
		return "", nil
	default:
		return "", fmt.Errorf("can't convert from type %T to string", val)
	}
}

func AsBool(val any) (bool, error) {
	switch val := val.(type) {
	case string:
		return val == "true", nil
	case bool:
		return val, nil
	case nil:
		return false, nil
	default:
		return false, fmt.Errorf("can't convert from type %T to bool", val)
	}
}

const (
	LayoutDateTime       = "2006-01-02 15:04:05"
	LayoutDateTimeT      = "2006-01-02T15:04:05"
	LayoutDateTimeLength = len(LayoutDateTime)
)

func DateTime(ds string) (time.Time, error) {
	t := strings.IndexByte(ds, 'T')
	n := len(ds)
	switch {
	case t == -1 && n == LayoutDateTimeLength:
		return time.ParseInLocation(ds, LayoutDateTime, time.Local)
	case t == -1 && n < LayoutDateTimeLength:
		return time.ParseInLocation(ds, LayoutDateTime[:n], time.Local)
	case t == -1 && n > LayoutDateTimeLength:
		return time.ParseInLocation(ds[:n], LayoutDateTime, time.Local)
	case t != -1 && n == LayoutDateTimeLength:
		return time.ParseInLocation(ds, LayoutDateTimeT, time.Local)
	case t != -1 && n < LayoutDateTimeLength:
		return time.ParseInLocation(ds, LayoutDateTimeT[:n], time.Local)
	case t != -1 && n > LayoutDateTimeLength:
		return time.ParseInLocation(ds[:n], LayoutDateTimeT, time.Local)
	}
	return ZeroTime, fmt.Errorf("invalid datetime format %v", ds)
}

func AsTime(val any) (time.Time, error) {
	switch val := val.(type) {
	case string:
		return DateTime(val)
	case int:
		return time.Unix(int64(val), 0), nil
	case int8:
		return time.Unix(int64(val), 0), nil
	case int16:
		return time.Unix(int64(val), 0), nil
	case int32:
		return time.Unix(int64(val), 0), nil
	case int64:
		return time.Unix(int64(val), 0), nil
	case uint:
		return time.Unix(int64(val), 0), nil
	case uint8:
		return time.Unix(int64(val), 0), nil
	case uint16:
		return time.Unix(int64(val), 0), nil
	case uint32:
		return time.Unix(int64(val), 0), nil
	case uint64:
		return time.Unix(int64(val), 0), nil
	case float32:
		return time.Unix(int64(val), 0), nil
	case float64:
		return time.Unix(int64(val), 0), nil
	case time.Time:
		return val, nil
	case nil:
		return ZeroTime, nil
	default:
		return ZeroTime, fmt.Errorf("can't convert from type %T to time", val)
	}
}

func AsDuration(val any) (time.Duration, error) {
	switch val := val.(type) {
	case string:
		return time.ParseDuration(val)
	case int:
		return time.Second * time.Duration(val), nil
	case int8:
		return time.Second * time.Duration(val), nil
	case int16:
		return time.Second * time.Duration(val), nil
	case int32:
		return time.Second * time.Duration(val), nil
	case int64:
		return time.Second * time.Duration(val), nil
	case uint:
		return time.Second * time.Duration(val), nil
	case uint8:
		return time.Second * time.Duration(val), nil
	case uint16:
		return time.Second * time.Duration(val), nil
	case uint32:
		return time.Second * time.Duration(val), nil
	case uint64:
		return time.Second * time.Duration(val), nil
	case float32:
		return time.Second * time.Duration(val), nil
	case float64:
		return time.Second * time.Duration(val), nil
	case time.Duration:
		return val, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't convert from type %T to duration", val)
	}
}

func AsInt(val any) (int, error) {
	type Type = int
	switch val := val.(type) {
	case string:
		iv, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0, err
		}
		return Type(iv), nil
	case int:
		return Type(val), nil
	case int8:
		return Type(val), nil
	case int16:
		return Type(val), nil
	case int32:
		return Type(val), nil
	case int64:
		return Type(val), nil
	case uint:
		return Type(val), nil
	case uint8:
		return Type(val), nil
	case uint16:
		return Type(val), nil
	case uint32:
		return Type(val), nil
	case uint64:
		return Type(val), nil
	case float32:
		return Type(val), nil
	case float64:
		return Type(val), nil
	case time.Duration:
		return Type(val), nil
	case time.Time:
		return Type(val.Unix()), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't convert from type %T to int", val)
	}
}

func AsInt8(val any) (int8, error) {
	type Type = int8
	switch val := val.(type) {
	case string:
		iv, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0, err
		}
		return Type(iv), nil
	case int:
		return Type(val), nil
	case int8:
		return Type(val), nil
	case int16:
		return Type(val), nil
	case int32:
		return Type(val), nil
	case int64:
		return Type(val), nil
	case uint:
		return Type(val), nil
	case uint8:
		return Type(val), nil
	case uint16:
		return Type(val), nil
	case uint32:
		return Type(val), nil
	case uint64:
		return Type(val), nil
	case float32:
		return Type(val), nil
	case float64:
		return Type(val), nil
	case time.Duration:
		return Type(val), nil
	case time.Time:
		return Type(val.Unix()), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't convert from type %T to int8", val)
	}
}

func AsInt16(val any) (int16, error) {
	type Type = int16
	switch val := val.(type) {
	case string:
		iv, err := strconv.ParseInt(val, 10, 16)
		if err != nil {
			return 0, err
		}
		return Type(iv), nil
	case int:
		return Type(val), nil
	case int8:
		return Type(val), nil
	case int16:
		return Type(val), nil
	case int32:
		return Type(val), nil
	case int64:
		return Type(val), nil
	case uint:
		return Type(val), nil
	case uint8:
		return Type(val), nil
	case uint16:
		return Type(val), nil
	case uint32:
		return Type(val), nil
	case uint64:
		return Type(val), nil
	case float32:
		return Type(val), nil
	case float64:
		return Type(val), nil
	case time.Duration:
		return Type(val), nil
	case time.Time:
		return Type(val.Unix()), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't convert from type %T to int16", val)
	}
}

func AsInt32(val any) (int32, error) {
	type Type = int32
	switch val := val.(type) {
	case string:
		iv, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return 0, err
		}
		return Type(iv), nil
	case int:
		return Type(val), nil
	case int8:
		return Type(val), nil
	case int16:
		return Type(val), nil
	case int32:
		return Type(val), nil
	case int64:
		return Type(val), nil
	case uint:
		return Type(val), nil
	case uint8:
		return Type(val), nil
	case uint16:
		return Type(val), nil
	case uint32:
		return Type(val), nil
	case uint64:
		return Type(val), nil
	case float32:
		return Type(val), nil
	case float64:
		return Type(val), nil
	case time.Duration:
		return Type(val), nil
	case time.Time:
		return Type(val.Unix()), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't convert from type %T to int32", val)
	}
}

func AsInt64(val any) (int64, error) {
	type Type = int64
	switch val := val.(type) {
	case string:
		iv, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0, err
		}
		return Type(iv), nil
	case int:
		return Type(val), nil
	case int8:
		return Type(val), nil
	case int16:
		return Type(val), nil
	case int32:
		return Type(val), nil
	case int64:
		return Type(val), nil
	case uint:
		return Type(val), nil
	case uint8:
		return Type(val), nil
	case uint16:
		return Type(val), nil
	case uint32:
		return Type(val), nil
	case uint64:
		return Type(val), nil
	case float32:
		return Type(val), nil
	case float64:
		return Type(val), nil
	case time.Duration:
		return Type(val), nil
	case time.Time:
		return Type(val.Unix()), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't convert from type %T to int64", val)
	}
}

func AsUint(val any) (uint, error) {
	type Type = uint
	switch val := val.(type) {
	case string:
		iv, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return 0, err
		}
		return Type(iv), nil
	case int:
		return Type(val), nil
	case int8:
		return Type(val), nil
	case int16:
		return Type(val), nil
	case int32:
		return Type(val), nil
	case int64:
		return Type(val), nil
	case uint:
		return Type(val), nil
	case uint8:
		return Type(val), nil
	case uint16:
		return Type(val), nil
	case uint32:
		return Type(val), nil
	case uint64:
		return Type(val), nil
	case float32:
		return Type(val), nil
	case float64:
		return Type(val), nil
	case time.Duration:
		return Type(val), nil
	case time.Time:
		return Type(val.Unix()), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't convert from type %T to uint", val)
	}
}

func AsUint8(val any) (uint8, error) {
	type Type = uint8
	switch val := val.(type) {
	case string:
		iv, err := strconv.ParseUint(val, 10, 8)
		if err != nil {
			return 0, err
		}
		return Type(iv), nil
	case int:
		return Type(val), nil
	case int8:
		return Type(val), nil
	case int16:
		return Type(val), nil
	case int32:
		return Type(val), nil
	case int64:
		return Type(val), nil
	case uint:
		return Type(val), nil
	case uint8:
		return Type(val), nil
	case uint16:
		return Type(val), nil
	case uint32:
		return Type(val), nil
	case uint64:
		return Type(val), nil
	case float32:
		return Type(val), nil
	case float64:
		return Type(val), nil
	case time.Duration:
		return Type(val), nil
	case time.Time:
		return Type(val.Unix()), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't convert from type %T to uint8", val)
	}
}

func AsUint16(val any) (uint16, error) {
	type Type = uint16
	switch val := val.(type) {
	case string:
		iv, err := strconv.ParseUint(val, 10, 16)
		if err != nil {
			return 0, err
		}
		return Type(iv), nil
	case int:
		return Type(val), nil
	case int8:
		return Type(val), nil
	case int16:
		return Type(val), nil
	case int32:
		return Type(val), nil
	case int64:
		return Type(val), nil
	case uint:
		return Type(val), nil
	case uint8:
		return Type(val), nil
	case uint16:
		return Type(val), nil
	case uint32:
		return Type(val), nil
	case uint64:
		return Type(val), nil
	case float32:
		return Type(val), nil
	case float64:
		return Type(val), nil
	case time.Duration:
		return Type(val), nil
	case time.Time:
		return Type(val.Unix()), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't convert from type %T to uint16", val)
	}
}

func AsUint32(val any) (uint32, error) {
	type Type = uint32
	switch val := val.(type) {
	case string:
		iv, err := strconv.ParseUint(val, 10, 32)
		if err != nil {
			return 0, err
		}
		return Type(iv), nil
	case int:
		return Type(val), nil
	case int8:
		return Type(val), nil
	case int16:
		return Type(val), nil
	case int32:
		return Type(val), nil
	case int64:
		return Type(val), nil
	case uint:
		return Type(val), nil
	case uint8:
		return Type(val), nil
	case uint16:
		return Type(val), nil
	case uint32:
		return Type(val), nil
	case uint64:
		return Type(val), nil
	case float32:
		return Type(val), nil
	case float64:
		return Type(val), nil
	case time.Duration:
		return Type(val), nil
	case time.Time:
		return Type(val.Unix()), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't convert from type %T to uint32", val)
	}
}

func AsUint64(val any) (uint64, error) {
	type Type = uint64
	switch val := val.(type) {
	case string:
		iv, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return 0, err
		}
		return Type(iv), nil
	case int:
		return Type(val), nil
	case int8:
		return Type(val), nil
	case int16:
		return Type(val), nil
	case int32:
		return Type(val), nil
	case int64:
		return Type(val), nil
	case uint:
		return Type(val), nil
	case uint8:
		return Type(val), nil
	case uint16:
		return Type(val), nil
	case uint32:
		return Type(val), nil
	case uint64:
		return Type(val), nil
	case float32:
		return Type(val), nil
	case float64:
		return Type(val), nil
	case time.Duration:
		return Type(val), nil
	case time.Time:
		return Type(val.Unix()), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't convert from type %T to uint64", val)
	}
}

func AsFloat32(val any) (float32, error) {
	type Type = float32
	switch val := val.(type) {
	case string:
		iv, err := strconv.ParseFloat(val, 32)
		if err != nil {
			return 0, err
		}
		return Type(iv), nil
	case int:
		return Type(val), nil
	case int8:
		return Type(val), nil
	case int16:
		return Type(val), nil
	case int32:
		return Type(val), nil
	case int64:
		return Type(val), nil
	case uint:
		return Type(val), nil
	case uint8:
		return Type(val), nil
	case uint16:
		return Type(val), nil
	case uint32:
		return Type(val), nil
	case uint64:
		return Type(val), nil
	case float32:
		return Type(val), nil
	case float64:
		return Type(val), nil
	case time.Duration:
		return Type(val), nil
	case time.Time:
		return Type(val.Unix()), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't convert from type %T to float32", val)
	}
}

func AsFloat64(val any) (float64, error) {
	type Type = float64
	switch val := val.(type) {
	case string:
		iv, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0, err
		}
		return Type(iv), nil
	case int:
		return Type(val), nil
	case int8:
		return Type(val), nil
	case int16:
		return Type(val), nil
	case int32:
		return Type(val), nil
	case int64:
		return Type(val), nil
	case uint:
		return Type(val), nil
	case uint8:
		return Type(val), nil
	case uint16:
		return Type(val), nil
	case uint32:
		return Type(val), nil
	case uint64:
		return Type(val), nil
	case float32:
		return Type(val), nil
	case float64:
		return Type(val), nil
	case time.Duration:
		return Type(val), nil
	case time.Time:
		return Type(val.Unix()), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't convert from type %T to float64", val)
	}
}

func AsUintptr(val any) (uintptr, error) {
	switch val := val.(type) {
	case uintptr:
		return val, nil
	case unsafe.Pointer:
		return uintptr(val), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't convert from type %T to uintptr", val)
	}
}

func AsUnsafePointer(val any) (unsafe.Pointer, error) {
	switch val := val.(type) {
	case uintptr:
		return unsafe.Pointer(val), nil
	case unsafe.Pointer:
		return val, nil
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("can't convert from type %T to unsafe.Pointer", val)
	}
}

func AsComplex128(val any) (complex128, error) {
	switch val := val.(type) {
	case complex128:
		return val, nil
	case complex64:
		return complex128(val), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("can't convert from type %T to unsafe.Pointer", val)
	}
}

func AsBytes(val any) ([]byte, error) {
	switch val := val.(type) {
	case []byte:
		return val, nil
	case string:
		return []byte(val), nil
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("can't convert from type %T to []byte", val)
	}
}

func AsRunes(val any) ([]rune, error) {
	switch val := val.(type) {
	case []rune:
		return val, nil
	case []byte:
		ret := make([]rune, len(val))
		for i, v := range val {
			ret[i] = rune(v)
		}
		return ret, nil
	case string:
		return []rune(val), nil
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("can't convert from type %T to []rune", val)
	}
}

/************************************************
 * 类型转换函数
 ************************************************/

func ToString(val any) string {
	v, _ := AsString(val)
	return v
}

func ToBool(val any) bool {
	v, _ := AsBool(val)
	return v
}

func ToTime(val any) time.Time {
	v, _ := AsTime(val)
	return v
}

func ToDuration(val any) time.Duration {
	v, _ := AsDuration(val)
	return v
}

func ToInt(val any) int {
	v, _ := AsInt(val)
	return v
}

func ToInt8(val any) int8 {
	v, _ := AsInt8(val)
	return v
}

func ToInt16(val any) int16 {
	v, _ := AsInt16(val)
	return v
}

func ToInt32(val any) int32 {
	v, _ := AsInt32(val)
	return v
}

func ToInt64(val any) int64 {
	v, _ := AsInt64(val)
	return v
}

func ToUint(val any) uint {
	v, _ := AsUint(val)
	return v
}

func ToUint8(val any) uint8 {
	v, _ := AsUint8(val)
	return v
}

func ToUint16(val any) uint16 {
	v, _ := AsUint16(val)
	return v
}

func ToUint32(val any) uint32 {
	v, _ := AsUint32(val)
	return v
}

func ToUint64(val any) uint64 {
	v, _ := AsUint64(val)
	return v
}

func ToFloat32(val any) float32 {
	v, _ := AsFloat32(val)
	return v
}

func ToFloat64(val any) float64 {
	v, _ := AsFloat64(val)
	return v
}
