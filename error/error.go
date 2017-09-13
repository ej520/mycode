package main

import "errors"
import "fmt"

/*
from: https://www.golangtc.com/t/589be068b09ecc2e18000207
错误和异常处理:
Go 和其他编程语言不太一样的是，它把错误和异常分开。
Go 提供了 error 类型用来表达错误，语法上允许函数有多个返回值，
因此，返回 error 的设计在 runtime 中随处可见。
另外，由于 Go 的普及率不高，所以我列一下代码（如图 1，2 所示），与大家分享一下 Go 的错误写法。
*/

/*
这个函数是两个偶数相加，
如果参数不是偶数的话，就会返回一个 error，
error 这个包是 Go 内置的，可以让你快速返回一个错误，不需要重新手动构造错误类型。
*/
func EvenNumAdd(a, b int) (int, error) {
	if a%2 != 0 || b%2 != 0 {
		return 0, errors.New(
			"parameter not even number",
		)
	}
	return a + b, nil
}

/*
调用它的时候可以通过第二个参数是否 nil 来判断它是否出错，然后把对应的错误信息打印出来。
*/
func main() {
	a, b := 1, 2
	c, err := EvenNumAdd(a, b)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%d + %d = %d", a, b, c)
	}
}
