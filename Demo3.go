package main

import "fmt"

/*switch语句*/

func main() {
	var temp1 int = 11
	str := ""
	switch temp1 {
	case 11:
		str = "A"
	case 22:
		str = "B"
	case 33:
		str = "C"
	default:
		str = "D"
	}
	fmt.Println("str:", str)
	fmt.Println("type：", &temp1)

}
