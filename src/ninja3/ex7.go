package main

import "fmt"

func main() {
	name := "Foo"
	if name == "Mac" {
		fmt.Println("Hi Mac!")
	} else if name == "John" {
		fmt.Println("Hi John!")
	} else {
		fmt.Println("You're not Mac or John!")
	}
}
