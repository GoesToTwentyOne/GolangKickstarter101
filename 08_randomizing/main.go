package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(10000)
	if x <= 5000 {
		fmt.Println("You have to much save for future")
	} else if x <= 8000 {
		fmt.Println("you have to pay your bank interest card")
	} else {
		fmt.Println("all time ok Now I am ok")
	}

}
