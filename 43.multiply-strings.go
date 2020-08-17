package main

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 *
 * https://leetcode-cn.com/problems/multiply-strings/description/
 *
 * algorithms
 * Medium (37.38%)
 * Total Accepted:    9.4K
 * Total Submissions: 25K
 * Testcase Example:  '"2"\n"3"'
 *
 * 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
 *
 * 示例 1:
 *
 * 输入: num1 = "2", num2 = "3"
 * 输出: "6"
 *
 * 示例 2:
 *
 * 输入: num1 = "123", num2 = "456"
 * 输出: "56088"
 *
 * 说明：
 *
 *
 * num1 和 num2 的长度小于110。
 * num1 和 num2 只包含数字 0-9。
 * num1 和 num2 均不以零开头，除非是数字 0 本身。
 * 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
 *
 *
 */
func multiply(num1 string, num2 string) string {
	pos := make([]int64, len(num1)+len(num2))
	ans := ""

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			mul := (num1[i] - '0') * (num2[j] - '0')
			p1, p2 := i+j, i+j+1
			sum := int64(mul) + pos[p2]

			pos[p1] += sum / 10
			pos[p2] = sum % 10

		}
	}

	for _, v := range pos {
		ans += strconv.Itoa(int(v))
	}
	ans = strings.TrimLeft(ans, "0")
	if len(ans) == 0 {
		return "0"
	}

	return ans
}
