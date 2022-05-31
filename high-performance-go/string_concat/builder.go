package string_concat

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func plusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}

func sprintfConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}
	return s
}

func builderConcat(n int, str string) string {
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteString(str)
	}
	return b.String()
}

func bufferConcat(n int, str string) string {
	var buff = new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buff.WriteString(str)
	}
	return buff.String()
}

func byteConcat(n int, str string) string {
	b := make([]byte, 0)
	for i := 0; i < n; i++ {
		b = append(b, str...)
	}
	return string(b)
}

//////
func preByteConcat(n int, str string) string {
	b := make([]byte, 0, n*len(str))
	for i := 0; i < n; i++ {
		b = append(b, str...)
	}
	return string(b)
}

func preBuilderConcat(n int, str string) string {
	var b strings.Builder
	b.Grow(n * len(str))
	for i := 0; i < n; i++ {
		b.WriteString(str)
	}
	return b.String()
}
