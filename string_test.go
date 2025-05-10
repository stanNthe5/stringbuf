package stringbuf

import (
	"testing"
)

func TestString(t *testing.T) {
	sb := New("Hello ", "world,")
	sb.Append("I am ", "StringBuf")
	sb.Prepend("StringbBuf ", "testing: ")
	str := sb.String()
	expectedStr := "StringbBuf testing: Hello world,I am StringBuf"
	if sb.String() != expectedStr {
		t.Errorf("Expected \n\"%s\" \n but got \n\"%s\"", expectedStr, str)
	}
}
