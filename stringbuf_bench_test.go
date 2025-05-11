package stringbuf

import (
	"strings"
	"testing"
)

const sample = "abcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghij"
const times = 2000

// stringbuf
func BenchmarkStringBuf_Append(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb StringBuf
		for j := 0; j < times; j++ {
			sb.Append(sample)
		}
		_ = sb.String()
	}
}

// strings.Builder
func BenchmarkStringsBuilder_Append(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		for j := 0; j < times; j++ {
			sb.WriteString(sample)
		}
		_ = sb.String()
	}
}

func BenchmarkStringBuf_Prepend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb StringBuf
		for j := 0; j < times; j++ {
			sb.Prepend(sample)
		}
		_ = sb.String()
	}
}

func BenchmarkStringsBuilder_PrependSimulated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := ""
		for j := 0; j < times; j++ {
			result = sample + result
		}
	}
}
