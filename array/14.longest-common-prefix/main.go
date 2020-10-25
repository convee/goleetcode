package main

import "strings"

//14. 最长公共前缀
//https://leetcode-cn.com/problems/longest-common-prefix/
//思路：
//将第一个元素设置为基准元素x0
//依次将基准元素和后面的元素进行比较（假定后面的元素依次为x1,x2,x3....），
//不断更新基准元素，直到基准元素和所有元素都满足最长公共前缀的条件，就可以得到最长公共前缀。
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	//将第一个元素设置为基准元素x0
	prefix := strs[0]
	for _, k := range strs {
		for strings.Index(k, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
func longestCommonPrefix2(strs []string) string {
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
