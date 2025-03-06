package core

import (
	"encoding/json"
	"unsafe"
)

// UnsafeBytes string到[]byte的不安全转换
// For more Details, see https://github.com/golang/go/issues/53003#issuecomment-1140276077.
func UnsafeBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// UnsafeString []byte到string的不安全转换
// For more Details, see https://github.com/golang/go/issues/53003#issuecomment-1140276077.
func UnsafeString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// ToJson 输出json格式
func ToJson(v any) string {
	bs, _ := json.Marshal(v)
	return string(bs)
}
