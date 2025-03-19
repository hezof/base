package base

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"
)

var (
	ErrInvalidStructPointer      = errors.New("invalid struct pointer")        // 非法结构体指针
	ErrInvalidMemoryOrNilPointer = errors.New("invalid memory or nil pointer") // 不可用内存或空指针
	SimpleStructBinder           = &StructBinder{Fields: SimpleStructField}    // struct解析无缓存
	CachedStructBinder           = &StructBinder{Fields: CacheStructField}     // struct解析带缓存
)

var (
	RTypeBool     = reflect.TypeOf(false)
	RTypeInt      = reflect.TypeOf(int(0))
	RTypeInt8     = reflect.TypeOf(int8(0))
	RTypeInt16    = reflect.TypeOf(int16(0))
	RTypeInt32    = reflect.TypeOf(int32(0))
	RTypeInt64    = reflect.TypeOf(int64(0))
	RTypeUint     = reflect.TypeOf(uint(0))
	RTypeUint8    = reflect.TypeOf(uint8(0))
	RTypeUint16   = reflect.TypeOf(uint16(0))
	RTypeUint32   = reflect.TypeOf(uint32(0))
	RTypeUint64   = reflect.TypeOf(uint64(0))
	RTypeFloat32  = reflect.TypeOf(float32(0.0))
	RTypeFloat64  = reflect.TypeOf(0.0)
	RTypeTime     = reflect.TypeOf(ZeroTime)
	RTypeDuration = reflect.TypeOf(time.Duration(0))
	RTypeBytes    = reflect.TypeOf(([]byte)(nil))
	RTypeRunes    = reflect.TypeOf(([]rune)(nil))
	RTypeString   = reflect.TypeOf("")
	RTypeAnyMap   = reflect.TypeOf((map[string]any)(nil))
	RTypeAnySlice = reflect.TypeOf(([]any)(nil))
)

var structFieldsCache sync.Map

func CacheStructField(typ reflect.Type) []*StructField {
	ret, ok := structFieldsCache.Load(typ)
	if !ok {
		ret = SimpleStructField(typ)
		structFieldsCache.Store(typ, ret)
	}
	return ret.([]*StructField)
}

func SimpleStructField(typ reflect.Type) (ret []*StructField) {
	n := typ.NumField()
	ret = make([]*StructField, 0, n)
	for i := 0; i < n; i++ {
		fld := typ.Field(i)
		//  It is empty for upper case (exported) field names.
		if fld.PkgPath == "" || fld.Anonymous {
			// break unexported fields
			ret = append(ret, &StructField{
				Name:      fld.Name,
				Type:      fld.Type,
				Tags:      AllTag(string(fld.Tag)),
				Index:     i,
				Anonymous: fld.Anonymous,
			})
		}
	}
	n = len(ret)
	ret = ret[0:n:n] // trim
	return
}

type StructField struct {
	Name      string            // 名称
	Type      reflect.Type      // 类型
	Tags      map[string]string // tag
	Index     int               // 下标值. unexported字段会忽略掉
	Anonymous bool              // is an embedded field
}

type StructBinder struct {
	Fields func(typ reflect.Type) []*StructField
}

/*
MapStruct 将map转换成struct
- org: map[string]interface{}
- dst: struct
- tag: 用于关联的tag名字
*/
func (sb StructBinder) MapStruct(org map[string]any, dst any, tag string) error {

	if org == nil {
		return nil
	}

	var dstTyp = reflect.TypeOf(dst)
	if dstTyp == nil || dstTyp.Kind() != reflect.Ptr {
		return ErrInvalidStructPointer
	}

	var dstVal = reflect.ValueOf(dst)
	if dstVal.IsNil() {
		return ErrInvalidMemoryOrNilPointer
	}

	dstTyp = dstTyp.Elem()
	dstVal = dstVal.Elem()
	for dstTyp.Kind() == reflect.Ptr {
		dstTyp = dstTyp.Elem()
		if dstVal.IsNil() {
			dstVal.Set(reflect.New(dstTyp))
		}
		dstVal = dstVal.Elem()
	}
	if dstTyp.Kind() != reflect.Struct {
		return ErrInvalidStructPointer
	}

	_, err := sb.AdaptStruct(org, dstTyp, &dstVal, tag)
	return err
}

