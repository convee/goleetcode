package string

//1221. 分割平衡字符串
//https://leetcode-cn.com/problems/split-a-string-in-balanced-strings/
func balancedStringSplit(s string) int {
	if len(s) == 0 {
		return 0
	}
	stack := []byte{s[0]}
	var res int
	for i := 1; i < len(s); i++ {
		if len(stack) > 0 && ((s[i] == 'L' && stack[len(stack)-1] == 'R') || (s[i] == 'R' && stack[len(stack)-1] == 'L')) {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				res++
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return res
}
