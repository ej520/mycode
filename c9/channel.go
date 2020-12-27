// +build ignore

package main

import "fmt"

func main() {
	ChanFun5()
}

//ChanFun1
//有缓存的channel
func ChanFun1() {
	//定义一个有缓存的channel
	c1 := make(chan int, 5)

	//将数据1 存入chan c1
	c1 <- 1
	//将数据从chan c1中取出
	fmt.Println(<-c1)

}

//ChanFun2
//无缓存的channel
func ChanFun2() {
	//定义一个无缓存的channel
	c1 := make(chan int)

	//将数据1 存入chan c1
	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
		}

	}()
	//将数据从chan c1中取出
	for i := 0; i < 10; i++ {
		fmt.Println(<-c1)
	}

}

//ChanFun3
//可读可取、可读、可取的channel
func ChanFun3() {
	c1 := make(chan int, 5)
	var readc <-chan int = c1
	var writec chan<- int = c1
	writec <- 1
	fmt.Println(<-readc)
}

//chan close
func ChanFun4() {
	c1 := make(chan int, 5)
	c1 <- 1
	c1 <- 2
	c1 <- 3
	c1 <- 4
	c1 <- 5
	close(c1)

	for v := range c1 {
		fmt.Println(v)
	}

	// fmt.Println(<-c1)
	// fmt.Println(<-c1)
	// fmt.Println(<-c1)
	// fmt.Println(<-c1)
	// fmt.Println(<-c1)

}

// chan   select
func ChanFun5() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	ch1 <- 1
	ch1 <- 2

	//保证不会panic
	select {
	case <-ch1:
		fmt.Println("ch1:", <-ch1)
	case <-ch2:
		fmt.Println("ch2:", <-ch2)
	case <-ch3:
		fmt.Println("ch3:", <-ch3)
	default:
		fmt.Println("ALL is not valid")
	}
}
