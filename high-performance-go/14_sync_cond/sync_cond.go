package main

import (
	"log"
	"sync"
	"time"
)

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()

	for !done {
		c.Wait()
	}
	log.Println(name, "start reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	log.Println(name, "start writing")
	time.Sleep(time.Second)

	c.L.Lock()
	done = true
	c.L.Unlock()
	log.Println(name, "wake all")
	//for i := 0; i < 3; i++ {
	//	c.Signal()
	//	time.Sleep(time.Second)
	//}
	c.Broadcast()
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)

	write("writer", cond)

	time.Sleep(time.Second * 3)
}
