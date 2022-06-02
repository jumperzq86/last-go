package main

import (
	"sync"
	"testing"
	"time"
)

type RW interface {
	Write()
	Read()
}

const cost = 10 * time.Microsecond

type Lock struct {
	count int
	mu    sync.Mutex
}

func (this *Lock) Write() {
	this.mu.Lock()
	this.count++
	time.Sleep(cost)
	this.mu.Unlock()
}

func (this *Lock) Read() {
	this.mu.Lock()
	_ = this.count
	time.Sleep(cost)
	this.mu.Unlock()
}

type RWLock struct {
	count int
	mu    sync.RWMutex
}

func (this *RWLock) Write() {
	this.mu.Lock()
	this.count++
	time.Sleep(cost)
	this.mu.Unlock()
}

func (this *RWLock) Read() {
	this.mu.RLock()
	_ = this.count
	time.Sleep(cost)
	this.mu.RUnlock()
}

func benchmark(b *testing.B, rw RW, read int, write int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for i := 0; i < read*100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				rw.Read()
			}()
		}

		for i := 0; i < write*100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				rw.Write()
			}()
		}

		wg.Wait()
	}
}

func BenchmarkReadMore(b *testing.B) {
	benchmark(b, &Lock{}, 9, 1)
}

func BenchmarkReadMoreRW(b *testing.B) {
	benchmark(b, &RWLock{}, 9, 1)
}

func BenchmarkWriteMore(b *testing.B) {
	benchmark(b, &Lock{}, 1, 9)
}

func BenchmarkWriteMoreRW(b *testing.B) {
	benchmark(b, &RWLock{}, 1, 9)
}

func BenchmarkEqual(b *testing.B) {
	benchmark(b, &Lock{}, 5, 5)
}

func BenchmarkEqualRW(b *testing.B) {
	benchmark(b, &RWLock{}, 5, 5)
}
