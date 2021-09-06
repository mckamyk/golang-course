package main

import "fmt"

const thisYear = 2021

const (
	_ = iota
	a = thisYear + iota
	b
	c
	d
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
