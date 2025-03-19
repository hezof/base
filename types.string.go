package base

import (
	"encoding/json"
	"github.com/hezof/base/internal/toml"
	"unsafe"
)

/************************************************
 * json & toml
 ************************************************/

func AsJson(v any) (string, error) {
	bs, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return UnsafeString(bs), nil
}

func AsToml(v any) (string, error) {
	bs, err := toml.Marshal(v)
	if err != nil {
		return "", err
	}
	return UnsafeString(bs), nil
}

// ToJson 输出json格式
func ToJson(v any) string {
	bs, _ := json.Marshal(v)
	return UnsafeString(bs)
}

// ToToml 输出toml格式
func ToToml(v any) string {
	bs, _ := toml.Marshal(v)
	return UnsafeString(bs)
}

/************************************************
 * string与bytes转换函数(非安全方式)
 ************************************************/

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
