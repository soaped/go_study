package main

import "fmt"

func main() {
	var b  = 2
	var a  *int
	a = &b
	fmt.Print(a)
}
