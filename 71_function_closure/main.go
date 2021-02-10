package main

import "fmt"

func main() {
	//closure syntax
	multiplication := func(x :=make([]int)) int {
		sum := 0
		for _, v := range x {
			sum += v

		}
		return sum

	}
	y := []int{20, 30, 40, 50, 60}
	fmt.Println(multiplication(y))
}
