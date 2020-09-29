package array

//136. 只出现一次的数字
//https://leetcode-cn.com/problems/single-number/
//位运算
// 任何数和 00 做异或运算，结果仍然是原来的数，即 a ^ 0 = a。
// 任何数和其自身做异或运算，结果是 00，即 a ^ a = 0。
// 异或运算满足交换律和结合律，即 a ^ b ^ a=b ^ a ^ a=b ^ (a ^ a)=b ^0 = b。

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
