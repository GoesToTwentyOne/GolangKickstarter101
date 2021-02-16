package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//book example
	for i := 0; i < 10; i++ {
		// go fmt.Println("This is i", i)
		go f(i)
	}
	var input string
	fmt.Scanln(&input)

}
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}
