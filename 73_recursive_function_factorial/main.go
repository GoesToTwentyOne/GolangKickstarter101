package main

import "fmt"

func main() {
	//call function and pass value
	fmt.Println(factorial(5))

}
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	//when n=0 it always enden
	return n * factorial(n-1)

}
