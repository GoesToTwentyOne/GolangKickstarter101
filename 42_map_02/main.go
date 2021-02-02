package main

import "fmt"

func main() {
	x := make(map[int]string)
	x[0] = "zero"
	x[1] = "one"
	x[2] = "two"
	x[3] = "three"
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])

}
