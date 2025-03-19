package base

import (
	"encoding/base64"
	"io"
	"strconv"
	"unicode/utf8"
)

/*
JsonEncoder 核心结构, 实现struct的编码.
*/

const MinimumBufferLength = 1024 // limit minimum length of buffer
const MaximumErrorLength = 16    // limit maximum length of error

func NewJsonEncoder(out io.Writer, size int) *JsonEncoder {
	if size < MinimumBufferLength {
		size = MinimumBufferLength
	}
	return &JsonEncoder{
		out:  out,
		buff: make([]byte, 0, size),
	}
}

type JsonEncoder struct {
	out        io.Writer
	buff       []byte
	number     [32]byte // 数值缓存区
	firstError error    // 上下文错误
}

func (w *JsonEncoder) reportError(err error) {
	if w.firstError == nil {
		w.firstError = err
	}
}

func (w *JsonEncoder) Write(p []byte) (int, error) {
	n := len(p)
	w.ensure(n)
	w.buff = append(w.buff, p...)
	return n, nil
}

var _ io.Writer = (*JsonEncoder)(nil)

func (w *JsonEncoder) Reset(out io.Writer) *JsonEncoder {
	w.out = out
	w.buff = w.buff[0:0]
	return w
}

func (w *JsonEncoder) Clean() *JsonEncoder {
	w.out = nil
	w.firstError = nil
	return w
}

func (w *JsonEncoder) ensure(n int) {
	if w.out != nil && cap(w.buff)-len(w.buff) < n {
		_, err := w.out.Write(w.buff)
		if err != nil {
			if w.firstError == nil {
				w.firstError = err
			}
			return
		}
		w.buff = w.buff[0:0]
	}
}

// Close 关闭写流, 并返回剩余buff
func (w *JsonEncoder) Close() error {
	if w.firstError != nil {
		return w.firstError
	}
	if w.out != nil && len(w.buff) > 0 {
		_, err := w.out.Write(w.buff)
		return err
	}
	return nil
}

func (w *JsonEncoder) Buffer() []byte {
	return w.buff
}

func (w *JsonEncoder) escape(s string) {

	// Portions of the string that contain no escapes are appended as
	// byte slices.

	p := 0 // last non-escape symbol

	for i := 0; i < len(s); {
		c := s[i]

		if c < utf8.RuneSelf {
			if noEscapeHtmlTable[c] {
				// single-width character, no escaping is required
				i++
				continue
			}

			w.buff = append(w.buff, s[p:i]...)
			switch c {
			case '\t':
				w.buff = append(w.buff, `\t`...)
			case '\r':
				w.buff = append(w.buff, `\r`...)
			case '\n':
				w.buff = append(w.buff, `\n`...)
			case '\\':
				w.buff = append(w.buff, `\\`...)
			case '"':
				w.buff = append(w.buff, `\"`...)
			default:
				w.buff = append(w.buff, '\\', 'u', '0', '0', chars[c>>4], chars[c&0xf])
			}

			i++
			p = i
			continue
		}

		// broken utf
		runeValue, runeWidth := utf8.DecodeRuneInString(s[i:])
		if runeValue == utf8.RuneError && runeWidth == 1 {
			w.buff = append(w.buff, s[p:i]...)
			w.buff = append(w.buff, '\\', 'u', 'f', 'f', 'f', 'd')
			i++
			p = i
			continue
		}

		// jsonp stuff - tab separator and line separator
		if runeValue == '\u2028' || runeValue == '\u2029' {
			w.buff = append(w.buff, s[p:i]...)
			w.buff = append(w.buff, '\\', 'u', '2', '0', '2', chars[runeValue&0xf])
			i += runeWidth
			p = i
			continue
		}
		i += runeWidth
	}
	w.buff = append(w.buff, s[p:]...)
}

