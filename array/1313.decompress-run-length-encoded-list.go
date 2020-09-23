package array

//1313. 解压缩编码列表
//https://leetcode-cn.com/problems/decompress-run-length-encoded-list/
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
