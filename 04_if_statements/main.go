package main

import "fmt"

func main() {
	Morning := 6
	if Morning < 5 || Morning >= 5 {
		if Morning >= 6 {
			fmt.Printf("This MOrning Good Morning")
		}
		fmt.Printf(" & Now it's not night")
	} else {
		fmt.Println("It was great night")

	}
}
