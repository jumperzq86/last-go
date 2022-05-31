package slice

import (
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func lastNumsBySlice(origin []int) []int {
	return origin[len(origin)-2:]
}

func lastNumsByCopy(origin []int) []int {
	c := make([]int, 2)
	copy(c, origin[len(origin)-2:])
	return c
}

func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func printMem(t *testing.T) {
	t.Helper()
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	t.Logf("%.2f MB", float64(rtm.Alloc)/1024./1024.)
}

func testLastChars(t *testing.T, f func([]int) []int) {
	t.Helper()
	ans := make([][]int, 0)
	for i := 0; i < 100; i++ {
		origin := generateWithCap(128 * 1024)
		ans = append(ans, f(origin))
	}
	printMem(t)
	_ = ans
}
