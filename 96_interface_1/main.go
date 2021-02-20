package main

import (
	"fmt"
)

func main() {
	a1 := agent{
		person: person{
			fn:  "Nihad",
			ln:  "Hossain",
			age: 20,
		},
		tp: "IDP Asian",
	}

	a2 := agent{
		person: person{
			fn:  "Rakib",
			ln:  "Khan",
			age: 30,
		},
		tp: "FBBCI Asian",
	}
	p1 := person{
		fn:  "Rowjatul",
		ln:  "Neha",
		age: 7,
	}

	fmt.Printf("%+v \n", a1)
	fmt.Printf("%+v \n", a2)
	//method call
	a2.pull()
	fmt.Println()
	////mathod call
	a1.pull()
	fmt.Println()

	//another person
	fmt.Println(p1)
	//polymorphisim->interface
	bar(a1)
	bar(a2)
	bar(p1)

}

type person struct {
	fn  string
	ln  string
	age int
}
type agent struct {
	person
	tp string
}

func (a agent) pull() {
	fmt.Printf("%v %v", a.fn, a.ln)
}
func (p person) pull() {
	fmt.Printf("%v %v", p.fn, p.ln)
}

type human interface {
	pull()
}

func bar(h human) {
	fmt.Println("I called human  ", h)
}
