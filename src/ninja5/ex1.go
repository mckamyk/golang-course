package main

import "fmt"

type person struct {
	first    string
	last     string
	iceCream []string
}

func main() {
	p1 := person{
		first:    "Mac",
		last:     "Kam",
		iceCream: []string{"Cookies and Cream", "Cookie Dough"},
	}

	p2 := person{
		first:    "Joanna",
		last:     "Kam",
		iceCream: []string{"Chocolate", "Vanilla"},
	}

	fmt.Printf("%s %s\n", p1.first, p1.last)
	for i, v := range p1.iceCream {
		fmt.Printf("\t%d: %s\n", i, v)
	}

	fmt.Printf("%s %s\n", p2.first, p2.last)
	for i, v := range p2.iceCream {
		fmt.Printf("\t%d: %s\n", i, v)
	}
}
