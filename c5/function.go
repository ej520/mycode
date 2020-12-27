// +build ignore

package main

import (
	"fmt"
)

func a(data1 int, data2 string) {
	if data1 > 1 {
		fmt.Println("func a ", data1)

	} else {
		fmt.Println("func a ", data2)
	}
}

func b(data1 int, data2 string) (ret1 int, ret2 string) {
	ret1 = data1 + 1
	ret2 = data2 + "FUN ADD String"

	return //默认返回已经规定好的ret1和ret2
}

func main() {
	//入参函数
	a(0, "现在是0小于1")

	//入参+出参 函数
	r1, r2 := b(0, "mystr1")
	fmt.Println("func b ", r1, r2)

	//匿名函数1  =》 没有名字的函数
	c := func(data1 string) {
		fmt.Println("func 匿名C ", data1)
	}
	c("我是匿名函数")

	d := func() {
		fmt.Println("func 匿名d ")
	}
	d()

	//匿名函数output
	num := func() int {
		fmt.Println("func 匿名x ")
		return 0
	}()
	fmt.Println(num)

}
