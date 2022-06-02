package exit_goroutine_expire

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func do2phases(phase1, done chan bool) {
	time.Sleep(time.Second)
	select {
	case phase1 <- true:
	default:
		return
	}
	time.Sleep(time.Second)
	done <- true
}

func timeoutFirstPhase() error {
	phase1 := make(chan bool)
	done := make(chan bool)

	go do2phases(phase1, done)

	select {
	case <-phase1:
		<-done
		return nil

	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout\n")
	}
}

func Test2phasesTimeout(t *testing.T) {
	for i := 0; i < 1000; i++ {
		timeoutFirstPhase()
	}
	time.Sleep(3 * time.Second)
	t.Log(runtime.NumGoroutine())
}
