package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	fmt.Println("Begin")

	sq := square{sideLength: 5.5}
	tr := triangle{height: 3.4, base: 5.6}

	printArea(sq)
	printArea(tr)

	fmt.Println("End")
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println("The area is:", s.getArea())
}
