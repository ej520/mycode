//from : http://blog.csdn.net/qwertyupoiuytr/article/details/55101258  Go语言goroutine+channel+select简介
/*
goroutine：
Go语言是原生支持语言级并发的，这个并发的最小逻辑单元就是goroutine。
goroutine就是Go语言提供的一种用户态线程，这种用户态线程是跑在内核级线程之上的，goroutine在运行时的调度是由Go语言提供的调度器来进行的，创建一个goroutine使用关键字go，go创建的goroutine不会阻塞主线程：
go func_name()
Go程序中没有语言级的关键字让你去创建一个内核线程，你只能创建goroutine，内核线程只能由调度器根据实际情况去创建。
调度器原理参考链接：http://studygolang.com/articles/1855

channel：
Channel是Go中的一个核心类型，可以把它看成一个管道，通过它可以发送或者接收数据进行通讯(communication)。配合goroutine，就形成了一种既简单又强大的请求处理模型，即N个工作goroutine将处理的中间结果或者最终结果放入一个channel，另外有M个工作goroutine从这个channel拿数据，再进行进一步加工，通过组合这种过程，可以胜任各种复杂的业务模型。
channel的基本操作：
var c chan int   //声明一个int类型的channel，注意，该语句仅声明，不初始化channel
c := make(chan type_name)   //创建一个无缓冲的type_name型的channel，无缓冲的channel当放入1个元素后，后续的输入便会阻塞
c := make(chan type_name, 100)   //创建一个缓冲区大小为100的type_name型的channel
c <- x   //将x发送到channel c中，如果channel缓冲区满，则阻塞当前goroutine
<- c   //从channel c中接收一个值，如果缓冲区为空，则阻塞
x = <- c   //从channel c中接收一个值并存到x中，如果缓冲区为空，则阻塞
x, ok = <- c   //从channel c中接收一个值，如果channel关闭了，那么ok为false（在没有defaultselect语句的前提下），在channel未关闭且为空的情况下，仍然阻塞
close(c)   //关闭channel c
for term := range c {}   //等待并取出channelc中的值，直到channel关闭，会阻塞
单向channel：
var ch1 chan<- float64    //只能向里面写入float64的数据，不能读取
var ch2 <-chan int        //只能读取int型数据

select：
select用于在多个channel上同时进行侦听并收发消息，当任何一个case满足条件时即执行，如果没有可执行的case则会执行default的case，如果没有指定defaultcase，则会阻塞程序，select的语法：
select {
   case communication clause :
      statement(s);
   case communication clause :
      statement(s);
   //可以定义任意数量的 case
   default :
      statement(s);
}
说明：
每个case都必须是一次通信
所有channel表达式都会被求值
所有被发送的表达式都会被求值
如果任意某个通信可以进行，它就执行；其他被忽略。
如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
否则： 1.如果有default子句，则执行该语句。
如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。
*/

package main

import (
	"fmt"
	"time"
)

func enqueue(q chan int) {
	time.Sleep(time.Second * 3)
	q <- 10
	q <- 20
	//close(q)
}

func main() {
	var c = make(chan int)
	go enqueue(c)
	for {
		select {
		case x, ok := <-c:
			if ok {
				fmt.Println(x)
			} else {
				fmt.Println("closed")
				return
			}
		default:
			fmt.Println("waiting")
			time.Sleep(time.Second)
		}
	}
}
