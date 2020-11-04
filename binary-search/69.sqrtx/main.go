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

func mySqrt2(x int) int {
	left := 0
	right := x
	for left <= right {
		mid := (left + right) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			left = mid + 1
		} else if mid*mid > x {
			right = mid - 1
		}
	}
	return right
}
