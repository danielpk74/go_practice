package main

import "fmt"

type shapes interface {
	getArea() float64
}

type area interface {
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
	t := triangle{
		base:   3.1,
		height: 2.3,
	}
	printArea(t)

	s := square{
		sideLength: 3.4,
	}
	printArea(s)

}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shapes) {
	fmt.Println(s.getArea())
}
