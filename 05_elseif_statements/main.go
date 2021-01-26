package main

import "fmt"

func main() {
	fmt.Println("Enter your subject mark  to check grade : ")
	var result int
	fmt.Scanln(&result)
	if result >= 80 && result <= 100 {
		fmt.Println("Your grade is A+")
	} else if result <= 79 && result >= 70 {
		fmt.Println("Your grade is A")
	} else if result <= 69 && result >= 60 {
		fmt.Println("Your grade is A-")
	} else if result <= 59 && result >= 50 {
		fmt.Println("Your grade is B")
	} else if result <= 49 && result >= 40 {
		fmt.Println("Your grade is C")
	} else if result == 33 && result >= 39 {
		fmt.Println("Your grade is D")
	} else if result <= 32 {
		fmt.Println("Your grade is F")
	} else {
		fmt.Println("Your data isn't correct")
	}

}
