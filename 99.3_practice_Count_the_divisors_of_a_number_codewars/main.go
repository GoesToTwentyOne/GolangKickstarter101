package main

import "fmt"

func main() {
	fmt.Println(divisors(4))

}
func divisors(n int) int {
	//Put your code here
	count := 0

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++

		}

	}
	return count

}
