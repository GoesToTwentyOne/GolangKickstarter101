package main

import "fmt"

func main() {
	input, inputbool := half(2)
	fmt.Printf("(%v  %v)", input, inputbool)

}

func half(a int) (int64, bool) {
	b := a * 1 / 2
	x := int64(b)
	var z bool
	if a%2 == 0 {
		z = true
	} else {
		z = false

	}

	return x, z

}
