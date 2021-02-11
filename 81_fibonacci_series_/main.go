package main

import "fmt"

func main() {
	var first int = 0
	var second int = 1
	var fibo int
	fmt.Printf("%v  %v", first, second)
	for i := 3; i <= 5; i++ {
		fibo = first + second
		fmt.Print(" ", fibo)
		first = second
		second = fibo

	}
	//fmt.Println(fibo)

}
