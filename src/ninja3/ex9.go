package main

import "fmt"

func main() {
	favSport := "Frisbee"
	switch favSport {
	case "Frisbee":
		fmt.Println("You're correct!")
	case "Soccer":
		fmt.Println("Also not bad")
	case "Footbal":
		fmt.Println("A great choice!")
	case "Cricket":
		fmt.Println("Get the fuck outta here")
	}
}
