package stringbuf

type StringBuf struct {
	buf          [][]string // for Append()
	reverseBuf   [][]string // for Prepend()
	index        int
	reverseIndex int
	len          int
}

func (s *StringBuf) Append(strs ...string) {
	if len(s.buf) == 0 {
		s.buf = append(s.buf, []string{})
	}

	for _, str := range strs {
		if len(s.buf[s.index]) > 1023 {
			s.index++
			if len(s.buf) < s.index+1 {
				s.buf = append(s.buf, make([]string, 0, 1024))
			}
		}
		s.len += len(str)
		s.buf[s.index] = append(s.buf[s.index], str)
	}
}

func (s *StringBuf) AppendRune(runes ...rune) {
	if len(runes) == 0 {
		return
	}
	strs := make([]string, len(runes))
	for i, r := range runes {
		strs[i] = string(r)
	}
	s.Append(strs...)
}

func (s *StringBuf) AppendByte(bytesArr ...[]byte) {
	if len(bytesArr) == 0 {
		return
	}
	strs := make([]string, len(bytesArr))
	for i, r := range bytesArr {
		strs[i] = string(r)
	}
	s.Append(strs...)
}

func (s *StringBuf) Prepend(strs ...string) {
	if len(s.reverseBuf) == 0 {
		s.reverseBuf = append(s.reverseBuf, []string{})
	}
	for _, str := range strs {
		if len(s.reverseBuf[s.reverseIndex]) > 1023 {
			s.reverseIndex++
			if len(s.reverseBuf) < s.reverseIndex+1 {
				s.reverseBuf = append(s.reverseBuf, make([]string, 0, 1024))
			}
		}
		s.len += len(str)
		s.reverseBuf[s.reverseIndex] = append(s.reverseBuf[s.reverseIndex], str)
	}
}

func (s *StringBuf) PrependRune(runes ...rune) {
	if len(runes) == 0 {
		return
	}
	strs := make([]string, len(runes))
	for i, r := range runes {
		strs[i] = string(r)
	}
	s.Prepend(strs...)
}

func (s *StringBuf) PrependByte(bytesArr ...[]byte) {
	if len(bytesArr) == 0 {
		return
	}
	strs := make([]string, len(bytesArr))
	for i, r := range bytesArr {
		strs[i] = string(r)
	}
	s.Prepend(strs...)
}

func (s *StringBuf) String() string {
	return string(s.Bytes())
}

func (s *StringBuf) Bytes() []byte {
	var b = make([]byte, 0, s.len)

	for i := len(s.reverseBuf) - 1; i >= 0; i-- {
		for _, bytes := range s.reverseBuf[i] {
			b = append(b, []byte(bytes)...)
		}
	}

	for _, chunk := range s.buf {
		for _, bytes := range chunk {
			b = append(b, []byte(bytes)...)
		}
	}
	return b
}

func (s *StringBuf) Equal(t StringBuf) bool {
	if s.len != t.len {
		return false
	}
	if s.String() != t.String() {
		return false
	}
	return true
}

func (s *StringBuf) Reset() {
	s.buf = [][]string{}
	s.reverseBuf = [][]string{}
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
		switch _input := any(input).(type) {
		case string:
			sb.Append(_input)
		case []byte:
			sb.AppendByte(_input)
		}
	}
	return sb
}
