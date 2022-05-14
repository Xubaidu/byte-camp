package main

import (
	"bytes"
	"strings"
	"testing"
)

func strcat(n int, str string) {
	var s string
	for i := 0; i < n; i++ {
		s += str
	}
}

func strBuilder(n int, str string) {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
}

func byteBuilfer(n int, str string) {
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buf.WriteString(str)
	}
}

func BenchmarkStrcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strcat(1, "abcde")
	}
}

func BenchmarkStrBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strBuilder(1, "abcde")
	}
}

func BenchmarkByteBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteBuilfer(1, "abcde")
	}
}
