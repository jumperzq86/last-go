package __range

import "testing"

type Item struct {
	id  int
	val [4096]byte
}

func BenchmarkForStruct(b *testing.B) {
	var items [1024]Item

	for i := 0; i < b.N; i++ {
		length := len(items)
		var tmp int
		for i := 0; i < length; i++ {
			tmp = items[i].id
		}
		_ = tmp
	}
}

func BenchmarkRangeIndexStruct(b *testing.B) {
	var items [1024]Item

	for i := 0; i < b.N; i++ {
		var tmp int
		for idx := range items {
			tmp = items[idx].id
		}
		_ = tmp
	}
}

func BenchmarkRangeStruct(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {
			tmp = item.id
		}
		_ = tmp
	}
}
