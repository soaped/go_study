// test project main.go
package main

import (
	"fmt"
)

func main() {
	var grade string = "B"
	var marks = 91

	switch marks {
	case 90:
		grade = "A"
	default:
		grade = "D"
	}

	switch  {
	case grade == "A": print("成绩不错")
	case grade == "D": print("没及格")
	}
}
