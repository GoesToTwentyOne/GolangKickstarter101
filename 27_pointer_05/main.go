package main

import "fmt"

func name(j *string) {
	*j = "Hossain"

}
func main() {
	j := "Nihad"
	fmt.Println(j)
	name(&j)
	fmt.Println(j)

}
