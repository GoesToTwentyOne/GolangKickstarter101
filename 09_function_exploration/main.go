package main

import "fmt"

func hello() {
	fmt.Println("Hello this is function started")
}

func exponenttwo(num int) int {
	return num * 2

}
func gpa(midtermGrade float32, finalGrade float32) (string, float32) {
	averageGrade := (midtermGrade + finalGrade) / 2
	var gradeLetter string
	if averageGrade > 90 {
		gradeLetter = "A"
	} else if averageGrade > 80 {
		gradeLetter = "B"
	} else if averageGrade > 70 {
		gradeLetter = "C"
	} else if averageGrade > 60 {
		gradeLetter = "D"
	} else {
		gradeLetter = "F"
	}

	return gradeLetter, averageGrade
}

func main() {
	hello()
	fmt.Println(exponenttwo(4))
	fmt.Println(gpa(84.4, 45.5))

}
