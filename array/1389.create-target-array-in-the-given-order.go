package main

//1389. 按既定顺序创建目标数组
func createTargetArray(nums []int, index []int) []int {
	var target = make([]int, len(nums))
	for k, i := range index {
		copy(target[i+1:], target[i:])
		target[i] = nums[k]
	}
	return target
}

func createTargetArray2(nums []int, index []int) []int {
	res := make([]int, len(index))
	sortedNum := 0 //当前已经排列总数
	for i := 0; i < len(nums); i++ {
		v := index[i]
		if i == v {
			res[i] = nums[i]
		} else {
			for n := sortedNum; n > v; n-- { //需要把数字往后移动
				res[n] = res[n-1]
			}
			res[v] = nums[i]
		}
		sortedNum++
	}
	return res

}
