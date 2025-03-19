package base

import (
	"fmt"
	"io"
	"strconv"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

/*
JsonDecoder 核心结构, 实现struct的解码.
*/

const MaximumNestingDepth = 256 // limit maximum depth of nesting, as allowed by https://tools.ietf.org/html/rfc7159#section-9

func NewJsonDecoder(in io.Reader, size int) *JsonDecoder {
	if size < MinimumBufferLength {
		size = MinimumBufferLength
	}
	r := &JsonDecoder{
		in:   in,
		buff: make([]byte, size),
		mark: 0,
		size: 0,
	}
	r.next()
	return r
}

func NewJsonBuffer(buf []byte) *JsonDecoder {
	r := &JsonDecoder{
		in:   nil,
		buff: buf,
		mark: 0,
		size: len(buf),
	}
	r.next()
	return r
}

type JsonDecoder struct {
	in         io.Reader // 读入流
	buff       []byte    // 缓存区
	mark       int       // 读位置(+1)
	size       int       // 末位置
	base       int       // 基位置(base + pos)
	token      JsonToken // token类型
	depth      int       // 嵌套深度
	number     []byte    // 数值缓存区
	firstError error     // 上下文错误
}

func (r *JsonDecoder) Read(p []byte) (int, error) {
	if r.mark < r.size || r.more() {
		n := copy(p, r.buff[r.mark:r.size])
		r.mark += n
		return n, nil
	}
	return 0, io.EOF
}

var _ io.Reader = (*JsonDecoder)(nil)

func (r *JsonDecoder) Reset(in io.Reader) *JsonDecoder {
	r.in = in
	r.mark = 0
	r.size = 0
	r.base = 0
	r.token = 0
	r.depth = 0
	r.next()
	return r
}

func (r *JsonDecoder) Clean() *JsonDecoder {
	r.in = nil
	r.firstError = nil
	return r
}

// Close 关闭读流
func (r *JsonDecoder) Close() error {
	// 解析结束,不能再有其他字符
	if r.next() != 0 {
		r.unexpectedEndError()
	}
	return r.firstError
}

func (r *JsonDecoder) expectedTokenError(t JsonToken) {
	if r.firstError == nil {
		r.token = -1
		r.firstError = newParseError(r, r.mark-1, fmt.Sprintf("expected %s, but got '%c'", tokens[t], r.buff[r.mark-1]))
	}
}

func (r *JsonDecoder) expectedDelimiterErrorMark(mark int, c byte) {
	if r.firstError == nil {
		r.token = -1
		if c < '\u0020' {
			r.firstError = newParseError(r, mark, fmt.Sprintf("expected <delimiter>, but got '\\x%x'", c))
		} else {
			r.firstError = newParseError(r, mark, fmt.Sprintf("expected <delimiter>, but got '%c'", c))
		}
	}
}

func (r *JsonDecoder) expectedDelimiterError() {
	idx := r.mark - 1
	if idx < 0 {
		idx = 0
	}
	r.expectedDelimiterErrorMark(idx, r.buff[idx])
}

func (r *JsonDecoder) invalidCharacterErrorMark(mark int, c byte) {
	if r.firstError == nil {
		r.token = -1
		if c < '\u0020' {
			r.firstError = newParseError(r, mark, fmt.Sprintf("invalid character '\\x%x'", c))
		} else {
			r.firstError = newParseError(r, mark, fmt.Sprintf("invalid character '%c'", c))
		}
	}
}

func (r *JsonDecoder) invalidCharacterError() {
	idx := r.mark - 1
	if idx < 0 {
		idx = 0
	}
	r.invalidCharacterErrorMark(idx, r.buff[idx])
}

func (r *JsonDecoder) unexpectedEndError() {
	if r.firstError == nil {
		r.token = -1
		r.firstError = newParseError(r, r.mark-1, "unexpected end of JSON input")
	}
}

func (r *JsonDecoder) exceedMaximumNestingDepthError() {
	if r.firstError == nil {
		r.token = -1
		r.firstError = newParseError(r, r.mark-1, "exceed maximum depth of nesting")
	}
}

func (r *JsonDecoder) reportError(err error) {
	if r.firstError == nil {
		r.token = -1
		r.firstError = err
	}
}

// more 从读入流读取缓存区
func (r *JsonDecoder) more() bool {
	if r.in == nil {
		return false
	}
	r.base += r.size
	r.mark = 0
	r.size = 0
	for r.size < cap(r.buff) {
		n, err := r.in.Read(r.buff[r.size:])
		r.size += n
		if err != nil {
			if r.size > 0 {
				return true
			}
			if err != io.EOF {
				r.reportError(err)
			}
			return false
		}
	}
	return true
}

func (r *JsonDecoder) next() JsonToken {
	if r.token != 0 {
		switch r.token {
		case ObjectBegin:
			r.skipObject()
		case ObjectEnd:
			r.token = 0
		case ArrayBegin:
			r.skipArray()
		case ArrayEnd:
			r.token = 0
		case Comma:
			r.token = 0
		case Colon:
			r.token = 0
		case String:
			r.skipString()
		case Number:
			r.skipNumber()
		case True:
			r.skipTrue()
		case False:
			r.skipFalse()
		case Null:
			r.skipNull()
		case -1:
			return -1
		}
	}

	for {
		for i, c := range r.buff[r.mark:r.size] {
			switch c {
			case '\u0020', '\u000A', '\u000D', '\u0009':
			case '{':
				if r.depth++; r.depth > MaximumNestingDepth {
					r.exceedMaximumNestingDepthError()
					r.depth--
					return -1
				}
				r.mark += i + 1 // mark永远指向下次读位置
				r.token = ObjectBegin
				return ObjectBegin
			case '}':
				r.depth--
				r.mark += i + 1 // mark永远指向下次读位置
				r.token = ObjectEnd
				return ObjectEnd
			case '[':
				if r.depth++; r.depth > MaximumNestingDepth {
					r.exceedMaximumNestingDepthError()
					r.depth--
					return 0
				}
				r.mark += i + 1 // mark永远指向下次读位置
				r.token = ArrayBegin
				return ArrayBegin
			case ']':
				r.depth--
				r.mark += i + 1 // mark永远指向下次读位置
				r.token = ArrayEnd
				return ArrayEnd
			case ',':
				r.mark += i + 1 // mark永远指向下次读位置
				r.token = Comma
				return Comma
			case ':':
				r.mark += i + 1 // mark永远指向下次读位置
				r.token = Colon
				return Colon
			case '"':
				r.mark += i + 1 // mark永远指向下次读位置
				r.token = String
				return String
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
				r.mark += i + 1 // mark永远指向下次读位置
				r.number = append(r.number[0:0], c)
				r.token = Number
				return Number
			case 't':
				r.mark += i + 1 // mark永远指向下次读位置
				r.token = True
				return True
			case 'f':
				r.mark += i + 1 // mark永远指向下次读位置
				r.token = False
				return False
			case 'n':
				r.mark += i + 1 // mark永远指向下次读位置
				r.token = Null
				return Null
			default:
				r.mark += i + 1 // mark永远指向下次读位置
				r.invalidCharacterError()
				return -1
			}
		}
		if !r.more() {
			return 0
		}
	}
}

func (r *JsonDecoder) readObject(fc FieldCodec) {
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

		switch r.next() {
		case 0:
			r.unexpectedEndError()
			return
		case Null:
			r.skipNull()
		default:
			fc.DecodeField(r, k)
		}

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
}

func (r *JsonDecoder) skipObject() {
	r.token = 0
	token := r.next()
	for token != ObjectEnd {
		if token != String {
			r.expectedTokenError(String)
			return
		}
		token = r.next()
		if token != Colon {
			r.expectedTokenError(Colon)
			return
		}
		token = r.next()
		switch token {
		case ObjectBegin:
			r.skipObject()
		case ObjectEnd:
			r.invalidCharacterError()
			return
		case ArrayBegin:
			r.skipArray()
		case ArrayEnd:
			r.invalidCharacterError()
			return
		case Comma:
			r.invalidCharacterError()
			return
		case Colon:
			r.invalidCharacterError()
			return
		case String:
			r.skipString()
		case Number:
			r.skipNumber()
		case True:
			r.skipTrue()
		case False:
			r.skipFalse()
		case Null:
			r.skipNull()
		case 0:
			r.unexpectedEndError()
			return
		case -1:
			return
		}
		token = r.next()
		if token == Comma {
			token = r.next()
			if token == ObjectEnd { // ",}"是错误格式
				r.invalidCharacterError()
				return
			}
		}
	}
	return
}

func (r *JsonDecoder) skipArray() {
	r.token = 0

	token := r.next()
	for token != ArrayEnd {
		switch token {
		case ObjectBegin:
			r.skipObject()
		case ObjectEnd:
			r.invalidCharacterError()
			return
		case ArrayBegin:
			r.skipArray()
		case ArrayEnd:
			r.invalidCharacterError()
			return
		case Comma:
			r.invalidCharacterError()
			return
		case Colon:
			r.invalidCharacterError()
			return
		case String:
			r.skipString()
		case Number:
			r.skipNumber()
		case True:
			r.skipTrue()
		case False:
			r.skipFalse()
		case Null:
			r.skipNull()
		case 0:
			r.unexpectedEndError()
			return
		case -1:
			return
		}
		token = r.next()
		if token == Comma {
			token = r.next()
			if token == ArrayEnd { // ",]"是错误格式
				r.invalidCharacterError()
				return
			}
		}
	}
	return
}

func (r *JsonDecoder) readInt64() int64 {
	r.readNumber()
	ret, err := strconv.ParseInt(UnsafeString(r.number), 10, 64)
	if err != nil && r.firstError == nil {
		r.firstError = err
	}
	return ret
}

func (r *JsonDecoder) readUint64() uint64 {
	r.readNumber()
	ret, err := strconv.ParseUint(UnsafeString(r.number), 10, 64)
	if err != nil && r.firstError == nil {
		r.firstError = err
	}
	return ret
}

func (r *JsonDecoder) readFloat64() float64 {
	r.readNumber()
	ret, err := strconv.ParseFloat(UnsafeString(r.number), 64)
	if err != nil && r.firstError == nil {
		r.firstError = err
	}
	return ret
}

func (r *JsonDecoder) readByte() byte {
	if r.mark < r.size || r.more() {
		c := r.buff[r.mark]
		r.mark++
		return c
	}
	return 0
}

func (r *JsonDecoder) unreadByte() {
	r.mark--
}

func (r *JsonDecoder) readU4() rune {
	var ret rune
	for i := 0; i < 4; i++ {
		c := r.readByte()
		switch {
		case c >= '0' && c <= '9':
			ret = ret*16 + rune(c-'0')
		case c >= 'a' && c <= 'f':
			ret = ret*16 + rune(c-'a'+10)
		case c >= 'A' && c <= 'F':
			ret = ret*16 + rune(c-'A'+10)
		default:
			if c != 0 {
				r.invalidCharacterError()
			} else {
				r.unexpectedEndError()
			}
			return 0
		}
	}
	return ret
}

func (r *JsonDecoder) appendEscape(str []byte, c byte) []byte {
	/*
		https://www.json.org/json-en.html
			escape
			'"'
			'\'
			'/'
			'b'
			'f'
			'n'
			'r'
			't'
			'u' hex hex hex hex
	*/
	switch c {
	case '"':
		str = append(str, '"')
	case '\\':
		str = append(str, '\\')
	case '/':
		str = append(str, '/')
	case 'b':
		str = append(str, '\b')
	case 'f':
		str = append(str, '\f')
	case 'n':
		str = append(str, '\n')
	case 'r':
		str = append(str, '\r')
	case 't':
		str = append(str, '\t')
	case 'u':
		u1 := r.readU4()
		if utf16.IsSurrogate(u1) {
			c = r.readByte()
			if c == 0 {
				return utf8.AppendRune(str, u1)
			}
			if c != '\\' {
				r.unreadByte()
				return utf8.AppendRune(str, u1)
			}
			c = r.readByte()
			if c == 0 {
				return utf8.AppendRune(str, u1)
			}
			if c != 'u' {
				str = utf8.AppendRune(str, u1)
				return r.appendEscape(str, c)
			}
			u2 := r.readU4()
			combined := utf16.DecodeRune(u1, u2)
			if combined == unicode.ReplacementChar {
				str = utf8.AppendRune(str, u1)
				str = utf8.AppendRune(str, u2)
			} else {
				str = utf8.AppendRune(str, combined)
			}
		} else {
			str = utf8.AppendRune(str, u1)
		}
	default:
		if c != 0 {
			r.invalidCharacterError()
		} else {
			r.unexpectedEndError()
		}
	}
	return str
}

func (r *JsonDecoder) skipString() {
	r.token = 0

	isEscape := false
	for {
		for i, c := range r.buff[r.mark:r.size] {
			switch {
			case c < '\u0020':
				r.invalidCharacterErrorMark(r.mark+i, c) // 特殊处理: 控制字符继续
			case c == '"' && !isEscape:
				r.mark += i + 1 // mark永远指向下次读位置
				if !isTokenEnd(r.readByte()) {
					r.expectedDelimiterError()
				}
				r.unreadByte()
				return
			case c == '\\' && !isEscape:
				isEscape = true
			default:
				isEscape = false
			}
		}
		if !r.more() {
			r.unexpectedEndError()
			return
		}
	}
}

func (r *JsonDecoder) readString() string {
	r.token = 0

	var buf []byte
__ENTRY__:
	for {
		for i, c := range r.buff[r.mark:r.size] {
			switch {
			case c < '\u0020':
				r.invalidCharacterErrorMark(r.mark+i, c) // 记录错误后继续
			case c == '"':
				mark := r.mark
				r.mark += i + 1 // mark永远指向下次读位置
				buf = append(buf, r.buff[mark:r.mark-1]...)
				if !isTokenEnd(r.readByte()) {
					r.expectedDelimiterError()
				}
				r.unreadByte()
				return UnsafeString(buf)
			case c == '\\':
				mark := r.mark
				r.mark += i + 1
				buf = append(buf, r.buff[mark:r.mark-1]...)
				buf = r.appendEscape(buf, r.readByte())
				continue __ENTRY__
			}
		}
		buf = append(buf, r.buff[r.mark:r.size]...)
		if !r.more() {
			r.unexpectedEndError()
			return UnsafeString(buf)
		}
	}
}

func (r *JsonDecoder) skipNumber() {
	r.token = 0

	hasE := false
	afterE := false
	hasDot := false
	for {
		for i, c := range r.buff[r.mark:r.size] {
			switch {
			case c >= '0' && c <= '9':
				afterE = false
			case c == '.' && !hasDot:
				hasDot = true
			case (c == 'e' || c == 'E') && !hasE:
				hasE = true
				hasDot = true
				afterE = true
			case (c == '+' || c == '-') && afterE:
				afterE = false
			default:
				r.mark += i // mark永远指向下次读位置
				if !isTokenEnd(c) {
					r.expectedDelimiterErrorMark(r.mark, c)
				}
				return
			}
		}
		if !r.more() {
			return
		}
	}
}

func (r *JsonDecoder) readNumber() {
	r.token = 0

	hasE := false
	afterE := false
	hasDot := false
	for {
		for i, c := range r.buff[r.mark:r.size] {
			switch {
			case c >= '0' && c <= '9':
				afterE = false
			case c == '.' && !hasDot:
				hasDot = true
			case (c == 'e' || c == 'E') && !hasE:
				hasE = true
				hasDot = true
				afterE = true
			case (c == '+' || c == '-') && afterE:
				afterE = false
			default:
				mark := r.mark
				r.mark += i
				r.number = append(r.number, r.buff[mark:r.mark]...)
				if !isTokenEnd(c) {
					r.expectedDelimiterErrorMark(r.mark, c)
				}
				return
			}
		}
		r.number = append(r.number, r.buff[r.mark:r.size]...)
		if !r.more() {
			return
		}
	}
}

func (r *JsonDecoder) skipTrue() {
	r.token = 0
	// 跳过't'
	if r.readByte() != 'r' {
		r.expectedTokenError(True)
	}
	if r.readByte() != 'u' {
		r.expectedTokenError(True)
	}
	if r.readByte() != 'e' {
		r.expectedTokenError(True)
	}
	if !isTokenEnd(r.readByte()) {
		r.expectedDelimiterError()
	}
	r.unreadByte()
}

func (r *JsonDecoder) skipFalse() {
	r.token = 0
	//  跳过'f'
	if r.readByte() != 'a' {
		r.expectedTokenError(True)
	}
	if r.readByte() != 'l' {
		r.expectedTokenError(True)
	}
	if r.readByte() != 's' {
		r.expectedTokenError(True)
	}
	if r.readByte() != 'e' {
		r.expectedTokenError(True)
	}
	if !isTokenEnd(r.readByte()) {
		r.expectedDelimiterError()
	}
	r.unreadByte()
}

func (r *JsonDecoder) skipNull() {
	r.token = 0
	// 跳过'n'
	if r.readByte() != 'u' {
		r.expectedTokenError(Null)
	}
	if r.readByte() != 'l' {
		r.expectedTokenError(Null)
	}
	if r.readByte() != 'l' {
		r.expectedTokenError(Null)
	}
	if !isTokenEnd(r.readByte()) {
		r.expectedDelimiterError()
	}
	r.unreadByte()
}

// dumpData, dump数据
func (r *JsonDecoder) dumpObjectOrArray(token JsonToken) []byte {

	var start, end byte
	switch token {
	case ObjectBegin:
		start, end = '{', '}'
	case ArrayBegin:
		start, end = '[', ']'
	default:
		return nil
	}

	r.token = 0 // 很关键!

	// 先添加start
	var data = append(make([]byte, 0, 256), start)

	level := 1
	inQuotes := false
	wasEscape := false
	for {
		for i, c := range r.buff[r.mark:r.size] {
			switch {
			case c == start && !inQuotes:
				level++
			case c == end && !inQuotes:
				level--
				if level == 0 {
					mark := r.mark
					r.mark += i + 1
					data = append(data, r.buff[mark:r.mark]...)
					return data
				}
			case c == '\\' && inQuotes:
				wasEscape = !wasEscape
				continue
			case c == '"' && inQuotes:
				inQuotes = wasEscape
			case c == '"':
				inQuotes = true
			}
			wasEscape = false
		}
		data = append(data, r.buff[r.mark:r.size]...)
		if !r.more() {
			return data
		}
	}
}

// JsonToken 词汇令牌. 0表示EOF, -1表示ERROR
type JsonToken int8

const (
	ObjectBegin JsonToken = 1
	ObjectEnd   JsonToken = 2
	ArrayBegin  JsonToken = 3
	ArrayEnd    JsonToken = 4
	Comma       JsonToken = 5
	Colon       JsonToken = 6
	String      JsonToken = 7
	Number      JsonToken = 8
	True        JsonToken = 9
	False       JsonToken = 10
	Null        JsonToken = 11
)

var tokens = map[JsonToken]string{
	ObjectBegin: `<object>`,
	ObjectEnd:   `<object>`,
	ArrayBegin:  `<array>`,
	ArrayEnd:    `<array>`,
	Comma:       `','`,
	Colon:       `':'`,
	String:      `<string>`,
	Number:      `<number>`,
	True:        `<bool>`,
	False:       `<bool>`,
	Null:        `<null>`,
}

func isTokenEnd(c byte) bool {
	return c == 0 || c == ' ' || c == '\t' || c == '\r' || c == '\n' || c == '[' || c == ']' || c == '{' || c == '}' || c == ',' || c == ':'
}

// ParseError Json解码错误
type ParseError struct {
	Reason string
	Offset int
	Data   string
}

func (l *ParseError) Error() string {
	return fmt.Sprintf("%s near offset %v: '%s'", l.Reason, l.Offset, l.Data)
}

func newParseError(ctx *JsonDecoder, mark int, reason string) error {

	if mark < 0 {
		mark = 0
	}

	var data string
	if ctx.size-mark <= MaximumErrorLength {
		data = string(ctx.buff[mark:ctx.size]) + "..."
	} else {
		data = string(ctx.buff[mark:mark+MaximumErrorLength]) + "..."
	}
	return &ParseError{
		Reason: reason,
		Offset: ctx.base + mark,
		Data:   data,
	}
}
