package main

import "fmt"

func main() {
	fmt.Println()
	a := animal{
		AnimalType: "domestic ",
		Name:       "My animal",

		NumberOfAnimal: 1,
		age:            1,
	}
	//dog
	d := dog{
		animal: animal{
			AnimalType:     "Domestic",
			Name:           "Tommy",
			NumberOfAnimal: 4,
			age:            12,
		},
		leg: 4,
	}
	h := hen{
		animal: animal{
			AnimalType:     "Domestic",
			Name:           "KUntu",
			NumberOfAnimal: 3,
			age:            2,
		},
		leg: 2,
	}
	//print animal
	fmt.Printf("%+v  \n", a)
	//print dog
	fmt.Printf("%+v  \n", d)
	//print hen
	fmt.Printf("%+v  \n", h)
	//method print
	a.allanimal()
	d.allanimal()
	h.allanimal()
	////polymorphism
	zoo(a)
	zoo(d)
	zoo(h)

}

type animal struct {
	AnimalType     string
	Name           string
	NumberOfAnimal int
	age            int
}
type dog struct {
	animal
	leg int
}
type hen struct {
	animal
	leg int
}

func (a animal) allanimal() {
	fmt.Printf("animal name : %v  animal age %v \n", a.Name, a.age)

}
func (d dog) allanimal() {
	fmt.Printf("Animal type %v animal leg %v \n", d.AnimalType, d.leg)

}
func (h hen) allanimal() {
	fmt.Printf("Animal type %v animal leg %v \n", h.AnimalType, h.leg)

}

type domestic interface {
	allanimal()
}

func zoo(do domestic) {
	fmt.Printf("This is Zoo : %+v  \n", do)
}
