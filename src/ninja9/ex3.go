package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	count := 0
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			v := count
			runtime.Gosched()
			v++
			count = v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
