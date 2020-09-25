package math

import (
	"math"
)

//633. 平方数之和
//给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c。
// 双指针法
// 解题思路
// 1、求出c的开方根，可以用循环的方法求，也可以用sqrt系统函数求，假设根为max
// 2、设定左右指针，分别指向1和max
// 3、求左右指针的平方和，恰好等于c,返回true，如果小于max，右移左指针；否则左移右指针。
// 4、直到左右指针相遇，还没找到相等的，返回false

func judgeSquareSum(c int) bool {
	max := int(math.Sqrt(float64(c)))
	left, right := 0, max
	for left <= max {
		sum := left*left + right*right
		if sum == c {
			return true
		}
		if sum < c {
			left++
		} else {
			right--
		}
	}
	return false
}
