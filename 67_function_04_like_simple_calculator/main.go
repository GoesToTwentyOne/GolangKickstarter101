package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//input take
	input := bufio.NewReader(os.Stdin)
	//input 1
	fmt.Print("Enter 1st number : ")
	inputdata1, _ := input.ReadString('\n')
	inputdata1conv1, _ := strconv.ParseInt(strings.TrimSpace(inputdata1), 10, 64)
	//input 2
	fmt.Print("Enter 2nd number : ")
	inputdata2, _ := input.ReadString('\n')
	inputdata1conv2, _ := strconv.ParseInt(strings.TrimSpace(inputdata2), 10, 64)
	//fmt.Println(inputdata1conv1, inputdata1conv2)
	//input call function
	fmt.Print("Enter what operation you want 1 for add | 2 for subtraction | 3 for multiplacation | 4 for division : ")
	funccall, _ := input.ReadString('\n')
	funccallconv, _ := strconv.ParseInt(strings.TrimSpace(funccall), 10, 64)
	//function calling
	if funccallconv == 1 {
		fmt.Println("Sum result :", add(inputdata1conv1, inputdata1conv2))
	} else if funccallconv == 2 {
		fmt.Println("subtraction result :", subtraction(inputdata1conv1, inputdata1conv2))
	} else if funccallconv == 3 {
		fmt.Println("multiplication result :", multiplacation(inputdata1conv1, inputdata1conv2))
	} else if funccallconv == 4 {
		fmt.Println("devide result :", devide(inputdata1conv1, inputdata1conv2))
	}

}

//make some function
func add(x, y int64) int64 {
	z := x + y
	return z

}
func subtraction(x, y int64) int64 {
	z := x - y
	return z
}
func multiplacation(x, y int64) int64 {
	z := x * y
	return z
}
func devide(x, y int64) int64 {
	z := x / y
	return z
	//return z

}
