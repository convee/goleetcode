package main

//392. 判断子序列
//给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//s = "abc", t = "ahbgdc"

//思路：双指针

func isSubsequence(s string, t string) bool {
	slen := len(s)
	tlen := len(t)
	i, j := 0, 0
	for j < tlen && i < slen {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == slen
}
