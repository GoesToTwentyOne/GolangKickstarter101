package main

import "fmt"

func main() {
	x := [5]int{20, 40, 60, 80, 100}
	sum := 0
	for i := 0; i < len(x); i++ {
		sum += x[i]

	}
	fmt.Println("Result add 5 elements : ", sum)
}
