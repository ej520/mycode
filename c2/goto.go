package lession2

import "fmt"

func FuncGoTo1() {
	a := 0

	//LB1:
	//	fmt.Println("FuncGoTo1: LB1 block")
	//
LB2:
	for a < 10 {
		a++
		fmt.Println("FuncGoTo1: LB2 block", a)
		if a == 3 {
			break LB2
			goto LB3
		}
	}

LB3:
	fmt.Println("FuncGoTo1: LB3 block")
	//goto LB1

}

//goto逆向跳，容易出现死循环
