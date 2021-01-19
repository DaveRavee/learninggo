package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

type shape interface {
	area() float64
}

func main() {
	c := circle {
		radius: 21.2,
	}

	l := square {
		length: 12.6,
	}

	info(l)
	info(c)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

func info(s shape) {
	fmt.Println(s.area())
}