package main

import "fmt"

var x, y int
var (  // 这种因式分解关键字的写法一般用于声明全局变量
    a int
    b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

/* 声明全局变量 */
var g int = 20

func main() {
	/* 声明局部变量 */
	var g int = 10

	fmt.Printf ("结果： g = %d\n",  g)
}