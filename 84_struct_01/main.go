package main

import "fmt"

func main() {
	s1 := students{"Nihad", 300, 299, "CSE"}
	fmt.Println(s1)

}

type students struct {
	Name       string
	Roll       int
	Reg        int
	Department string
}
