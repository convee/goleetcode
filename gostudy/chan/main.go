package main

import "fmt"



// 通过形参限定函数内部只能从channel中读取数据
func readChan(chanName <-chan int) {
	c := <-chanName
	fmt.Println(c)
}

// 通过形参限定函数内部只能向channel中写入数据
func writeChan(chanName chan<- int) {
	chanName <- 1
}
func main() {
	var mychan = make(chan int, 10)
	writeChan(mychan)
	readChan(mychan)
}
