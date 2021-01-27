package main

import "fmt"

func main() {
	const pie = 3.14159
	var R float64
	fmt.Scan(&R)

	A := pie * R * 2
	fmt.Printf("A=%.4f \n", A)

}
