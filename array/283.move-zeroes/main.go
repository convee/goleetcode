package main

func moveZeroes(nums []int) {
	l := len(nums) //递减非零计数器
	i := 0         //递增非零计数器
	for {
		if i >= l {
			break
		}
		if nums[i] == 0 {
			l = l - 1                               //遇到一个0，减少1，最后与i比较作为结束条件
			nums = append(nums[0:i], nums[i+1:]...) //把0前后两部分合并
			nums = append(nums, 0)                  //在末尾补回0
		} else {
			i = i + 1 //非0计数器自增
		}
	}
	return

}
