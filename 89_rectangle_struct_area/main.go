package main

import "fmt"

func main() {
	t := trangle{40, 60}
	fmt.Println(area(t))

}

type trangle struct {
	height int
	width  int
}

func area(t trangle) int {
	return t.height * t.width

}
