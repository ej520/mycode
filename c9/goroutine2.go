// +build ignore

package main

import (
	"fmt"
	"sync"
)

//协程管理器，无需sleep和for循环
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go GoRun2(&wg)
	go GoRun1(&wg)
	//wg 2 变成0 时结束
	wg.Wait()
}
func GoRun1(wg *sync.WaitGroup) {
	fmt.Println("GoRun1 going to run ...")
	wg.Done()
}

func GoRun2(wg *sync.WaitGroup) {
	fmt.Println("GoRun2 going to run ...")
	wg.Done()
}
