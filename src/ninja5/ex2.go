package main

import "fmt"

type person struct {
	first    string
	last     string
	iceCream []string
}

func main() {
	x := map[string]person{
		"kam_mac": {
			first:    "Mac",
			last:     "Kam",
			iceCream: []string{"Cookies and Cream", "Cookie Dough"},
		},
		"kam_jo": {
			first:    "Joanna",
			last:     "Kam",
			iceCream: []string{"Chocolate", "Vanilla"},
		},
	}

	for _, v := range x {
		printPerson(v)
	}
}

func printPerson(p person) {
	fmt.Printf("%s %s\n", p.first, p.last)
	for i, v := range p.iceCream {
		fmt.Printf("\t%d: %s\n", i, v)
	}
}
