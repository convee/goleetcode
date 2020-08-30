package main

import "fmt"

//1486. 数组异或操作
//https://leetcode-cn.com/problems/xor-operation-in-an-array/
func xorOperation(n int, start int) int {
	result := 0
	for i := 1; i < n; i++ {
		result = result ^ (start + 2*i)
	}
	return result
}

func main() {
	fmt.Println(xorOperation(5, 0))
}
