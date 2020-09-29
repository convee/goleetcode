package array

//350. 两个数组的交集 II
//给定两个数组，编写一个函数来计算它们的交集。
//https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/

//哈希表
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	m := map[int]int{}
	for _, num := range nums1 {
		m[num]++
	}

	intersection := []int{}
	for _, num := range nums2 {
		if m[num] > 0 {
			intersection = append(intersection, num)
			m[num]--
		}
	}
	return intersection
}
