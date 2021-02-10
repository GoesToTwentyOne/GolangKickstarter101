package main

import "fmt"

func main() {
	//dont clear at now
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
