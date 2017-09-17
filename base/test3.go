package main

import "fmt"
import "math"

//bool: true,false 不能用数字
//int16/uint16  int16有符号的  uint16 无符号的
// int / uint int8 / uint8 。。。。
//byte (uint8别名)
//浮点型 float32 float64  小数点精确到7/15小数位
// 复数 complex64 complex128
//其它值类型 array struct string
//引用类型 slice map chan
//函数类型 func

//类型零值： 值类型：0 bool：false  string：空字符串
func zero() {
	var a int
	var b string
	var c bool
	var d [3]int
	var e [1]string
	var f [2]bool

	fmt.Println(a, b, c, d, e, f)
	fmt.Println(math.MaxInt32) //int32的最大值，通过这种方式来判断值是否溢出
	fmt.Println(math.MaxUint32)
}

func varvlue() {
	var b int = 1 //方式1
	fmt.Println(b)
	var c = 1 //系统推断
	d := 1    //变量申明与赋值最简模式 系统推断

	fmt.Println(b, c, d)

}

//var v1 = 2 //全局变量只能用var， 不能省略var关键字，所以不能用:= ,

func test1() {
	const a, b = 1, 2
	fmt.Println(a)
	fmt.Println(b)

}

func test2() {
	const (
		a, b = 1, 2
		c, d
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
func test3() {
	const (
		a = 'A'
		b
		c = iota
		d
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}

func main() {
	var a string
	fmt.Println(a)
	var b int
	fmt.Println(b)
	var c bool
	fmt.Println(c)
	//v1 = 10
	v1 := 10
	fmt.Println(v1)

	var z1 int64 = 6555
	var z2 = string(z1)
	fmt.Println(z1)
	fmt.Println(z2)

	test3()
	zero()
}