func (w *JsonEncoder) base64(in []byte) {
	si := 0
	n := (len(in) / 3) * 3

	for si < n {
		// Convert 3x 8bit source bytes into 4 bytes
		val := uint(in[si+0])<<16 | uint(in[si+1])<<8 | uint(in[si+2])
		w.buff = append(w.buff, encode[val>>18&0x3F], encode[val>>12&0x3F], encode[val>>6&0x3F], encode[val&0x3F])
		si += 3
	}

	remain := len(in) - si
	if remain == 0 {
		return
	}

	// Add the remaining small block
	val := uint(in[si+0]) << 16
	if remain == 2 {
		val |= uint(in[si+1]) << 8
	}

	w.buff = append(w.buff, encode[val>>18&0x3F], encode[val>>12&0x3F])

	switch remain {
	case 2:
		w.buff = append(w.buff, encode[val>>6&0x3F], byte(padChar))
	case 1:
		w.buff = append(w.buff, byte(padChar), byte(padChar))
	}
}

/**************************************************
* 复用方法
**************************************************/

func encodeNull(w *JsonEncoder) {
	w.ensure(4)
	w.buff = append(w.buff, 'n', 'u', 'l', 'l')
}

func encodeTrue(w *JsonEncoder) {
	w.ensure(4)
	w.buff = append(w.buff, 't', 'r', 'u', 'e')
}

func encodeFalse(w *JsonEncoder) {
	w.ensure(5)
	w.buff = append(w.buff, 'f', 'a', 'l', 's', 'e')
}

func encodeZero(w *JsonEncoder) {
	w.ensure(1)
	w.buff = append(w.buff, '0')
}

func encodeInt32(w *JsonEncoder, value int32) {
	w.ensure(11)
	w.buff = append(w.buff, strconv.AppendInt(w.number[0:0], int64(value), 10)...)
}

func encodeInt64(w *JsonEncoder, value int64) {
	w.ensure(21)
	w.buff = append(w.buff, strconv.AppendInt(w.number[0:0], value, 10)...)
}

func encodeUint32(w *JsonEncoder, value uint32) {
	w.ensure(11)
	w.buff = append(w.buff, strconv.AppendUint(w.number[0:0], uint64(value), 10)...)
}

func encodeUint64(w *JsonEncoder, value uint64) {
	w.ensure(21)
	w.buff = append(w.buff, strconv.AppendUint(w.number[0:0], value, 10)...)
}

func encodeFloat(w *JsonEncoder, value float32) {
	w.ensure(21)
	w.buff = append(w.buff, strconv.AppendFloat(w.number[0:0], float64(value), 'g', -1, 32)...)
}

func encodeDouble(w *JsonEncoder, value float64) {
	w.ensure(21)
	w.buff = append(w.buff, strconv.AppendFloat(w.number[0:0], value, 'g', -1, 64)...)
}

func encodeString(w *JsonEncoder, value string) {
	w.ensure(2 + len(value))
	w.buff = append(w.buff, quotes)
	w.escape(value)
	w.buff = append(w.buff, quotes)
}

func encodeBytes(w *JsonEncoder, value []byte) {
	w.ensure(2 + base64.StdEncoding.EncodedLen(len(value)))
	w.buff = append(w.buff, quotes)
	w.base64(value)
	w.buff = append(w.buff, quotes)
}

func encodeStringEmpty(w *JsonEncoder) {
	w.ensure(2)
	w.buff = append(w.buff, quotes, quotes)
}

func encodeArrayEmpty(w *JsonEncoder) {
	w.ensure(2)
	w.buff = append(w.buff, leftBracket, rightBracket)
}

func encodeObject(w *JsonEncoder, value any) {
	if e, ok := value.(FieldCodec); ok {
		w.ensure(2)
		w.buff = append(w.buff, leftBrace)
		e.EncodeField(w)
		if last := len(w.buff) - 1; w.buff[last] == comma {
			w.buff[last] = rightBrace
		} else {
			w.buff = append(w.buff, rightBrace)
		}
	} else {
		bs, err := EncodeJson(value)
		if err != nil {
			w.reportError(err)
		} else {
			_, _ = w.Write(bs)
		}
	}

}

func encodeObjectEmpty(w *JsonEncoder) {
	w.ensure(2)
	w.buff = append(w.buff, leftBrace, rightBrace)
}

