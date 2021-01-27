package main

import "fmt"

func main() {
	π := 3.14159
	var R float64
	fmt.Scan(&R)

	A := π * R * 2
	fmt.Printf("A=%.4f", A)
	fmt.Println()

}
