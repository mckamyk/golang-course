package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	defer wg.Wait()

	go foo("Go 1", true)
	go foo("Go 2", true)
	foo("main", false)
}

func foo(s string, isRoutine bool) {
	fmt.Println(s)
	if isRoutine {
		wg.Done()
	}
}
