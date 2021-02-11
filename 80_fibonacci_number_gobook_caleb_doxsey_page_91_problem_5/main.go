package main

import "fmt"

func main() {
	fmt.Println(fib(10))

}
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	fmt.Println(fib(n-1) + fib(n-2))
	return fib(n-1) + fib(n-2)
}
