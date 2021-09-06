package main

import "fmt"

func main() {
	xYear := 1993
	currentYear := 2021

	for {
		if xYear > currentYear {
			break
		}
		fmt.Println(xYear)
		xYear++
	}
}
