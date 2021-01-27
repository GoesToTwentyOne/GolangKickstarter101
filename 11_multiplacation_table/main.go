package main

import "fmt"

func main() {
	fmt.Println("Enter your number :")
	var num int
	fmt.Scanln(&num)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v"+"X"+"%v"+"="+"%v", num, i, num*i)
		fmt.Println()

	}

}
