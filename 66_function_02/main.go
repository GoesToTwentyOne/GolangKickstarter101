package main

import "fmt"

func main() {
	x := []int{20, 30, 40, 50, 60}
	fmt.Println(avarage(x))

}
func avarage(x []int) float64 {
	sum := 0
	for _, v := range x {
		sum += v

	}
	return float64(sum / len(x))

}
