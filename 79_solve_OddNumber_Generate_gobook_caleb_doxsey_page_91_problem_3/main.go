package main

import "fmt"

func main() {
	fmt.Println(makeOddGenerator(1))

}
func makeOddGenerator(i int) int {
	if i%2 == 0 {
		i = i + 1

	} else {
		i = i + 2

	}
	return i

}
