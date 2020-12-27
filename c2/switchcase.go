package lession2

import "fmt"

func FuncS1(num int) {
	a := num
	switch a {
	case 0:
		fmt.Println("num is 0")
	case 1:
		fmt.Println("num is 1")
	case 2:
		fmt.Println("num is 2")
	default:
		fmt.Println("00000000")
	}
}
