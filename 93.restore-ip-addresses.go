package main

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原IP地址
 *
 * https://leetcode-cn.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (42.28%)
 * Total Accepted:    4.8K
 * Total Submissions: 11.4K
 * Testcase Example:  '"25525511135"'
 *
 * 给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
 *
 * 示例:
 *
 * 输入: "25525511135"
 * 输出: ["255.255.11.135", "255.255.111.35"]
 *
 */
func restoreIpAddresses(s string) (res []string) {
	buildRes(&res, s, make([]string, 0), 0)
	return
}

func buildRes(res *[]string, s string, str []string, index int) {
	if len(str) == 4 {
		if index == len(s) {
			*res = append(*res, strings.Join(str, "."))
		}
		return
	}

	for i := index; i < index+3 && i < len(s); i++ {
		candidate := s[index : i+1]
		if num, _ := strconv.Atoi(candidate); num > 255 {
			break
		}

		str = append(str, candidate)
		buildRes(res, s, str, i+1)
		str = str[:len(str)-1]

		if candidate == "0" {
			break
		}
	}
}
