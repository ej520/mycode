// +build ignore

package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Sex    bool
	Hobbys []string
}

func main() {
	//struct声明方式1
	var u User
	u.Name = "wyj"
	u.Age = 4
	u.Hobbys = []string{"玩"}
	fmt.Println(u)

	//struct声明方式2
	uu := User{"w", 12, true, []string{"学习"}}
	fmt.Println(uu)

	//struct声明方式3
	uuu := User{
		Name:   "abc",
		Age:    33,
		Sex:    true,
		Hobbys: []string{"xxx"},
	}
	fmt.Println(uuu)

	//struct声明方式4, 该方式直接获得是地址
	u4 := new(User)
	u4.Name = "Peter"
	fmt.Println(u4)

	//调用方法
	PrintUser(u)
	PrintUser(uu)
	PrintUser(uuu)
	PrintUser(*u4)
}

func PrintUser(u User) {
	fmt.Println(u)
}
