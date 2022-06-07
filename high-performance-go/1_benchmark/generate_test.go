package main

import "testing"

/////////////////////////////////////////////
// 减少内存分配次数，以便提高性能
//go test -bench='Generate' -benchmem .

//goos: darwin
//goarch: amd64
//pkg: example
//BenchmarkGenerateWithCap-8  43  24335658 ns/op  8003641 B/op    1 allocs/op
//BenchmarkGenerate-8         33  30403687 ns/op  45188395 B/op  40 allocs/op
//PASS
//ok      example 2.121s
/////////////////////////////////////////////

func BenchmarkGenerateWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateWithCap(1000000)
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generate(1000000)
	}
}

/////////////////////////////////////////////
// 可以测试函数复杂度， O(1), O(n), O(n^2)
//go test -bench='Generate1' .

//goos: darwin
//goarch: amd64
//pkg: example
//BenchmarkGenerate1000-8            34048             34643 ns/op
//BenchmarkGenerate10000-8            4070            295642 ns/op
//BenchmarkGenerate100000-8            403           3230415 ns/op
//BenchmarkGenerate1000000-8            39          32083701 ns/op
//PASS
//ok      example 6.597s
/////////////////////////////////////////////

func benchmarkGenerate(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateWithCap(i)
	}
}

func BenchmarkGenerate1000(b *testing.B) {
	benchmarkGenerate(1000, b)
}

func BenchmarkGenerate10000(b *testing.B) {
	benchmarkGenerate(10000, b)
}

func BenchmarkGenerate100000(b *testing.B) {
	benchmarkGenerate(100000, b)
}

func BenchmarkGenerate1000000(b *testing.B) {
	benchmarkGenerate(1000000, b)
}

/////////////////////////////////////////////
// 在benchmark 函数中若是需要耗时的准备工作可以使用 b.StopTimer() / b.StartTimer() 来暂停计时器
//go test -bench='Bubble' .

/////////////////////////////////////////////

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		nums := generateWithCap(10000)
		b.StartTimer()
		bubbleSort(nums)
	}
}
