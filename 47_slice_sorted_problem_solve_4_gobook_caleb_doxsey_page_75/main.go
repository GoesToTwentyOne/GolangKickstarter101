package main

import (
	"fmt"
)

func main() {
	var x = []int{
		48, 96, 86, 68, 57, 82, 63, 70, 17, 34, 83, 27, 19, 97, 9, 17,
	}
	var min int

	var i int
	var j int

	for i = 0; i < len(x); i++ {
		for j = i + 1; j < len(x); j++ {
			if x[i] < x[j] {
				min = x[i]

			}

		}
		//fmt.Println(min)

	}
	fmt.Println(min)

	fmt.Println()

}
