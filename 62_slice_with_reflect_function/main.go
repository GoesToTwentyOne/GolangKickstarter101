package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := make([]string, 2, 5)
	x[0] = "Nihad"
	x[1] = "Hossain"
	fmt.Println(x)
	fmt.Println(reflect.TypeOf(x).Kind().String())
}
