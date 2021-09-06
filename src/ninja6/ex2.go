package main

import "fmt"

func main() {
	xi := []int{3, 5, 2, 76, 8, 54, 3, 43}
	s := sum(xi...)
	fmt.Println(s)
	fmt.Println(sum2(xi))
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func sum2(xi []int) int {
	return sum(xi...)
}
