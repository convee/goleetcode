package main

/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 *
 * https://leetcode-cn.com/problems/permutation-in-string/description/
 *
 * algorithms
 * Medium (31.59%)
 * Total Accepted:    2.9K
 * Total Submissions: 9.2K
 * Testcase Example:  '"ab"\n"eidbaooo"'
 *
 * 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
 *
 * 换句话说，第一个字符串的排列之一是第二个字符串的子串。
 *
 * 示例1:
 *
 *
 * 输入: s1 = "ab" s2 = "eidbaooo"
 * 输出: True
 * 解释: s2 包含 s1 的排列之一 ("ba").
 *
 *
 *
 *
 * 示例2:
 *
 *
 * 输入: s1= "ab" s2 = "eidboaoo"
 * 输出: False
 *
 *
 *
 *
 * 注意：
 *
 *
 * 输入的字符串只包含小写字母
 * 两个字符串的长度都在 [1, 10,000] 之间
 *
 *
 */
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	for i := 0; i < len(s2)-len(s1)+1; i++ {
		if isAnagram(s2[i:i+len(s1)], s1) {
			return true
		}
	}
	return false
}
func isAnagram(s string, t string) bool {
	var xor, squaresum1, squaresum2 rune
	for _, char := range s {
		xor ^= char
		squaresum1 += char * char
	}
	for _, char := range t {
		xor ^= char
		squaresum2 += char * char
	}
	return xor == 0 && squaresum1 == squaresum2
}
