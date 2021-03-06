package main

// 239. 滑动窗口最大值
// https://leetcode-cn.com/problems/sliding-window-maximum/

// 思路：暴力解法 O(nk)
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) <= 0 {
		return []int{}
	}
	l := len(nums) - k + 1
	res := []int{}
	for i := 0; i < l; i++ {
		mx := -1 << 31
		for j := i; j < k+i; j++ {
			mx = max(mx, nums[j])
		}
		res = append(res, mx)
	}
	return res
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 思路：双端队列 O(n)

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	//用切片模拟一个双端队列
	queue := []int{}
	result := []int{}
	for i := range nums {
		for i > 0 && (len(queue) > 0) && nums[i] > queue[len(queue)-1] {
			//将比当前元素小的元素祭天
			queue = queue[:len(queue)-1]
		}
		//将当前元素放入queue中
		queue = append(queue, nums[i])
		if i >= k && nums[i-k] == queue[0] {
			//维护队列，保证其头元素为当前窗口最大值
			queue = queue[1:]
		}
		if i >= k-1 {
			//放入结果数组
			result = append(result, queue[0])
		}
	}
	return result
}
