package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	MyName, _ := rd.ReadString('\n')
	fmt.Println(MyName)

}
