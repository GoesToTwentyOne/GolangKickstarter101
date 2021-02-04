package main

import "fmt"

// //Pointers are references to address locations in memory. They
// are their own data types in Golang. You should use
// a pointer when you have a large amount of data to pass around or
// when you want to change a value at an address.
func value(xPtr *int) {
	*xPtr = 0

}
func main() {
	x := 5
	value(&x)
	fmt.Println(x)

}
