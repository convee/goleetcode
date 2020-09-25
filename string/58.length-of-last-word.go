package string

//58. 最后一个单词的长度
//给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。

// 如果不存在最后一个单词，请返回 0 。

// 说明：一个单词是指仅由字母组成、不包含任何空格字符的 最大子字符串。

// 输入: "Hello World"
// 输出: 5

// 思路：
// 从后向前遍历字符串
// 末尾空格记得忽略
// 只统计第一次遇到的字符串即可
func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	l := len(s)
	out := 0
	for i := l - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			out++
		}
		if out != 0 && string(s[i]) == " " {
			break
		}
	}
	return out
}
