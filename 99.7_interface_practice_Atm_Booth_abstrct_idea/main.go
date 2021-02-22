package main

import "fmt"

func main() {
	fmt.Println()
	var cd cardHolder
	cd.Name = "Md.Nihad Hossain"
	cd.AccountNO = 1025896666
	cd.Money = 50000
	cd.withDraw(500)
}

type atmWithDraw interface {
	withDraw()
}
type cardHolder struct {
	Name      string
	AccountNO int
	Money     int
}

func (c cardHolder) withDraw(l int) {
	fmt.Println(c.Name)
	fmt.Println(c.AccountNO)
	fmt.Printf("now you have %v \n witdraw %v \n", c.Money-l, l)

}
