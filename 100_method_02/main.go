package main

import "fmt"

func main() {
	a1 := agent{
		person: person{
			fn:  "Nihad",
			ln:  "Hossain",
			age: 20,
		},
		tp: "IDP Asian",
	}
	fmt.Printf("%+v ", a1)
	////mathod call
	a1.pull()
	a2 := agent{
		person: person{
			fn:  "Rakib",
			ln:  "Khan",
			age: 30,
		},
		tp: "FBBCI Asian",
	}
	fmt.Printf("%+v ", a2)
	//method call
	a2.pull()

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
