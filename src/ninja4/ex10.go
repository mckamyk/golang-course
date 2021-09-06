package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being Evil", "Ice Cream", "Sunsets"},
	}

	printMap(x)

	x["kam_mac"] = []string{"Joanna", "Coding", "Games"}
	printMap(x)

	delete(x, "no_dr")
	printMap(x)
}

func printMap(m map[string][]string) {
	for k, favs := range m {
		fmt.Println(k)
		for i, fav := range favs {
			fmt.Printf("\t%d: %s\n", i, fav)
		}
	}
}
