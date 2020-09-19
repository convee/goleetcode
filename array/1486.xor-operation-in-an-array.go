package main

import "fmt"

//1486. 数组异或操作
//https://leetcode-cn.com/problems/xor-operation-in-an-array/

//给你两个整数，n 和 start 。

//数组 nums 定义为：nums[i] = start + 2*i（下标从 0 开始）且 n == nums.length 。

//请返回 nums 中所有元素按位异或（XOR）后得到的结果。

func xorOperation(n int, start int) int {
	result := 0
	for i := 1; i < n; i++ {
		result = result ^ (start + 2*i)
	}
	return result
}

func xorOperation2(n int, start int) int {
	if (start & 3) < 2 {
		if (n & 1) == 0 {
			return n & 3
		} else {
			return start + 2*n - 3 + (n & 3)
		}
	} else {
		if (n & 1) == 0 {
			return (start + (n-1)*2) ^ (start - 2 + (n & 3))
		} else {
			return start + 1 - (n & 3)
		}
	}
}

func main() {
	fmt.Println(xorOperation(5, 0))
}
