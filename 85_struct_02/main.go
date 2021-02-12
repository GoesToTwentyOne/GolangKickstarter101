package main

import "fmt"

func main() {
	u := user{"Md.Nihad Hossain", 20, "English"}
	fmt.Printf("%+v \n", u)
	//use new keyward
	u2 := new(user)
	u2.Name = "Sandy"
	u2.Age = 25
	u2.Language = "English"
	fmt.Printf("%+v \n", u2)
	//use ampersand sign
	u3 := &user{"lucky", 20, "English"}
	fmt.Printf("%+v", u3)

}

type user struct {
	Name     string
	Age      int
	Language string
}
