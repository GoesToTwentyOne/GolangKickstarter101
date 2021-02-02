package main

import "fmt"

func main() {
	elements := map[string]map[string]string{
		"N": map[string]string{
			"name":   "Nihad",
			"status": "student",
		},
		"R": map[string]string{
			"name":   "Rowjatul",
			"status": "student",
		},
		"A": map[string]string{
			"name":   "Annika",
			"status": "student",
		},
	}
	fmt.Println(elements["A"])
}
