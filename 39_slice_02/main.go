package main

import "fmt"

func main() {
	numberSlice := make([]int, 2, 3)
	numberSlice[0] = 10
	numberSlice[1] = 50
	fmt.Println(numberSlice)
	//dosen't give error as slice are special rule
	numberSlice = append(numberSlice, 45, 55)
	fmt.Println(numberSlice)
	fmt.Println(cap(numberSlice))

}
