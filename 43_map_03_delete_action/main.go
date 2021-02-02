package main

import "fmt"

func main() {
	x := make(map[int]int)
	x[1] = 100
	x[2] = 200
	fmt.Println(x[1])
	//delete build in function
	delete(x, 2)
	fmt.Println(x[2])

}
