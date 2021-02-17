package main

import "fmt"

func main() {
	p1 := publisher{
		Name:        "English Applied",
		Place:       "India,Mumbai",
		Established: 1990,
	}
	fmt.Printf("%+v publisher \n", p1)
	b1 := book{
		publisher: publisher{
			Name:        "Invincible",
			Place:       "India,Mumbai",
			Established: 1990,
		},
		BookName: "Be a network marketing millionaire",
	}
	fmt.Printf("%+v Book \n", b1)

}

type publisher struct {
	Name        string
	Place       string
	Established int
}

//embeded type means OOP inharitance

type book struct {
	publisher
	BookName string
}
