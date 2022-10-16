package main

import "fmt"

func main() {
	score := -79

	grade := ""

	switch {
	case score < 50:
		grade = "F"
	case score < 60:
		grade = "D"
	case score < 70:
		grade = "C"
	case score < 80:
		grade = "B"
	case score >= 80:
		grade = "A"
	default:
		fmt.Println("Undefined Score")
	}

	fmt.Println(grade)

}
