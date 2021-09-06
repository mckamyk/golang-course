package main

import "fmt"

type person struct {
	first string
	last  string
}

func (p *person) changeMeMethod() {
	(*p).first = "Jo"
}

func changeMeFunc(p *person) {
	(*p).first = "Flem"
}

func main() {
	p := person{
		first: "Mac",
		last:  "Kam",
	}
	fmt.Println(p)
	p.changeMeMethod()
	fmt.Println(p)
	changeMeFunc(&p)
	fmt.Println(p)
}
