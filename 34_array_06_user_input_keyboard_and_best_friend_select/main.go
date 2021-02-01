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
	var x [5]int64
	var i int
	var sum int64 = 0
	for i = 0; i < len(x); i++ {
		fmt.Printf("Enter your array[%v] :", i)
		arrayinput, _ := input.ReadString('\n')
		arrayinputconv, _ := strconv.ParseInt(strings.TrimSpace(arrayinput), 10, 64)
		x[i] = arrayinputconv
		sum = sum + x[i]

	}
	// specific look type conversion
	fmt.Printf("avarage : %v", sum/int64(len(x)))

}
