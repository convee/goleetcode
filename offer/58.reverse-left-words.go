package main

import "fmt"

//剑指 Offer 58 - II. 左旋转字符串
//https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
func reverseLeftWords(s string, n int) string {
	return string(s[n:] + s[:n])
}

func main() {
	fmt.Println(reverseLeftWords("abcdefg", 2))
}
