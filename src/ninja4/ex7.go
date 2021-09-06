package main

import "fmt"

func main() {
	x := [][]string{}
	fmt.Println(x)
	x = append(x, []string{"James", "Bond", "Shaken, not stirred."})
	x = append(x, []string{"Miss", "Moneypenny", "Helloooooo, James."})

	for _, v := range x {
		for _, vv := range v {
			fmt.Println(vv)
		}
	}
}