func encodeNullMember(w *JsonEncoder, name string) {
	w.ensure(8 + len(name))
	w.buff = append(w.buff, quotes)
	w.buff = append(w.buff, name...)
	w.buff = append(w.buff, quotes, colon, 'n', 'u', 'l', 'l', comma)
}

func encodeTrueMember(w *JsonEncoder, name string) {
	w.ensure(8 + len(name))
	w.buff = append(w.buff, quotes)
	w.buff = append(w.buff, name...)
	w.buff = append(w.buff, quotes, colon, 't', 'r', 'u', 'e', comma)
}

func encodeFalseMember(w *JsonEncoder, name string) {
	w.ensure(9 + len(name))
	w.buff = append(w.buff, quotes)
	w.buff = append(w.buff, name...)
	w.buff = append(w.buff, quotes, colon, 'f', 'a', 'l', 's', 'e', comma)
}

func encodeZeroMember(w *JsonEncoder, name string) {
	w.ensure(5 + len(name))
	w.buff = append(w.buff, quotes)
	w.buff = append(w.buff, name...)
	w.buff = append(w.buff, quotes, colon, '0', comma)
}

func encodeInt32Member(w *JsonEncoder, name string, value int32) {
	w.ensure(15 + len(name))
	w.buff = append(w.buff, quotes)
	w.buff = append(w.buff, name...)
	w.buff = append(w.buff, quotes, colon)
	w.buff = append(w.buff, strconv.AppendInt(w.number[0:0], int64(value), 10)...)
	w.buff = append(w.buff, comma)
}

func encodeInt64Member(w *JsonEncoder, name string, value int64) {
	w.ensure(25 + len(name))
	w.buff = append(w.buff, quotes)
	w.buff = append(w.buff, name...)
	w.buff = append(w.buff, quotes, colon)
	w.buff = append(w.buff, strconv.AppendInt(w.number[0:0], value, 10)...)
	w.buff = append(w.buff, comma)
}

func encodeUint32Member(w *JsonEncoder, name string, value uint32) {
	w.ensure(15 + len(name))
	w.buff = append(w.buff, quotes)
	w.buff = append(w.buff, name...)
	w.buff = append(w.buff, quotes, colon)
	w.buff = append(w.buff, strconv.AppendUint(w.number[0:0], uint64(value), 10)...)
	w.buff = append(w.buff, comma)
}

func encodeUint64Member(w *JsonEncoder, name string, value uint64) {
	w.ensure(25 + len(name))
	w.buff = append(w.buff, quotes)
	w.buff = append(w.buff, name...)
	w.buff = append(w.buff, quotes, colon)
	w.buff = append(w.buff, strconv.AppendUint(w.number[0:0], value, 10)...)
	w.buff = append(w.buff, comma)
}

func encodeFloatMember(w *JsonEncoder, name string, value float32) {
	w.ensure(25 + len(name))
	w.buff = append(w.buff, quotes)
	w.buff = append(w.buff, name...)
	w.buff = append(w.buff, quotes, colon)
	w.buff = append(w.buff, strconv.AppendFloat(w.number[0:0], float64(value), 'g', -1, 32)...)
	w.buff = append(w.buff, comma)
}

func encodeDoubleMember(w *JsonEncoder, name string, value float64) {
	w.ensure(25 + len(name))
	w.buff = append(w.buff, quotes)
	w.buff = append(w.buff, name...)
	w.buff = append(w.buff, quotes, colon)
	w.buff = append(w.buff, strconv.AppendFloat(w.number[0:0], value, 'g', -1, 64)...)
	w.buff = append(w.buff, comma)
}

func encodeStringMember(w *JsonEncoder, name string, value string) {
	w.ensure(4 + len(name) + len(value))
	w.buff = append(w.buff, quotes)
	w.buff = append(w.buff, name...)
	w.buff = append(w.buff, quotes, colon, quotes)
	w.escape(value)
	w.buff = append(w.buff, quotes, comma)
}

func encodeStringEmptyMember(w *JsonEncoder, name string) {
	w.ensure(6 + len(name))
	w.buff = append(w.buff, quotes)
	w.buff = append(w.buff, name...)
	w.buff = append(w.buff, quotes, colon, quotes, quotes, comma)
}

