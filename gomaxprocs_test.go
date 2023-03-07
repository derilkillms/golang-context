package golang_context

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			time.Sleep(3 * time.Second)
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("CPU", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Goroutine", totalGoroutine)
	group.Wait()
}
