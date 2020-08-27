package goleetcode

//1295. 统计位数为偶数的数字
//https://leetcode-cn.com/problems/find-numbers-with-even-number-of-digits/
func findNumbers(nums []int) int {
	n := 0
	for _, v := range nums {
		cur := 0
		for 0 < v {
			v /= 10
			cur++
		}
		if cur%2 == 0 {
			n++
		}
	}
	return n
}
