package main

import (
	"fmt"
	"math"
	"strconv"
)

type point struct {
	x float32
	y float32
}

type square struct {
	topLeft point
	width float32
	height float32
}

type circle struct {
	center point
	radius float32
}

func (s square) area() float32 {
	return s.height * s.width
}

func (c circle) area() float32 {
	return float32(math.Pow(float64(c.radius), 2) * math.Pi);
}

type shape interface {
	area() float32
}

func info(s shape) {
	fmt.Println("area for the shape: " + strconv.FormatFloat(
		float64(s.area()), 'f',  -1, 32))
}

func main() {
	sq1 := square {
		point {
			0,
			0,
		},
		100,
		100,
	}
	cr1 := circle {
		point{
			0,
			0,
		},
		10,
	}
	info(sq1)
	info(cr1)
}