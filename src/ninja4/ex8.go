package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being Evil", "Ice Cream", "Sunsets"},
	}

	for k, favs := range x {
		fmt.Println(k)
		for i, fav := range favs {
			fmt.Printf("\t%d: %s\n", i, fav)
		}
	}

	fmt.Println("")
}
