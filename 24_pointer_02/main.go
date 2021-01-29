package main

import "fmt"

func value(xPtr *int) {
	*xPtr = 0

}
func main() {
	x := 5
	value(&x)
	fmt.Println(x)

}
