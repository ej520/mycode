//## +build ignore

package main

import "fmt"

//值传参： 不需要改变外面的值
func ppValue(v string) {
	v = "我不会改变外面"
}

//指针传参：需要改变外面的值的情况
func ppPoint(p *string) {
	*p = "我变了，会改变外面"
}

func main() {
	a := "值：我是我"
	b := "指针：我是我"
	//值传参： 不需要改变外面的值
	ppValue(a)
	fmt.Println(a)

	//指针传参：需要改变外面的值的情况
	ppPoint(&b)
	fmt.Println(b)
}

//结构体的指针
//map的指针
//回参的指针
