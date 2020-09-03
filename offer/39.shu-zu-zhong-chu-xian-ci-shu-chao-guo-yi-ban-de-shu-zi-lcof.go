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

func majorityElement2(nums []int) int {
	x, votes := nums[0], 1 // x 为假设的众数，votes为票数
	for i := 1; i < len(nums); i++ {
		if votes == 0 { // 票数归零，重新设当前元素为众数，准备开始新一轮的计票
			x = nums[i]
		}

		if nums[i] == x { // 若当前数字等于众数x，则票数+1，否则-1
			votes++
		} else {
			votes--
		}
	}

	return x
}
