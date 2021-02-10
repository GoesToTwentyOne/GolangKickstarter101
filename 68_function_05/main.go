package main

import "fmt"

func main() {
	x, y := foo()
	fmt.Println(x, y)

}
func foo() (int, int) {
	return 20, 30
}
