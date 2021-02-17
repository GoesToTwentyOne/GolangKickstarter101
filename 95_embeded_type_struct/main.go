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
		Place:    "Bangladesh",
		BookName: "Be a network marketing millionaire",
	}
	fmt.Printf("%+v Book \n", b1)
	//specific call
	fmt.Println(b1.publisher.Place, b1.BookName)
	fmt.Println(p1.Name, p1.Place)

}

type publisher struct {
	Name        string
	Place       string
	Established int
}

//embeded type means OOP inharitance

type book struct {
	publisher
	Place    string
	BookName string
}
