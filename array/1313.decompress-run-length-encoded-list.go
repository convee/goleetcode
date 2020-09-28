package array

//1313. 解压缩编码列表
//https://leetcode-cn.com/problems/decompress-run-length-encoded-list/

//给你一个以行程长度编码压缩的整数列表 nums 。
// 考虑每对相邻的两个元素 [freq, val] = [nums[2*i], nums[2*i+1]] （其中 i >= 0 ），每一对都表示解压后子列表中有 freq 个值为 val 的元素，你需要从左到右连接所有子列表以生成解压后的列表。
// 请你返回解压后的列表。

func decompressRLElist(nums []int) []int {
	var res []int
	l := len(nums)
	i := 0
	for i < l {
		for 0 < nums[i] {
			nums[i] -= 1
			res = append(res, nums[i+1])
		}
		i += 2
	}
	return res
}

func decompressRLElist2(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i = i + 2 {
		for j := 0; j < nums[i]; j++ {
			result = append(result, nums[i+1])
		}
	}
	return result
}
