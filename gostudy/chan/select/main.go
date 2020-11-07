package main

import (
	"fmt"
	"time"
)

func addNumberToChan(chanName chan int) {
	for {
		chanName <- 1
		time.Sleep(1 * time.Second)
	}
}

// 通过range可以持续从channel中读出数据，好像在遍历一个数组一样，
// 当channel中没有数据时会阻塞当前goroutine，与读channel时阻塞处理机制一样
// 注意：如果向此channel写数据的goroutine退出时，系统检测到这种情况后会panic，否则range将会永久阻塞
func chanRange(chanName chan int) {
	for e := range chanName {
		fmt.Printf("Get element from chan: %d\n", e)
	}
}

// 使用select可以监控多channel，比如监控多个channel，当其中某一个channel有数据时，就从其读出数据。
// 从channel中读出数据的顺序是随机的，事实上select语句的多个case执行顺序是随机的

// select的case语句读channel不会阻塞，尽管channel中没有数据。
// 这是由于case语句编译后调用读channel时会明确传入不阻塞的参数，此时读不到数据时不会将当前goroutine加入到等待队列，而是直接返回。

func main() {
	var chan1 = make(chan int, 10)
	var chan2 = make(chan int, 10)
	go addNumberToChan(chan1)
	go addNumberToChan(chan2)
	chanRange(chan1)

	for {
		select {
		case e := <-chan1:
			fmt.Printf("Get element from chan1: %d\n", e)
		case e := <-chan2:
			fmt.Printf("Get element from chan2: %d\n", e)
		default:
			fmt.Printf("No element in chan1 and chan2.\n")
			time.Sleep(1 * time.Second)
		}
	}
}
