package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, p.last, "and I am", p.age)
}

func main() {
	me := person{
		first: "Mac",
		last:  "Kam",
		age:   27,
	}

	me.speak()
}
