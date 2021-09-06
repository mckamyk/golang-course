package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var count int64 = 0
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			runtime.Gosched()
			fmt.Println(atomic.LoadInt64(&count))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
