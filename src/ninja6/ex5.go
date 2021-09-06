package main

import "fmt"

type circle struct {
	radius int
}

type square struct {
	length int
}

func (c circle) area() float64 {
	pi := 3.14159
	return pi * float64(c.radius) * 2
}

func (s square) area() float64 {
	return float64(s.length * s.length)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c := circle{radius: 4}
	s := square{length: 12}

	info(c)
	info(s)
}
