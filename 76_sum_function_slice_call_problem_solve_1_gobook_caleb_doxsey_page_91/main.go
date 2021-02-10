package main

import "fmt"

func main() {
	x := []int{20, 30, 40, 50, 60, 70}
	fmt.Println(sum(x))

}
func sum(x []int) int {
	total := 0
	for _, v := range x {
		total += v

	}
	return total
}
