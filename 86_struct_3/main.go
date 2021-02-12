package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println()
	v := new(vote)
	v.Name = "Nihad"
	v.Wardno = 5
	v.Center = "Thakurgaon Mohila Collage"
	v.Date = time.Now().String()
	fmt.Printf("%+v", v)

}

type vote struct {
	Name   string
	Wardno int
	Center string
	Date   string
}
