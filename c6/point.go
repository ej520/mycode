// +build ignore

package main

import "fmt"

func pp1() {
	var a int
	a = 123
	fmt.Println(a)

	//指针不能指向一个具体的值， 需要指向和他的类型一致的内存地址
	var b *int //*号+数据类型 申明指针
	b = &a     //指针是指向一个变量的原始内存地址，&号+变量表示取地址
	*b = 999   //*+内存地址 表明正在通过原始内存地址修改这个地址上的值
	fmt.Println(a, b)

	fmt.Println(a, b, *b)
	fmt.Println(a == *b, &a == b)
}

//数组指针
func arrpp() {
	var arr [5]string
	arr = [5]string{"1", "2", "3", "4", "5"}
	//数组指针的定义
	var arrP *[5]string
	arrP = &arr
	fmt.Println(arr, arrP)
}

//指针数组
func pparr() {
	//var arr [5]string
	//arr = [5]string{"1", "2", "3", "4", "5"}
	//定义指针数组
	var arrP [5]*string
	var str1 string = "str1"
	var str2 string = "str2"
	var str3 string = "str3"
	var str4 string = "str4"
	var str5 string = "str5"
	arrP = [5]*string{&str1, &str2, &str3, &str4, &str5}
	*arrP[3] = "NewString"
	fmt.Println(arrP[0], &str1)
	fmt.Println(*arrP[1], str2)
	fmt.Println(*arrP[2], str3)
	fmt.Println(*arrP[3], str4)
	fmt.Println(*arrP[4], str5)
}
func main() {
	pp1()
	//数组指针
	arrpp()
	//指针的数组
	pparr()
}
