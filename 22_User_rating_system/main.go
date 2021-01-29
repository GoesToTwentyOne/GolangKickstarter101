package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	//food input
	fmt.Println("Enter your food name : ")
	youeat, _ := input.ReadString('\n')
	//name input
	fmt.Println("Enter your name : ")
	name, _ := input.ReadString('\n')
	// fmt.Println(youeat, name)
	// input rating
	fmt.Println("Enter your rating : ")
	rating, _ := input.ReadString('\n')
	ratingConv, _ := strconv.ParseInt(strings.TrimSpace(rating), 10, 64)
	//backend development
	fmt.Println()
	fmt.Printf("Hello,\n%v"+"Thanks for eat %v"+"your rating is recorded my system\n"+"your rating is %v "+
		"  Time : %v\n", name, youeat, ratingConv, time.Now())

	if ratingConv == 5 {
		fmt.Println("your service is great")
	} else if ratingConv == 4 {
		fmt.Println("your service is good")
	} else if ratingConv < 4 {
		fmt.Println("be serious and this is worning")
	}

}
