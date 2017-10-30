package main


func main() {

	/*var x interface{}
	

	switch i := x.(type) {
	case nil:
		fmt.Print("X 的类型 %T",i)
	case int :
		fmt.Print("X 的类型  int")
	default:
		fmt.Print("X 的类型 未知",)
	}*/
	var a int = 100
	var b int = 20

	print(max(a,b))

}

func max (a,b int) int {
	if a > b {
		    return a
	}else {
		return b
	}
}