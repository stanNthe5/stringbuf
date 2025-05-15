package stringbuf

import (
	"bytes"
	"unsafe"
)

type StringBuf struct {
	buf          [][]string // for Append()
	reverseBuf   [][]string // for Prepend()
	index        int
	reverseIndex int
	len          int
}

func (s *StringBuf) Write(p []byte) (n int, err error) {
	s.Append(string(p))
	return len(p), nil
}

func (s *StringBuf) WriteString(str string) (int, error) {
	if len(str) == 0 {
		return 0, nil
	}
	if len(s.buf) == 0 {
		s.buf = append(s.buf, []string{})
	}
	if len(s.buf[s.index]) > 1023 {
		s.index++
		if len(s.buf) < s.index+1 {
			s.buf = append(s.buf, make([]string, 0, 1024))
		}
	}
	s.buf[s.index] = append(s.buf[s.index], str)
	s.len += len(str)
	return len(str), nil
}

func (s *StringBuf) prependStr(str string) {
	if len(str) == 0 {
		return
	}
	if len(s.reverseBuf) == 0 {
		s.reverseBuf = append(s.reverseBuf, []string{})
	}
	if len(s.reverseBuf[s.reverseIndex]) > 1023 {
		s.reverseIndex++
		if len(s.reverseBuf) < s.reverseIndex+1 {
			s.reverseBuf = append(s.reverseBuf, make([]string, 0, 1024))
		}
	}
	s.len += len(str)
	s.reverseBuf[s.reverseIndex] = append(s.reverseBuf[s.reverseIndex], str)
}

func (s *StringBuf) Append(strs ...string) {
	if len(strs) == 0 {
		return
	}
	for _, str := range strs {
		s.WriteString(str)
	}
}

func (s *StringBuf) AppendRune(runes ...rune) {
	if len(runes) == 0 {
		return
	}
	for _, r := range runes {
		s.WriteString(string(r))
	}
}

func (s *StringBuf) AppendByte(bytesArr ...[]byte) {
	if len(bytesArr) == 0 {
		return
	}
	for _, r := range bytesArr {
		s.WriteString(string(r))
	}
}

func (s *StringBuf) Prepend(strs ...string) {
	if len(s.reverseBuf) == 0 {
		s.reverseBuf = append(s.reverseBuf, []string{})
	}

	for i := len(strs) - 1; i >= 0; i-- {
		if len(strs[i]) == 0 {
			continue
		}
		if len(s.reverseBuf[s.reverseIndex]) > 1023 {
			s.reverseIndex++
			if len(s.reverseBuf) < s.reverseIndex+1 {
				s.reverseBuf = append(s.reverseBuf, make([]string, 0, 1024))
			}
		}
		s.len += len(strs[i])
		s.reverseBuf[s.reverseIndex] = append(s.reverseBuf[s.reverseIndex], strs[i])
	}
}

func (s *StringBuf) PrependRune(runes ...rune) {
	if len(runes) == 0 {
		return
	}
	for i := len(runes) - 1; i >= 0; i-- {
		s.prependStr(string(runes[i]))
	}
}

func (s *StringBuf) PrependByte(bytesArr ...[]byte) {
	if len(bytesArr) == 0 {
		return
	}
	for i := len(bytesArr) - 1; i >= 0; i-- {
		s.prependStr(string(bytesArr[i]))
	}
}

func (s *StringBuf) String() string {
	if s.len == 0 {
		return ""
	}
	// safe: Bytes() returns freshly allocated, immutable data
	return unsafe.String(unsafe.SliceData(s.Bytes()), s.len)
}

func (s *StringBuf) Bytes() []byte {
	var b = make([]byte, 0, s.len)

	for i := len(s.reverseBuf) - 1; i >= 0; i-- {
		for j := len(s.reverseBuf[i]) - 1; j >= 0; j-- {
			b = append(b, s.reverseBuf[i][j]...)
		}
	}

	for _, chunk := range s.buf {
		for _, str := range chunk {
			b = append(b, str...)
		}
	}
	return b
}

func (s *StringBuf) Equal(t StringBuf) bool {
	if s.len != t.len {
		return false
	}
	return bytes.Equal(s.Bytes(), t.Bytes())
}

func (s *StringBuf) Reset() {
	if s.len == 0 {
		return
	}
	s.buf = s.buf[:0]
	s.reverseBuf = s.reverseBuf[:0]
	s.len = 0
	s.index = 0
	s.reverseIndex = 0
}

func (s *StringBuf) Len() int {
	return s.len
}

func New[T string | []byte](inputs ...T) StringBuf {
	var sb StringBuf
	for _, input := range inputs {
		switch input := any(input).(type) {
		case string:
			sb.Append(input)
		case []byte:
			sb.AppendByte(input)
		}
	}
	return sb
}
