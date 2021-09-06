package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y1 := x[:5]
	y2 := x[5:]
	y3 := x[2:7]
	y4 := x[1:6]

	fmt.Println(x)
	fmt.Println(y1)
	fmt.Println(y2)
	fmt.Println(y3)
	fmt.Println(y4)
}
