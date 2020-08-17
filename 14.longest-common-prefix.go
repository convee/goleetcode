package main

import "strings"

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (31.68%)
 * Total Accepted:    49.3K
 * Total Submissions: 155.7K
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 * 示例 1:
 *
 * 输入: ["flower","flow","flight"]
 * 输出: "fl"
 *
 *
 * 示例 2:
 *
 * 输入: ["dog","racecar","car"]
 * 输出: ""
 * 解释: 输入不存在公共前缀。
 *
 *
 * 说明:
 *
 * 所有输入只包含小写字母 a-z 。
 *
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	s1 := strs[0]
	prefixs := make([]string, len(s1))
	for i := range s1 {
		prefixs[i] = s1[:i+1]
	}
	for _, s := range strs {
		for i, pre := range prefixs {
			if !strings.HasPrefix(s, pre) {
				prefixs = prefixs[:i]
				break
			}
		}
	}
	if len(prefixs) > 0 {
		return prefixs[len(prefixs)-1]
	}
	return ""
}
