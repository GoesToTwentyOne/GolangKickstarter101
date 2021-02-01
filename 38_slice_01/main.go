package main

import "fmt"

func main() {
	x := []string{"Nihad", "Rowjatul", "Neha"}

	fmt.Println(x)
	x = append(x, "Nafis")
	fmt.Println(x)
	//deffrent from python
	fmt.Println(x[1 : len(x)-1])

}
