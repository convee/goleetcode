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

//斐波那契数列：后一个数是前两个数的和的一种数列
//1 1 2 3 5 8 13 21 ... N-1 N 2N-1
func F(n int, a1, a2 int) int {
	if n == 0 {
		return a1
	}
	return F(n-1, a2, a1+a2)
}

func main() {
	fmt.Println(Rescuvie(5))
	fmt.Println(RescuvieTail(5, 1))
	fmt.Println(F(5, 1, 1))
}
