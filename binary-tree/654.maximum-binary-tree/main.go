package main

// 654. 最大二叉树
// https://leetcode-cn.com/problems/maximum-binary-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 主函数
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(0, len(nums)-1, nums)
}

// 将 nums[lo..hi] 构造成符合条件的树，返回根节点
func build(lo int, hi int, nums []int) *TreeNode {
	if lo > hi {
		return nil
	}
	// 找到数组中的最大值和对应的索引
	maxVal := nums[lo]
	index := lo
	for i := lo; i <= hi; i++ {
		if maxVal < nums[i] {
			index = i
			maxVal = nums[i]
		}
	}
	// 递归调用构造左右子树
	return &TreeNode{
		Val:   maxVal,
		Left:  build(lo, index-1, nums),
		Right: build(index+1, hi, nums),
	}

}
