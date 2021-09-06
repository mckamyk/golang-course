package main

import "fmt"

func main() {
	foo(func() {
		fmt.Println("In callback")
	})
	fmt.Println("After callback")
}

func foo(cb func()) {
	fmt.Println("Before callback")
	cb()
}
