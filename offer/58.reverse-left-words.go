package main

import "fmt"

//剑指 Offer 58 - II. 左旋转字符串
//https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
func reverseLeftWords(s string, n int) string {
	return string(s[n:] + s[:n])
}

func reverseLeftWords2(s string, n int) string {
	var str string
	var str2 string
	for i := 0; i < len(s); i++ {
		if i < n {
			str += string(s[i])
		} else {
			str2 += string(s[i])
		}
	}
	return str2 + str
}

func main() {
	fmt.Println(reverseLeftWords2("abcdefg", 2))
}
