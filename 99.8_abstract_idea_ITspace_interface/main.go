package main

import "fmt"

func main() {
	var it materials
	it.NameOfOwner = "Kanuuu Vai"
	it.IntelProcessor = 280
	it.Monitor = 369
	it.GraphicsCard = 500
	it.SystemUnit = 8000
	it.Itspace()

}

type itShop interface {
	Itspace()
}
type materials struct {
	NameOfOwner    string
	Monitor        int
	GraphicsCard   int
	SystemUnit     int
	IntelProcessor int
}

func (m materials) Itspace() {
	fmt.Println("hi welcome everyone ", m.NameOfOwner)
	fmt.Println("I have many computer goods ")
	fmt.Println(m.IntelProcessor)
	fmt.Println(m.GraphicsCard)
	fmt.Println(m.Monitor)
	fmt.Println(m.SystemUnit)

}
