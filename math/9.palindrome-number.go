package main

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 *
 * https://leetcode-cn.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (56.05%)
 * Total Accepted:    70.6K
 * Total Submissions: 125.9K
 * Testcase Example:  '121'
 *
 * 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
 *
 * 示例 1:
 *
 * 输入: 121
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: -121
 * 输出: false
 * 解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
 *
 *
 * 示例 3:
 *
 * 输入: 10
 * 输出: false
 * 解释: 从右向左读, 为 01 。因此它不是一个回文数。
 *
 *
 * 进阶:
 *
 * 你能不将整数转为字符串来解决这个问题吗？
 *
 */
func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	reverse := 0
	num := x
	for num != 0 {
		reverse = reverse*10 + num%10
		num = num / 10
	}
	if x == reverse {
		return true
	}
	return false
}

func isPalindrome2(x int) bool {
	str := fmt.Sprint(x)
	ttStr := []rune(str)
	lenStr := len(str)
	for i, j := 0, lenStr-1; i < j; i, j = i+1, j-1 {
		ttStr[i], ttStr[j] = ttStr[j], ttStr[i]
	}
	if str == string(ttStr) {
		return true
	}
	return false
}

func isPalindrome3(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 && x > 0 {
		return true
	}
	s := strconv.Itoa(x)
	length := len(s)
	for i := 0; i <= length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(isPalindrome(0))
}
