package main

//反转字符串
//https://leetcode-cn.com/problems/reverse-string/
func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left <= right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
