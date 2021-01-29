package main

import "fmt"

func demo(z *int) {
	*z = 0

}
func main() {
	y := 10
	fmt.Println(y)
	demo(&y)
	fmt.Println(y)

}
