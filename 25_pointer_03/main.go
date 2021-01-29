package main

import "fmt"

func foo(y *int) {
	fmt.Println("before ", y)
	fmt.Println("before ", *y)
	*y = 40
	fmt.Println("after ", y)
	fmt.Println("after ", *y)
}

func main() {
	x := 2
	fmt.Println("before main ", &x)
	fmt.Println("before main ", x)
	foo(&x)
	fmt.Println("after main ", &x)

	fmt.Println("after main ", x)

}
