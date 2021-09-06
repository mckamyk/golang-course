package main

import "fmt"

func main() {
	r1 := foo()
	r2, r3 := bar()
	fmt.Println(r1)
	fmt.Println(r2, r3)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 43, "Hi"
}
