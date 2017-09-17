package main

import (
	"fmt"
)

func map1() {
	//类似hash或者字典 。 key-value  key 必须支持=或者!=
	//优先选择数组和slice
	var n map[int]string = make(map[int]string) //定义方式1
	m := make(map[int]string)                   //定义方式2
	m[1] = "OK"                                 //add key=value => 1=ok
	m[1] = "OKKKK"                              //modify key:1
	//delete(m, 1)                                //delete key:1
	a := m[1] //get key:1
	fmt.Println(a)
	fmt.Println(m, n)

}

func main() {

	map1()
}