func encodeBytesMember(w *JsonEncoder, name string, value []byte) {
	w.ensure(6 + len(name) + base64.StdEncoding.EncodedLen(len(value)))
	w.buff = append(w.buff, quotes)
	w.buff = append(w.buff, name...)
	w.buff = append(w.buff, quotes, colon, quotes)
	w.base64(value)
	w.buff = append(w.buff, quotes, comma)
}

func encodeArrayEmptyMember(w *JsonEncoder, name string) {
	w.ensure(6 + len(name))
	w.buff = append(w.buff, quotes)
	w.buff = append(w.buff, name...)
	w.buff = append(w.buff, quotes, colon, leftBracket, rightBracket, comma)
}

func encodeObjectMember(w *JsonEncoder, name string, value any) {
	w.ensure(4 + len(name))
	w.buff = append(w.buff, quotes)
	w.buff = append(w.buff, name...)
	w.buff = append(w.buff, quotes, colon)
	encodeObject(w, value)
	w.buff = append(w.buff, comma)
}

func encodeObjectEmptyMember(w *JsonEncoder, name string) {
	w.ensure(6 + len(name))
	w.buff = append(w.buff, quotes)
	w.buff = append(w.buff, name...)
	w.buff = append(w.buff, quotes, colon, leftBrace, rightBrace, comma)
}

func encodeObjectWith[V any](w *JsonEncoder, value map[string]V, enc func(w *JsonEncoder, v V)) {
	w.ensure(2)
	w.buff = append(w.buff, leftBrace)
	for k, v := range value {
		encodeString(w, k)
		w.buff = append(w.buff, colon)
		enc(w, v)
		w.buff = append(w.buff, comma)
	}
	w.buff[len(w.buff)-1] = rightBrace
}

func encodeObjectMemberWith[V any](w *JsonEncoder, name string, value map[string]V, enc func(w *JsonEncoder, v V)) {
	w.ensure(4 + len(name))
	w.buff = append(w.buff, quotes)
	w.buff = append(w.buff, name...)
	w.buff = append(w.buff, quotes, colon, leftBrace)
	for k, v := range value {
		encodeString(w, k)
		w.buff = append(w.buff, colon)
		enc(w, v)
		w.buff = append(w.buff, comma)
	}
	w.buff[len(w.buff)-1] = rightBrace
	w.buff = append(w.buff, comma)
}

func encodeArrayWith[V any](w *JsonEncoder, value []V, enc func(w *JsonEncoder, v V)) {
	w.ensure(2)
	w.buff = append(w.buff, leftBracket)
	for _, v := range value {
		enc(w, v)
		w.buff = append(w.buff, comma)
	}
	w.buff[len(w.buff)-1] = rightBracket
}

func encodeArrayMemberWith[V any](w *JsonEncoder, name string, value []V, enc func(w *JsonEncoder, v V)) {
	w.ensure(4 + len(name))
	w.buff = append(w.buff, quotes)
	w.buff = append(w.buff, name...)
	w.buff = append(w.buff, quotes, colon, leftBracket)
	for _, v := range value {
		enc(w, v)
		w.buff = append(w.buff, comma)
	}
	w.buff[len(w.buff)-1] = rightBracket
	w.buff = append(w.buff, comma)
}

/**********************************************************
 * encoder常量与亦是
 **********************************************************/

const (
	leftBrace    byte = '{'
	rightBrace   byte = '}'
	leftBracket  byte = '['
	rightBracket byte = ']'
	comma        byte = ','
	colon        byte = ':'
	quotes       byte = '"'
)

const encode = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
const padChar = '='

const chars = "0123456789abcdef"

func table(falseValues ...int) [128]bool {
	ret := [128]bool{}

	for i := 0; i < 128; i++ {
		ret[i] = true
	}

	for _, v := range falseValues {
		ret[v] = false
	}

	return ret
}

var (
	//escapeHtmlTable   = table(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, '"', '&', '<', '>', '\\')
	noEscapeHtmlTable = table(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, '"', '\\')
)
