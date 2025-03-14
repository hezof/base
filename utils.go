package core

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func Uniq[V comparable](s []V) []V {
	m := make(map[V]any)
	for _, v := range s {
		m[v] = nil
	}
	s = s[0:0]
	for v, _ := range m {
		s = append(s, v)
	}
	return s
}

func Keys[K comparable, V any](m map[K]V) []K {
	if m == nil {
		return nil
	}
	s := make([]K, 0, len(m))
	for k, _ := range m {
		s = append(s, k)
	}
	return s
}

func Vals[K comparable, V any](m map[K]V) []V {
	if m == nil {
		return nil
	}
	s := make([]V, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

func NvlS(ss ...string) string {
	for _, s := range ss {
		if s != "" {
			return s
		}
	}
	return ""
}

func NvlI[V int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](is ...V) V {
	for _, i := range is {
		if i != 0 {
			return i
		}
	}
	return 0
}

func NvlB(bs ...bool) bool {
	for _, b := range bs {
		if b {
			return b
		}
	}
	return false
}

var ZeroTime = time.Unix(0, 0)

func NvlT(ts ...time.Time) time.Time {
	for _, t := range ts {
		if t != ZeroTime {
			return t
		}
	}
	return ZeroTime
}

// Exist 判断路径是否存在
func Exist(path string) bool {
	stat, err := os.Stat(path)
	return stat != nil || os.IsExist(err)
}

// LookupExec 查找命令路径
func LookupExec(name string) string {
	// 1. 如果name存在
	if Exist(name) {
		return name
	}
	// 2. 当前目录查找
	if path := filepath.Join(filepath.Dir(os.Args[0]), name); Exist(path) {
		return path
	}
	// 3. 工作目录查找
	if cwd, err := os.Getwd(); err == nil {
		if path := filepath.Join(cwd, name); Exist(path) {
			return path
		}
	}
	// 4. 系统PATH查找
	if path, err := exec.LookPath(name); err == nil {
		if Exist(path) {
			return path
		}
	}
	// 5. 上述步骤都失败,返回初值碰运气!
	return name
}

// LocatePath 定位文件真实路径
func LocatePath(name string) (string, error) {
	// 1. 绝对名称直接返回
	if Exist(name) {
		return name, nil
	}
	// 1. 启动目录的默认配置
	path := filepath.Join(filepath.Dir(os.Args[0]), name)
	if Exist(path) {
		return path, nil
	}
	// 2. 工作目录的默认配置
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path = filepath.Join(cwd, name)
	if Exist(path) {
		return path, nil
	}
	return "", nil
}

// Hex 16进制编码
func Hex(src []byte) string {
	return hex.EncodeToString(src)
}

// Base64 标准base64编码
func Base64(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

// Base64URL url base64编码
func Base64URL(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

// UnHex 16进制解码
func UnHex(dst string) ([]byte, error) {
	return hex.DecodeString(dst)
}

// UnBase64 标准base64解码
func UnBase64(dst string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(dst)
}

// UnBase64URL url base64解码
func UnBase64URL(dst string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(dst)
}

// UnBase64Compatible base64兼容解码
func UnBase64Compatible(dst string) ([]byte, error) {
	if strings.IndexAny(dst, "-_") != -1 {
		return base64.URLEncoding.DecodeString(dst)
	}
	return base64.StdEncoding.DecodeString(dst)
}

func Md5(src []byte) []byte {
	h := md5.New()
	h.Write(src)
	return h.Sum(nil)
}

func Sha1(src []byte) []byte {
	h := sha1.New()
	h.Write(src)
	return h.Sum(nil)
}

func Sha256(src []byte) []byte {
	h := sha256.New()
	h.Write(src)
	return h.Sum(nil)
}

func HmacSha1(key []byte, val []byte) []byte {
	h := hmac.New(sha1.New, key)
	h.Write(val)
	return h.Sum(nil)
}

func HmacSha256(key []byte, val []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(val)
	return h.Sum(nil)
}
