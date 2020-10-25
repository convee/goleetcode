package main

// 第350题：两个数组的交集
// 思路：
// 设定两个为0的指针，比较两个指针的元素是否相等。
// 如果指针的元素相等，我们将两个指针一起向后移动，并且将相等的元素放入空白数组。

// 首先拿到这道题，我们基本马上可以想到，此题可以看成是一道传统的映射题（map映射），为
// 什么可以这样看呢，因为我们需找出两个数组的交集元素，同时应与两个数组中出现的次数一
// 致。这样就导致了我们需要知道每个值出现的次数，所以映射关系就成了<元素,出现次数>。剩下
// 的就是顺利成章的解题。

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
