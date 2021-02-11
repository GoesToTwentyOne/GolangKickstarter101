package main

import "fmt"

func main() {
	x := 5
	fmt.Println(&x)
	pointer(&x)
	fmt.Println(&x)

}
func pointer(i *int) {
	*i = 2
}
