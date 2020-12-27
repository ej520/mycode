package lession2

import "fmt"

func Panduan(c1 int) {
	a := c1
	if a > 3 {
		fmt.Println("a > 3")
	} else if a == 3 {
		fmt.Println("a == 3")
	} else {
		fmt.Println("a < 3")
	}
}