/*
AdaptStruct 调用者必须确保org不为nil
*/
func (sb StructBinder) AdaptStruct(org map[string]any, dstTyp reflect.Type, dstVal *reflect.Value, tag string) (set bool, err error) {

__NEXT__:
	for _, fld := range sb.Fields(dstTyp) {
		fldKey := fld.Tags[tag]
		if fldKey == "-" {
			continue __NEXT__
		}
		fldTyp := fld.Type
		fldVal := dstVal.Field(fld.Index)
		if fld.Anonymous {
			if fldTyp.Kind() == reflect.Ptr {
				fldTyp = fldTyp.Elem()
				tmpVal := reflect.New(fldTyp)
				subVal := tmpVal.Elem()
				subSet := false
				if subSet, err = sb.AdaptStruct(org, fldTyp, &subVal, tag); err != nil {
					return
				} else if subSet {
					set = true
					fldVal.Set(tmpVal)
				}
			} else {
				if _, err = sb.AdaptStruct(org, fldTyp, &fldVal, tag); err != nil {
					return
				} else {
					set = true
				}
			}
			continue __NEXT__
		}
		if fldKey == "" {
			fldKey = fld.Name
		}
		if val := org[fldKey]; val != nil {
			_set := false
			if _set, err = sb.AdaptValue(val, fldTyp, &fldVal, tag); err != nil {
				return
			} else if _set {
				set = true
			}
		}
	}
	return
}

