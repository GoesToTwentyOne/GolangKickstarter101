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
	fmt.Println("Enter your several number you want add :")
	check, _ := input.ReadString('\n')
	checkconv, _ := strconv.ParseInt(strings.TrimSpace(check), 10, 64)

	var sum int64
	var i, j int64
	for j = 1; j <= checkconv; j++ {
		fmt.Printf("Enter your  add number %v:", j)
		addnum, _ := input.ReadString('\n')
		addnumconv, _ := strconv.ParseInt(strings.TrimSpace(addnum), 10, 64)
		for i = j; i <= j; i++ {
			sum = sum + addnumconv
			//fmt.Println(sum)

		}
		//fmt.Println(sum)

	}
	fmt.Printf("add %v value and result = %v \n", checkconv, sum)

}
