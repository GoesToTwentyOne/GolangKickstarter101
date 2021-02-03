package main

import "fmt"

func main() {
	bit := []int{2, 4, 8, 16, 32, 64, 128, 256}
	for v := range bit {
		if bit[v] == 16 {
			//break
			//continue
			goto me
		}
		fmt.Println(bit[v])

	}
me:
	fmt.Println("Joy Bangla")

}
