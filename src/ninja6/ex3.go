package main

import "fmt"

func main() {
	defer printMe("First")
	printMe("Second")
}

func printMe(s string) {
	fmt.Println(s)
}
