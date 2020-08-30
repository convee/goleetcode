package array

//35. 搜索插入位置
//https://leetcode-cn.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	l := len(nums)

	for i := 0; i < l; i++ {
		if target == nums[i] {
			return i
		}
	}

	for i := l - 1; i >= 0; i-- {
		if target < nums[0] {
			return 0
		}
		if target > nums[i] {
			return i + 1
		}
	}
	return l

}
