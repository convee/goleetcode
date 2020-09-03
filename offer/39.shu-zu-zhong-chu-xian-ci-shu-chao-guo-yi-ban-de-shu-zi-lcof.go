package offer

//剑指 Offer 39. 数组中出现次数超过一半的数字
//https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
func majorityElement(nums []int) int {
	m := make(map[int]int)
	t := len(nums) / 2
	for _, v := range nums {
		if m[v] >= t {
			return v
		}
		m[v]++
	}
	return -1
}
