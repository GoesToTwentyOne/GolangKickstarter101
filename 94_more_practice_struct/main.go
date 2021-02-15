package main

import (
	"fmt"
	"math"
)

func main() {
	c := circle{0, 0, 58}
	fmt.Println(c)
	fmt.Println(c.area())
	//fmt.Println(circlearea(&c))
	fmt.Println(c.area())
	// c = circle{5, 3, 2}
	// fmt.Println(c)

}

type circle struct {
	x, y, r float64
}

///this is method

func (c *circle) area() float64 {
	return math.Pi * c.r * c.r

}
