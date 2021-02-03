package main

import "fmt"

func main() {
	//new -> only allocates - no initial limits of memory
	//make-> allocate and intial memory - non zero storage
	hitt := make(map[int]string)
	hitt[1950] = "Sunset Boulevard"
	hitt[2010] = "Inception"
	hitt[2008] = "The Dark Knight"
	hitt[2019] = "1917"
	hitt[1980] = "The Shining"
	delete(hitt, 1950)
	fmt.Println(hitt)
	for k, v := range hitt {
		fmt.Printf("This is hit year %v and name of %v \n\n", k, v)
	}
}
