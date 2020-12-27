//## +build ignore

package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	var w chan<- int = ch
	var r <-chan int = ch

	var wg sync.WaitGroup
	wg.Add(2)
	go Setchan1(w, &wg)
	go Getchan1(r, &wg)
	wg.Wait()

}

func Setchan1(w chan<- int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println("Setchan: Set", i)
		w <- i
	}
	wg.Done()
}

func Getchan1(r <-chan int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println("get info in chan w:", <-r)
	}
	wg.Done()
}
