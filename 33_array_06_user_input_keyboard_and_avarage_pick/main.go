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
	// fmt.Printf("Enter your limit value to get avarage :")
	// arraylimit, _ := input.ReadString('\n')
	// arraylimitconv, _ := strconv.ParseInt(strings.TrimSpace(arraylimit), 10, 64)

	//take type as error remove
	var x [5]int64
	var i int
	var sum int64 = 0
	for i = 0; i < len(x); i++ {
		//take input
		fmt.Printf("Enter your array[%v] :", i)
		arrayinput, _ := input.ReadString('\n')
		arrayinputconv, _ := strconv.ParseInt(strings.TrimSpace(arrayinput), 10, 64)
		x[i] = arrayinputconv
		sum = sum + x[i]

	}
	fmt.Printf("avarage : %v", sum/5)

}
