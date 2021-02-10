package main

import "fmt"

func main() {
	//test function
	x := 0
	increment := func() int {
		x++
		return x

	}
	fmt.Println(increment())
}
