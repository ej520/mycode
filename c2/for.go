package lession2

import "fmt"

// break 终止循环
func FunFor1() {
	a := 0
	for {
		a++
		fmt.Println("FunFor1:", a)
		if a > 10 {
			break
		}
	}

}
func FunFor2() {
	for a := 0; a < 10; a++ {
		fmt.Println("FunFor2:", a)
	}
}

func FunFor3() {
	a := 0
	for a < 10 {
		a++
		fmt.Println("FunFor3:", a)
	}
}

//continue 跳过本次循环，即a=5跳过，继续6-10
func FunFor4() {
	a := 0
	for a < 10 {
		a++
		if a == 5 {
			fmt.Println("FunFor4: ...")
			continue
		}
		fmt.Println("FunFor4:", a)
	}
}
