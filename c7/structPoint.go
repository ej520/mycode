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

func (h *Home) Open() {
	fmt.Println("Home=>OPEN():", h.P)
}

func StructPoint() {
	//struct声明方式3
	uuu := User{
		Name:   "abc",
		Age:    33,
		Sex:    true,
		Hobbys: []string{"xxx"},
	}
	fmt.Println(uuu)

	var u3p *User
	u3p = &uuu

	//复杂对象 不加可以*修改
	u3p.Name = "Name-New-ABC"
	(*u3p).Age = 10
	fmt.Println(*u3p)
}

//结构体的方法
func (u *User) Song(name string) (ret string) {
	ret = u.Name + "正在唱" + name
	return
}

func main() {
	StructPoint()

	//唱歌
	u := User{
		Name:   "东方不败",
		Age:    33,
		Sex:    true,
		Hobbys: []string{"xxx"},
	}

	u.P = "芝麻开门"
	u.Open()

	fmt.Println(u.Song("东方红"))

	ua := User{
		Name: "西方",
	}
	fmt.Println(ua)
	fmt.Println(ua.Song("夕阳西下"))
}
