package main

//1342. 将数字变成 0 的操作次数
//https://leetcode-cn.com/problems/number-of-steps-to-reduce-a-number-to-zero/
func numberOfSteps(num int) int {
	n := 0
	for {
		if num == 0 {
			return n
		} else if num%2 == 0 {
			n++
			num = num / 2
		} else {
			n++
			num = num - 1
		}
	}
}

func numberOfSteps2(num int) int {
	var res int
	for num != 0 {
		r := num % 2
		if r == 0 {
			num = num / 2
		} else {
			num = num - 1
		}
		res++
	}
	return res
}
