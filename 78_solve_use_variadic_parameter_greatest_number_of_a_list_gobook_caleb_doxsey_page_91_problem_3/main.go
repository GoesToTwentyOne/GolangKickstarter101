package main

import "fmt"

func main() {

	fmt.Println(greatestNumber(5, 1, 3, 2, 4, 6, 7))

}
func greatestNumber(list ...int) int {
	var max int = 0
	var i int
	var j int
	for i = 0; i < len(list); i++ {
		for j = i + 1; j < len(list); j++ {
			if list[i] < list[j] {
				max = list[i]

			}

		}

	}
	//fmt.Println(max)
	//fmt.Println(list[4])
	return max

}
