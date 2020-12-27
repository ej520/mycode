// +build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	go GoRun()

	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
func GoRun() {
	fmt.Println("going to run ...")
}
