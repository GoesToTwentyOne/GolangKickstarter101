package main

import (
	"fmt"
	"math"
)

func main() {
	c := circle{0, 0, 5}
	//fmt.Printf("%+v", c)
	fmt.Println(circleArea(c))

}

type circle struct {
	x, y, r float64
}

func circleArea(c circle) float64 {
	return math.Pi * c.r * c.r
}
