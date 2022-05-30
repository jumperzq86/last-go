package main

import (
	"github.com/pkg/profile"
	"math/rand"
	"strings"
)

//func generate(n int) []int {
//	rand.Seed(time.Now().UnixNano())
//	nums := make([]int, 0)
//	for i := 0; i < n; i++ {
//		nums = append(nums, rand.Int())
//	}
//	return nums
//}
//func bubbleSort(nums []int) {
//	for i := 0; i < len(nums); i++ {
//		for j := 1; j < len(nums)-i; j++ {
//			if nums[j] < nums[j-1] {
//				nums[j], nums[j-1] = nums[j-1], nums[j]
//			}
//		}
//	}
//}
//
//func main() {
//	//f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
//	//defer f.Close()
//	//pprof.StartCPUProfile(f)
//	//defer pprof.StopCPUProfile()
//
//	defer profile.Start().Stop()
//
//	n := 10
//	for i := 0; i < 5; i++ {
//		nums := generate(n)
//		bubbleSort(nums)
//		n *= 10
//	}
//}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func concat(n int) string {
	var s strings.Builder
	for i := 0; i < n; i++ {
		s.WriteString(randomString(n))
	}
	return s.String()

	//s := ""
	//for i := 0; i < n; i++ {
	//	s += randomString(n)
	//}
	//return s
}

func main() {
	//runtime.SetCPUProfileRate(200)
	//defer profile.Start().Stop()
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()

	//f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	//defer f.Close()
	//runtime.SetCPUProfileRate(200)
	//pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()

	concat(1000)
}
