package main

import "fmt"

const birthYear = 1993

func main() {
	currentYear := 2021
	i := birthYear
	for i <= currentYear {
		fmt.Println(i)
		i++
	}
}
