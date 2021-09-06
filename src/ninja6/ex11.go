package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
}

func factorial(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
