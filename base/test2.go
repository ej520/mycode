// 当前程序的包名
package main

//单行注释
/* 多行注释 */

// 导入其它的包
//import std "fmt"
//声明引用包组，常量组，全局变量组，一般类型声明（非接口、非结构） 都可以用这种方法。 注意：函数体内 不能使用这种方法申明变量
// import() const() var() type()(非接口、非结构)
import (
	std "fmt"
	"time"
)

// 常量的定义
const PI = 3.14

// 全局变量的声明与赋值,整个包都可以使用
var name = "gopher"

//首字母大写，可以跨包访问；首字母小写，只能包内调用

// 一般类型声明
type newType int

// 结构的声明
type gopher struct{}

// 接口的声明
type golang interface{}

// 由 main 函数作为程序入口点启动
func main() {
	std.Println("Hello world!你好，世界！")
	std.Println(time.Now(), "Hello world!	你好，世界！")
}
