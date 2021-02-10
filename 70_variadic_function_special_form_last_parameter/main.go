package main

import "fmt"

func main() {
	fmt.Println(addvariadic(1, 2, 3, 4, 5))

}
func addvariadic(args ...int) int {
	sum := 0
	for v := range args {
		sum += v
	}
	return sum
}
