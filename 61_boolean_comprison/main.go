package main

import "fmt"

func main() {
	var a bool
	var b bool = true
	fmt.Println(a)
	// it always print false as it automatically assign 0
	fmt.Printf("%v %v  %v\n", a && b, a || b, !b)
	//false and true =false
	//true and true =true
	//not b =false
}
