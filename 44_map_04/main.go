package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hello"
	elements["G"] = "Good"
	elements["T"] = "Talent"
	elements["B"] = "Boy"
	fmt.Println(elements["B"])
	_, ok := elements["F"]
	fmt.Println(ok)
}
