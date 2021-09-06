package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Printf("%T", wg)
	wg.Add(1)
	defer wg.Wait()
	go foo("go")
	foo("main")
}

func foo(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Millisecond * 500)
	}
	wg.Done()
}
