package main

import "sort"

// 第350题：两个数组的交集
// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
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

// 相比上一个方法利用 nums数组，执行用时和内存消耗没有明显区别

func intersect2(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	// 遍历 nums1 初始化 map
	for _, v := range nums1 {
		m[v]++
	}
	k := 0
	for _, v := range nums2 {
		// 如果元素相同，将其存入 nums2(提示：解答中我们并没有创建空白数组，因为遍历后的数组其实就没用了。我们可以将相等的元素放入用过的数组中，就为我们节省下了空间。)
		if m[v] > 0 {
			m[v]--
			nums2[k] = v
			k++
		}
	}
	return nums2[0:k]
}

//进阶:
//如果给定的数组已经排好序呢？将如何优化你的算法呢？
//思路：双指针解法
//解题步骤：
// <1> 设定两个为0的指针，比较两个指针的元素是否相等。如果指针的元素相等，我们将两个指针一起
// 向后移动，并且将相等的元素放入空白数组。下图中我们的指针分别指向第一个元素，判断元素相等之
// 后，将相同元素放到空白的数组。
// <2> 如果两个指针的元素不相等，我们将小的一个指针后移。图中我们指针移到下一个元素，判断不相
// 等之后，将元素小的指针向后移动，继续进行判断。
//<3> 反复以上步骤。

func intersect3(nums1 []int, nums2 []int) []int {
	i, j, k := 0, 0, 0
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}
