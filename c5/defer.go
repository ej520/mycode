// +build ignore

package main

import (
	"fmt"
	"time"
)

//延迟调用 defer： 无论放哪里都是最后调用

func mo() {
	fmt.Println("我想最先执行！！！")
}

func sp() {
	for i := 0; i < 15; i++ {
		fmt.Println(" 我在睡觉")
		time.Sleep(time.Second * 1)
	}

}

func main() {
	//延迟调用， 最后执行
	defer mo()
	fmt.Println("1")
	fmt.Println("2")
	go sp()
	time.Sleep(time.Second * 10)
	fmt.Println("3")

}
