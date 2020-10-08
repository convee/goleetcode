package gostudy

//给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从 0 开始)。如果不存在，则返回 -1。
//需要注意点
//循环时，i 不需要到 len-1
//如果找到目标字符串，len(needle)==j
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	var i, j int
	// i不需要到len-1
	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		// 判断字符串长度是否相等
		if len(needle) == j {
			return i
		}
	}
	return -1
}
