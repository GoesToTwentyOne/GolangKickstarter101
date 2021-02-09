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
	fmt.Printf("Enter your  value How many value add and want to avg?example:10,20 etc :\n")
	arraylimitinput, _ := input.ReadString('\n')
	arraylimitinputconv, _ := strconv.ParseInt(strings.TrimSpace(arraylimitinput), 10, 64)
	x := make([]int64, arraylimitinputconv)
	//try to constant but I am stuck now
	var i int
	var sum int64 = 0
	for i = 0; i < len(x); i++ {
		fmt.Printf("Enter your value array[%v] :", i)
		arrayinput, _ := input.ReadString('\n')
		arrayinputconv, _ := strconv.ParseInt(strings.TrimSpace(arrayinput), 10, 64)
		x[i] = arrayinputconv
		sum = sum + x[i]

	}
	fmt.Printf("sum : %v \n", sum)

	// specific look type conversion
	fmt.Printf("avarage : %v", sum/int64(len(x)))

}
