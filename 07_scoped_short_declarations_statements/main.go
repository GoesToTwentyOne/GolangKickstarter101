package main

import "fmt"

func main() {
	switch name := "Rowjatul"; name {
	case "Rowjatul":
		fmt.Println("My name is Rowjatul")
	case "Annika":
		fmt.Println("My name is annika")
	case "Aspia":
		fmt.Println("My name is Aspia")
	case "Nihad":
		fmt.Println("my name is Nihad ")
	default:
		fmt.Println("Sorry your name isn't here")

	}
	fmt.Printf("\n \n \n \n")
	x := 60
	y := 70
	if product := x * y; product <= 300 {
		fmt.Println("This is 3 hundred")
	} else if product <= 400 {
		fmt.Println("This is 4 hundred")
	} else {
		fmt.Println("This is 420")
	}

}
