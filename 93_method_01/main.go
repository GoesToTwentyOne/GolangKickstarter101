package main

import "fmt"

func main() {
	r := rectangle{5, 10, 20, 30}
	fmt.Println(r.area())

}

type rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *rectangle) area() float64 {
	l := r.x1 * r.x2
	w := r.y1 * r.y2
	return l * w

}
