package main

import "fmt"

func main() {
	s1 := seed()
	s2 := seed()

	fmt.Println("Seed 1", s1())
	fmt.Println("Seed 1", s1())
	fmt.Println("Seed 1", s1())
	fmt.Println("Seed 2", s2())
	fmt.Println("Seed 2", s2())
}

func seed() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
