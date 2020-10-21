package main

//88. 合并两个有序数组
//https://leetcode-cn.com/problems/merge-sorted-array/
//因为输入的两个数组自身是有序的，所以我们只需要从后往前比较两个数组取最大的那个数放到nums1的后面
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m + n - 1; i >= 0; i-- {
		//n为零时说明nums2中所有的元素已经放到nums1里面了
		if n == 0 {
			break
		}
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
	}

}
