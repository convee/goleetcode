package main

import "fmt"

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (27.93%)
 * Total Accepted:    76K
 * Total Submissions: 272.2K
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 * 示例 1:
 *
 * 输入: "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

 */

//自己假设更为一般的 case: abcabd，在遇到第二个 a 时最长子串从 abc 更新为了 bca，接着从 bca 更新为了 cab，最后发现最长子串 cabd
//类比 TCP 做流量控制的滑动窗口，向后遍历时若遇到重复字符，则丢弃重复字符之前的部分，并记录最大子串长即可
func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	start := -1 // 指向当前不重复字符串的起始位置
	m := make(map[rune]int)
	for k, v := range s {
		if last, ok := m[v]; ok && start < last { //字符再次出现
			start = last
		}
		if k-start > maxLen {
			maxLen = k - start
		}
		m[v] = k
	}
	return maxLen
}

func lengthOfLongestSubstring2(s string) int {
	res := 0
	i, j := 0, 0
	strs := []rune(s)
	l := len(strs)
	m := make(map[rune]int)
	for i < l && j < l {
		if _, ok := m[strs[j]]; !ok {
			m[strs[j]] = j
			j++
			res = max(res, j-i)
		} else {
			delete(m, strs[i])
			i++
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {

	fmt.Println(lengthOfLongestSubstring2("pwwkew"))
}
