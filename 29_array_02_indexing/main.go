package main

import "fmt"

func main() {
	x := [6]string{"mango", "apple", "orange", "jack fruit", "banana", "grape"}
	fmt.Println(x[2:]) // last item print without indexing
	fmt.Println(x[:3]) // print 0 indexing yto 3
	fmt.Println(x[0])
}
