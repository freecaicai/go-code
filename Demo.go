package main

import "fmt"

func main() {
	var age = 123
	var test = "chenghaojie"
	var url = "code=%d&endDate=%s"
	var tarage_url = fmt.Sprintf(url, age, test)
	var b bool = true
	var c, d = 1, "字符串"
	//const name = "qq1"
	//fmt包含了IO/输入/输出的函数
	a := 1
	var (
		vname1 = 1
		vname2 = 2
		vname3 = 3
	)

	fmt.Println("Hello  world" + "  who is she")
	fmt.Println("菜鸟教程")
	fmt.Println(tarage_url)
	fmt.Println(b)
	fmt.Println(c, d)
	fmt.Println(a)
	fmt.Println(vname1, vname2, vname3)
	demo1()
}

func demo1() {
	fmt.Println("你试试啊！")
}