/*
AdaptValue 调用者必须确保org不为nil
*/
func (sb StructBinder) AdaptValue(org interface{}, dstTyp reflect.Type, dstVal *reflect.Value, tag string) (bool, error) {

	switch dstTyp.Kind() {
	case reflect.Bool:
		if x, err := AsBool(org); err == nil {
			dstVal.SetBool(x)
			return true, nil
		}
	case reflect.Int:
		if x, err := AsInt64(org); err == nil {
			dstVal.SetInt(x)
			return true, nil
		}
	case reflect.Int8:
		if x, err := AsInt64(org); err == nil {
			dstVal.SetInt(x)
			return true, nil
		}
	case reflect.Int16:
		if x, err := AsInt64(org); err == nil {
			dstVal.SetInt(x)
			return true, nil
		}
	case reflect.Int32:
		if x, err := AsInt64(org); err == nil {
			dstVal.SetInt(x)
			return true, nil
		}
	case reflect.Int64:
		if dstTyp == RTypeDuration {
			if x, err := AsDuration(org); err == nil {
				dstVal.SetInt(int64(x))
				return true, nil
			}
		} else {
			if x, err := AsInt64(org); err == nil {
				dstVal.SetInt(x)
				return true, nil
			}
		}
	case reflect.Uint:
		if x, err := AsUint64(org); err == nil {
			dstVal.SetUint(x)
			return true, nil
		}
	case reflect.Uint8:
		if x, err := AsUint64(org); err == nil {
			dstVal.SetUint(x)
			return true, nil
		}
	case reflect.Uint16:
		if x, err := AsUint64(org); err == nil {
			dstVal.SetUint(x)
			return true, nil
		}
	case reflect.Uint32:
		if x, err := AsUint64(org); err == nil {
			dstVal.SetUint(x)
			return true, nil
		}
	case reflect.Uint64:
		if x, err := AsUint64(org); err == nil {
			dstVal.SetUint(x)
			return true, nil
		}
	case reflect.Uintptr:
		if x, err := AsUintptr(org); err == nil {
			dstVal.Set(reflect.ValueOf(x))
			return true, nil
		}
	case reflect.Float32:
		if x, err := AsFloat64(org); err == nil {
			dstVal.SetFloat(x)
			return true, nil
		}
	case reflect.Float64:
		if x, err := AsFloat64(org); err == nil {
			dstVal.SetFloat(x)
			return true, nil
		}
	case reflect.Complex64:
		if x, err := AsComplex128(org); err == nil {
			dstVal.SetComplex(x)
			return true, nil
		}
	case reflect.Complex128:
		if x, err := AsComplex128(org); err == nil {
			dstVal.SetComplex(x)
			return true, nil
		}
	case reflect.Array:
		orgTyp := reflect.TypeOf(org)
		if orgTyp == dstTyp || orgTyp.AssignableTo(dstTyp) {
			dstVal.Set(reflect.ValueOf(org))
			return true, nil
		}
	case reflect.Chan:
		orgTyp := reflect.TypeOf(org)
		if orgTyp == dstTyp || orgTyp.AssignableTo(dstTyp) {
			dstVal.Set(reflect.ValueOf(org))
			return true, nil
		}
	case reflect.Func:
		orgTyp := reflect.TypeOf(org)
		if orgTyp == dstTyp || orgTyp.AssignableTo(dstTyp) {
			dstVal.Set(reflect.ValueOf(org))
			return true, nil
		}
	case reflect.Interface:
		orgTyp := reflect.TypeOf(org)
		if orgTyp == dstTyp || orgTyp.AssignableTo(dstTyp) {
			dstVal.Set(reflect.ValueOf(org))
			return true, nil
		}
	case reflect.Map:
		orgTyp := reflect.TypeOf(org)
		if orgTyp == dstTyp || orgTyp.AssignableTo(dstTyp) {
			dstVal.Set(reflect.ValueOf(org))
			return true, nil
		}
		// map[string]interface到map[string]struct
		if orgTyp == RTypeAnyMap {
			if dstTyp.Key().Kind() == reflect.String {

				// 需要判空
				if dstVal.IsNil() {
					dstVal.Set(reflect.MakeMap(dstTyp))
				}

				var set bool
				typ := dstTyp.Elem()
				// 已经判断orgTyp是GenericStruct
				for k, v := range org.(map[string]interface{}) {
					if v == nil {
						continue
					}
					val := reflect.New(typ).Elem()
					if _set, err := sb.AdaptValue(v, typ, &val, tag); err != nil {
						return false, err
					} else if _set {
						dstVal.SetMapIndex(reflect.ValueOf(k), val)
						set = true
					}
				}
				return set, nil
			}
		}
	case reflect.Ptr:

		typ := dstTyp.Elem()
		if dstVal.IsNil() {
			// 如果是空指针,则分配新值. 并判断是否有set操作, 有的话才设置
			tmp := reflect.New(typ)
			val := tmp.Elem()
			if _set, err := sb.AdaptValue(org, typ, &val, tag); err != nil {
				return false, err
			} else if _set {
				dstVal.Set(tmp) // 中间值是个指针
				return true, nil
			} else {
				return false, nil
			}
		} else {
			val := dstVal.Elem()
			return sb.AdaptValue(org, typ, &val, tag)
		}

	case reflect.Slice:
		// []byte
		if dstTyp == RTypeBytes {
			if x, err := AsBytes(org); err == nil {
				dstVal.Set(reflect.ValueOf(x))
				return true, nil
			}
		} else if dstTyp == RTypeRunes {
			if x, err := AsRunes(org); err == nil {
				dstVal.Set(reflect.ValueOf(x))
				return true, nil
			}
		} else {
			orgTyp := reflect.TypeOf(org)
			if orgTyp == dstTyp || orgTyp.AssignableTo(dstTyp) {
				dstVal.Set(reflect.ValueOf(org))
				return true, nil
			}
			// []interface{}
			if orgTyp == RTypeAnySlice {
				// 已经判断orgTyp是GenericSlice
				orgVal := org.([]interface{})
				orgLen := len(orgVal)
				if n := dstVal.Cap(); n > orgLen {
					// 重置
					dstVal.SetLen(orgLen)
					zero := reflect.Zero(dstTyp.Elem())
					for i := 0; i < orgLen; i++ {
						dstVal.Index(i).Set(zero)
					}
				} else {
					dstVal.Set(reflect.MakeSlice(dstTyp, orgLen, orgLen))
				}
				var set bool
				for i, v := range orgVal {
					if v == nil {
						continue
					}
					typ := dstTyp.Elem()
					val := dstVal.Index(i)
					for typ.Kind() == reflect.Ptr {
						typ = typ.Elem()
						if val.IsNil() {
							val.Set(reflect.New(typ))
						}
						val = val.Elem()
					}
					if _set, err := sb.AdaptValue(v, typ, &val, tag); err != nil {
						return false, err
					} else if _set {
						set = true
					}
				}
				return set, nil
			}
		}

	case reflect.String:
		if x, err := AsString(org); err == nil {
			dstVal.SetString(x)
			return true, nil
		}
	case reflect.Struct:
		// time.Time
		if dstTyp == RTypeTime {
			if x, err := AsTime(org); err == nil {
				dstVal.Set(reflect.ValueOf(x))
				return true, nil
			}
		} else {
			orgTyp := reflect.TypeOf(org)
			if orgTyp == dstTyp || orgTyp.AssignableTo(dstTyp) {
				dstVal.Set(reflect.ValueOf(org))
				return true, nil
			}
			// map[string]interface{}
			if orgTyp == RTypeAnyMap {
				// 已经判断orgTyp是GenericStruct
				if _set, err := sb.AdaptStruct(org.(map[string]interface{}), dstTyp, dstVal, tag); err != nil {
					return false, err
				} else {
					return _set, nil
				}
			}
		}
	case reflect.UnsafePointer:
		if x, err := AsUnsafePointer(org); err == nil {
			dstVal.Set(reflect.ValueOf(x))
			return true, nil
		}
	}
	return false, fmt.Errorf("can't convert from type %T to %v", org, dstTyp)
}

func MapStruct(org map[string]interface{}, dst interface{}, tag string) error {
	return CachedStructBinder.MapStruct(org, dst, tag)
}
