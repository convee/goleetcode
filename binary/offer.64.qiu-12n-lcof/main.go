package main

var res int

// 剑指 Offer 64. 求1+2+…+n
// 求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
// https://leetcode-cn.com/problems/qiu-12n-lcof/
// 思路：递归
// 利用 && 特性，A&&B

func sumNums(n int) int {
	res = 0
	sum(n)
	return res
}

func sum(n int) bool {
	res += n
	return n > 0 && sum(n-1)
}
