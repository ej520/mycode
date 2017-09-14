//from :http://www.jb51.net/article/56812.htm   interface 断言  嵌入interface  反射
/*
interface
Go语言里面设计最精妙的应该算interface，它让面向对象，内容组织实现非常的方便，当你看完这一章，你就会被interface的巧妙设计所折服。
什么是interface
简单的说，interface是一组method的组合，我们通过interface来定义对象的一组行为。
我们前面一章最后一个例子中Student和Employee都能SayHi，虽然他们的内部实现不一样，但是那不重要，重要的是他们都能say hi
让我们来继续做更多的扩展，Student和Employee实现另一个方法Sing，然后Student实现方法BorrowMoney而Employee实现SpendSalary。
这样Student实现了三个方法：SayHi、Sing、BorrowMoney；而Employee实现了SayHi、Sing、SpendSalary。
上面这些方法的组合称为interface(被对象Student和Employee实现)。例如Student和Employee都实现了interface：SayHi和Sing，也就是这两个对象是该interface类型。而Employee没有实现这个interface：SayHi、Sing和BorrowMoney，因为Employee没有实现BorrowMoney这个方法。
interface类型
interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。详细的语法参考下面这个例子
*/

package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

//Human对象实现Sayhi方法
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s ,you can call me on %s \n", h.name, h.phone)
}

//Human对象实现Sing方法
func (h *Human) Sing(str string) {
	fmt.Println("lalala .. lalala ... ", str)
}

//Human对象实现Guzzle方法
func (h *Human) Guzzle(str string) {
	fmt.Println("guzzle guzzle .. guzzzle ..", str)
}

//Student对象实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
	s.BorrowMoney += amount
}

//Employee对象重载Human的Sayhi方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am is %s , work at %s,call me on %s\n", e.name, e.company, e.phone)
}

//Employee对象实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount
}

//定义接口interface
type IMen interface {
	SayHi()
	Sing(str string)
	Guzzle(str string)
}

type IYoungChap interface {
	SayHi()
	Sing(str string)
	BorrowMoney(amount float32)
}

type IElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}
