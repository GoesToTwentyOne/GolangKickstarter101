package main

import "fmt"

func main() {
	// only run main function

	book := struct {
		Name      string
		TotalPage int
		Author    string
		Review    float64
	}{
		Name:      "As a Man Thinketh",
		TotalPage: 22,
		Author:    "James Allen",
		Review:    5.00,
	}
	fmt.Printf("%+v", book)

}
