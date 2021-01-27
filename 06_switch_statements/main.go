package main

import "fmt"

func main() {
	fmt.Println()
	ClothChoice := "sweater"
	switch ClothChoice {
	case "shirt":
		fmt.Println("This is a shirt")
	case "pant":
		fmt.Println("This is a pant")
	case "jackets":
		fmt.Printf("This is jacket")
	case "sweater":
		fmt.Println("This is sweater")
	default:
		fmt.Println("sorry out of stock")
	}

}
