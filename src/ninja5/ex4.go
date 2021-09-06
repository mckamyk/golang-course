package main

import "fmt"

func main() {
	x := struct {
		field string
		value int
	}{
		field: "foo",
		value: 42,
	}

	fmt.Println(x)
}
