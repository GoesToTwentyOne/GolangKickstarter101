package main

import "fmt"

func main() {
	s := studentResult{"Md.Nihad Hossain", 85, 90, 99, 99, 87}
	//fmt.Println(s)
	fmt.Println(calculatorResult(s))

}

type studentResult struct {
	Name          string
	BanglaNo      int
	EnglishNo     int
	MathNo        int
	ScienceNo     int
	AgrecaltureNo int
}

func calculatorResult(s studentResult) (string, float64) {

	return s.Name, (float64(s.BanglaNo+s.EnglishNo+s.MathNo+s.ScienceNo+s.AgrecaltureNo) / 5)

}
