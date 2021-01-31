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
	for i := 0; i < len(x); i++ {
		fmt.Printf("Enter your value array[%v]:", i)
		arrayInput, _ := input.ReadString('\n')
		arrayInputconv, _ := strconv.ParseInt(strings.TrimSpace(arrayInput), 10, 64)
		x[i] = arrayInputconv
	}
	fmt.Println(x)

}
