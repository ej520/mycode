// +build ignore

package main

import (
	"fmt"
)

//不定项输入参数
func mo(data1 int, data2 ...string) {
	fmt.Println(data1, data2)
	for k, v := range data2 {
		fmt.Println(k, v)
	}
}

//闭包函数： 函数返回一个函数，就交闭包函数
func bb() func(int) {
	//返回一个函数，是否有入参和出参，看定义，要一致
	return func(num int) {
		fmt.Println("闭包函数", num)
	}
}

func main() {
	mo(111, "a", "1", "b", "2")

	//传入不定项参数
	ar := []string{"a", "b", "c", "d"}
	mo(9527, ar...)

	(func() {
		fmt.Println("自己执行，不用调用")
	}())

	//闭包函数调用
	bb()(100)
}
