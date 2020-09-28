package array

//189. 旋转数组
//https://leetcode-cn.com/problems/rotate-array/
//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
// 解题思路
// 1、通过 k 把数组分为两个部分
// 2、 k 在数组切割的位置 len(nums) - k
// 3、两部分调换位置
// 4、 k 的值可能大于数组长度要对其取余

func rotate(nums []int, k int) {
	k = k % len(nums)
	copy(nums, append(nums[len(nums)-k:], nums[0:len(nums)-k]...))
}
