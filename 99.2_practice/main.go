package main

import "fmt"

func main() {
	x := []int{5, 3, 4, 5}

	fmt.Println(grow(x))

}
func grow(arr []int) int {
	var mul int = 1
	for i := 0; i < len(arr); i++ {
		mul = mul * arr[i]

	}
	return mul
}
