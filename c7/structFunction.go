// +build ignore

package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Sex    bool
	Hobbys []string
	Home
}
type Home struct {
	P string
}

//指针会改变外面值
func (u *User) Myfunc1(name string) {
	fmt.Println(name, "的年纪", u.Age)
	u.Age += u.Age
}

//值传递不改变外面值
func (u User) Myfunc2(name string) {
	fmt.Println(name, "的年纪", u.Age)
	u.Age += u.Age
}

func main() {
	u1 := User{
		Name: "ej520",
		Age:  22,
	}

	u2 := User{
		Name: "abc",
		Age:  22,
	}

	//22 44 88 指针会改变外面值，每次累加
	u1.Myfunc1("aaa")
	u1.Myfunc1("aaa")
	u1.Myfunc1("aaa")

	//22 22 22 值传递不会改变值，每次一样
	u2.Myfunc2("bbb")
	u2.Myfunc2("bbb")
	u2.Myfunc2("bbb")
}
