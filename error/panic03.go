//Go error（错误） panic（异常）, defer, recover 的异常处理
package main

import (
	"errors"
	"fmt"
)

// from : http://blog.csdn.net/yyxyong/article/details/73734000
/*
1.go语言不支持传统的 try…catch…finally 这种异常，因为Go语言的设计者们认为，将异常与控制结构混在一起会很容易使得代码变得混乱。在Go语言中，使用多值返回来返回错误。不要用异常代替错误，更不要用来控制流程。在极个别的情况下，也就是说，遇到真正的异常的情况下（比如除数为0了）。才使用Go中引入的Exception处理：defer, panic, recover。
2.注意: 返回错误值与异常属于两种不同场景, 错误值属于普通正常的函数返回值,并不会导致程序中断执行; 而Go异常在不捕捉的情况下, 会导致程序中断执行;
3.代码演示:
普通函数,错误值返回
*/

/*
func test_common_ret() (int, error) {
	return 1, errors.New("hello error") //普通的 错误值返回, 不会导致程序中断执行, 返回后继续进行下面的打印语句
}

func main() {
	ret, err := test_common_ret()
	fmt.Println(ret, err) //打印语句
}
*/
//运行结果:
//1 hello error

//2 进行异常抛出, 但不捕捉, 导致程序中断执行
/*
func test_panic_ret() (int, error) {
	err := errors.New("hello error")
	ret := 34
	panic(err)      //此处抛出异常, 未进行异常捕捉,直接导致程序中断执行, 不会继续执行接下来的 return语句和打印语句
	return ret, err // 普通错误值返回,
}
func main() {
	ret, err := test_panic_ret()
	fmt.Println("start print")
	fmt.Println(ret, err) //打印语句
	fmt.Println("end print")
}
*/
/*
运行结果:
panic: hello error
goroutine 1 [running]:      //下面的输出为系统出错的堆栈
main.test_panic_ret(0xc042068060, 0xc042037f68, 0x0)
        F:/Go/src/add/addtesttime.go:268 +0x7f
main.main()
        F:/Go/src/add/addtesttime.go:272 +0x2d
*/

//3.抛出异常, 进行捕捉, 程序不会出现异常中断, 会继续执行后面的语句(不过要注意的是，异常捕捉之后，逻辑并不会恢复到panic那个点去，函数还是会在defer之后返回。)

func test_recov_panic_ret() (int, error) {
	defer func() { //进行异常捕捉
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	err := errors.New("hello error")
	ret := 34
	fmt.Println("start panic")
	panic(err)                //此处抛出异常, 进行异常捕捉, 不会导致程序中断执行
	fmt.Println("stop panic") //永远不会执行的代码
	return ret, err           //普通错误值返回
}
func main() {
	ret, err := test_recov_panic_ret()
	fmt.Println(ret, err) //打印语句
}
