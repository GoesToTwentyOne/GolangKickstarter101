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
	fmt.Println("enter your value :")
	MyName, _ := input.ReadString('\n')
	MynameRating, _ := strconv.ParseFloat(strings.TrimSpace(MyName), 64)
	fmt.Println(MynameRating + 2)
}
