package lession3

import (
	"fmt"
)

func FuncArray1() {
	a := [3]int{0, 1, 2}

	b := [...]int{1, 2, 3, 4, 5, 6, 7, 222, 223}

	var c = new([10]int)
	c[5] = 3
	fmt.Println(a, b, c)

	zoom := [...]string{"狗子", "猫咪", "猴子"}

	for i := 0; i < len(zoom); i++ {
		fmt.Println(zoom[i] + "跑。。。")
	}

	for k, v := range zoom {
		fmt.Println(k, v)
	}
}

func FuncArray2() {

	// len cap
	var a = [10]int{0, 1, 2}
	a[3] = 3
	fmt.Println(a)

}
