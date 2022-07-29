package main

import "fmt"

/*select 语句*/
func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for sum <= 10 {
		sum += sum
	}
	for {
		sum++
	}
	fmt.Println(sum)
}
