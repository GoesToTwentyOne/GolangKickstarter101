package main

import "fmt"

func main() {
	//A defer statement defers the execution of a function until the surrounding function returns.
	defer motherSon()
	fatherSon()
	sisterBrother()

}
func motherSon() {
	fmt.Println("My son  my peace")

}
func fatherSon() {
	fmt.Println("Little a boy ")

}
func sisterBrother() {
	fmt.Println("He didn't give me choclate")

}
