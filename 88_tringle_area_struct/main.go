package main

import "fmt"

func main() {
	t := trangle{5, 3}
	fmt.Println(area(t))

}

type trangle struct {
	width  float32
	height float32
}

func area(t trangle) float32 {
	return .5 * t.width * t.height
}
