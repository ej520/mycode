//## +build ignore

package main

import "fmt"

type Animal interface {
	Eat()
	Run()
}

type Cat struct {
	Name string
	Sex  string
}
type Dog struct {
	Name string
}

func (c *Cat) Run() {
	fmt.Println(c.Name, "start to run ....")
}

func (c *Cat) Eat() {
	fmt.Println(c.Name, "start to eat ...")
}

func (d Dog) Eat() {
	fmt.Println(d.Name, "start to eat")
}
func (d Dog) Run() {
	fmt.Println(d.Name, "start to run ...")
}

func MyFun(a Animal) {
	a.Run()
	a.Eat()
}

var L Animal

func MyFun2(a Animal) {
	L = a
}

// func main() {
// 	var a Animal

// 	c := Cat{
// 		Name: "tom",
// 		Sex:  "male",
// 	}

// 	d := Dog{
// 		Name: "gougou",
// 	}
// 	a = &c
// 	fmt.Println(a, c)
// 	a.Run()
// 	a.Eat()

// 	a = d
// 	fmt.Println(a, d)
// 	MyFun(a)
// }

func main() {
	c := Cat{
		Name: "tom",
		Sex:  "male",
	}
	MyFun2(&c)
	L.Run()

}
