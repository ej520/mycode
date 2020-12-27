// +build ignore

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	var w chan<- int = ch
	var r <-chan int = ch
	go Setchan(w)

	Getchan(r)

}

func Setchan(w chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Setchan: Set", i)
		w <- i
	}
}

func Getchan(r <-chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("get info in chan w:", <-r)
	}
}
