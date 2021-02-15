package main

import "fmt"

func main() {
	// p := person{"Nihad Hossain"}
	// fmt.Println(p)
	// fmt.Println()
	a := new(android)
	a.talk()

}

type person struct {
	Name string
}
type android struct {
	person
	model string
}

func (p *person) talk() {
	fmt.Println(p.Name)
}
