package main

//191. 位1的个数
//编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’ 的个数（也被称为汉明重量）。

//示例 1：
//
//输入：00000000000000000000000000001011
//输出：3
//解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。

func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		num = num & (num - 1)
		res++
	}
	return res
}

func hammingWeight2(num uint32) int {
	result := 0
	var mask uint32
	mask = 1 //初始化掩码
	for i := 0; i < 32; i++ {
		if num&mask != 0 {
			result++
		}
		mask <<= 1
	}
	return result
}
