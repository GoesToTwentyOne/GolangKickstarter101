package main

import "fmt"

func main() {
	var name string //single variable declare
	name = "Nihad"  //value assign or initialization
	fmt.Println(name)
	var name1 string = "Nihad Hossain" // at on time single variable declare & value assign or initialization
	fmt.Println(name1)
	var num1, num2, num3 int = 4, 6, 8 // multiple variable declare and value assign
	fmt.Println(num1, num2, num3)
	result := 5.00 // shorthand variable declaration
	fmt.Println(result)

	//When we doesn't assign an value that time all value is 0
	var a int
	var Myname string
	var b bool
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", Myname)
	fmt.Printf("%v \n", b)

	//type check with shorhand variable declare
	fmt.Printf("\n \n \n")
	w := 2.0
	x := 10
	y := true
	z := "Zoo"
	fmt.Printf("%v %T \n", w, w)
	fmt.Printf("%v %T \n", x, x)
	fmt.Printf("%v %T \n", y, y)
	fmt.Printf("%v %T \n", z, z)

	fmt.Printf("\n \n \n")

	h := 10 //binary check
	fmt.Printf("%b \n", h)

	i := 1000 // decemal print
	fmt.Printf("%d \n", i)

	j := 40.3688 // after . two number print
	fmt.Printf("%.2f", j)
	fmt.Printf("\n \n \n")
	//constant
	const Name = "This is me and This is my sister "
	fmt.Printf("%v \n", Name)
	fmt.Printf("\n \n \n")

	// multiple variable declaration
	Num, Result, MyNameIs, Love := 20, 5.0, "Nihad Hossain", true
	fmt.Printf("\n %v \n %.2f \n %v \n %v \n", Num, Result, MyNameIs, Love)

}
