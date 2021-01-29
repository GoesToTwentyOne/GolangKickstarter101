package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your Started multiplication Table value you want started :")
	multiplacationstart, _ := input.ReadString('\n')
	multiplacationstartfinal, _ := strconv.ParseInt(strings.TrimSpace(multiplacationstart), 10, 64)
	fmt.Println("Enter your Started multiplication Table value you want ended :")

	multiplacationend, _ := input.ReadString('\n')
	multiplacationendfinal, _ := strconv.ParseInt(strings.TrimSpace(multiplacationend), 10, 64)
	var i int64
	var k int64
	for k = multiplacationstartfinal; k <= multiplacationendfinal; k++ {
		var num int64 = k
		fmt.Printf("Multiplacation Table  of = %v \n \n", num)
		for i = 1; i <= 10; i++ {

			fmt.Printf("%v X %v = %v \n", num, i, num*i)

		}
		fmt.Println()

	}

}
