package main

import "fmt"

func main() {
	aim := 20
	aimref := &aim
	//multiply aim
	aim = aim * 4
	fmt.Println(aim)
	//pointer print
	fmt.Println(*aimref)

}
