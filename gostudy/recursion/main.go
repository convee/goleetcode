package main

import "fmt"

//阶乘递归
//1*2*3*4*5*...*N
func Rescuvie(n int) int {
	if n == 0 {
		return 1
	}
	return n * Rescuvie(n-1)
}

//尾递归
func RescuvieTail(n int, a int) int {
	if n == 1 {
		return a
	}
	return RescuvieTail(n-1, a*n)
}

func main() {
	fmt.Println(Rescuvie(5))
	fmt.Println(RescuvieTail(5, 1))
}
