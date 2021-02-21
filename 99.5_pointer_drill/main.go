package main

import "fmt"

func main() {
	//value assign
	a := 10
	var b *int //pointer declare
	//print normal variable =>accessing
	fmt.Println(a)
	//print normal variable memory address
	fmt.Println("memory address a : ", &a)
	//print pointer variable default it's hold nil
	fmt.Println(b)
	//memory address b
	fmt.Println("memory address b :", &b)

	//referencin address of a
	b = &a
	//print b reference address
	fmt.Println("after referencing :", b)
	//b dereferencing and print value
	fmt.Println("after dereferencing and print value :", *b)
	//function value
	pointerdril(&a)

}
func pointerdril(c *int) {
	fmt.Println(c)
	//before reassign and dereferencing
	fmt.Println("before reassign and dereferencing : ", *c)
	*c = *c + 20
	fmt.Println(c)
	fmt.Println("after reassign and dereferencing", *c)
}
