package array

//217. 存在重复元素
//https://leetcode-cn.com/problems/contains-duplicate/
// 给定一个整数数组，判断是否存在重复元素。
// 如果任意一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。
//暴力解法
func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
