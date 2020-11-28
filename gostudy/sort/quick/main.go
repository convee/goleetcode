package main

// 快速排序（前序遍历）
func sort(nums []int, lo int, hi int) {
	// 临界点 p
	p := partition(nums, lo, hi)
	sort(nums, lo, p-1)
	sort(nums, hi, p+1)
}

// 交换位置
func partition(nums []int, lo int, hi int) int {

}
