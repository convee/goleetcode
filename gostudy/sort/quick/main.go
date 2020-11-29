package main

import "fmt"

// 快速排序（前序遍历）
// 快速排序的逻辑是，若要对 nums[lo..hi] 进行排序，
// 我们先找一个分界点 p，通过交换元素使得 nums[lo..p-1] 都小于等于 nums[p]，
// 且 nums[p+1..hi] 都大于 nums[p]，
// 然后递归地去 nums[lo..p-1] 和 nums[p+1..hi] 中寻找新的分界点，最后整个数组就被排序了
func sort(nums []int, lo int, hi int) {
	if len(nums) <= 0 {
		return
	}
	if hi <= lo {
		return
	}
	// 临界点
	p := partition(nums, lo, hi)
	sort(nums, lo, p-1)
	sort(nums, hi, p+1)
}

// 交换位置，得到临界点位置
func partition(nums []int, lo int, hi int) int {
	v := nums[lo]
	k := lo
	for i := lo + 1; i <= hi; i++ {
		if nums[i] <= v {
			nums[k] = nums[i]
			nums[i] = nums[k+1]
			k++
		}
	}
	nums[k] = v
	return k

}

func main() {
	b := [...]int{8, 7, 5, 4, 3, 10, 15}
	sort(b[:], 0, len(b)-1)
	fmt.Println(b)
}
