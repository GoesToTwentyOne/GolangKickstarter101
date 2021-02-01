package main

import "fmt"

func main() {
	slice1 := []int64{4, 6, 8, 10, 20}
	slice2 := make([]int64, 3)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}
