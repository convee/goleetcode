package main

// 1342. 将数字变成 0 的操作次数
// 给你一个非负整数 num ，请你返回将它变成 0 所需要的步数。 如果当前数字是偶数，你需要把它除以 2 ；否则，减去 1 。
// https://leetcode-cn.com/problems/number-of-steps-to-reduce-a-number-to-zero/

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
