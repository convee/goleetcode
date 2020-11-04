package main

// 69. x 的平方根

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left := 1
	right := x / 2
	for left < right {
		mid := (left+right)/2 + 1
		if x/mid < mid {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}
