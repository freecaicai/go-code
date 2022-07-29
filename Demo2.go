package main

import "fmt"

/*

 */
func main() {
	const (
		a = iota
		b
		c
		d = "11"
		e = 100
		f
		g = iota
		h
	)
	i := 19
	fmt.Println(a, b, c, d, e, f, g, h)
	if i < 20 {
		fmt.Println("i<20")
		if b == 1 {
			fmt.Println("b：", b)
		}
	} else {
		fmt.Println("i的值为", i, "11weq")
	}
}
