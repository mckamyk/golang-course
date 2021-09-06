package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	var wg sync.WaitGroup
	wg.Add(100)

	var mu sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			v := count
			v++
			count = v
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
