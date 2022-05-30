package main

import (
	"testing"
	"time"
)

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(30)
	}
}

///////
// 若是在benchmark函数中需要进行耗时的准备工作，可以使用 b.ResetTimer() 来重置计时器
//go test -bench='Fib' -benchtime=50x .

//goos: darwin
//goarch: amd64
//pkg: high-performance-go/benchmark
//cpu: Intel(R) Core(TM) i3-3220 CPU @ 3.30GHz
//BenchmarkFib-4                        50           5883139 ns/op
//BenchmarkFibWithTimer-4               50           5731431 ns/op
//PASS
//ok      high-performance-go/benchmark   6.606s

func BenchmarkFibWithTimer(b *testing.B) {
	time.Sleep(3 * time.Second)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		fib(30)
	}
}
