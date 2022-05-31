package string_concat

import "testing"

func benchmark(b *testing.B, f func(int, string) string) {
	str := randomString(10)
	for i := 0; i < b.N; i++ {
		f(10000, str)
	}
}

func BenchmarkPlusConcat(b *testing.B) {
	benchmark(b, plusConcat)
}

func BenchmarkSprintfConcat(b *testing.B) {
	benchmark(b, sprintfConcat)
}

func BenchmarkBuilderConcat(b *testing.B) {
	benchmark(b, builderConcat)
}

func BenchmarkBufferConcat(b *testing.B) {
	benchmark(b, bufferConcat)
}

func BenchmarkByteConcat(b *testing.B) {
	benchmark(b, byteConcat)
}

func BenchmarkPreBuilderConcat(b *testing.B) {
	benchmark(b, preBuilderConcat)
}

func BenchmarkPreByteConcat(b *testing.B) {
	benchmark(b, preByteConcat)
}
