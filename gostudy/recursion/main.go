package main

import "fmt"

//阶乘递归
//1*2*3*4*5*...*N
func Rescuvie(n int) int {
	if n == 0 {
		return 1
	}
	return n * Rescuvie(n-1)
}

//尾递归
func RescuvieTail(n int, a int) int {
	if n == 1 {
		return a
	}
	return RescuvieTail(n-1, a*n)
}

//斐波那契数列：后一个数是前两个数的和的一种数列
//1 1 2 3 5 8 13 21 ... N-1 N 2N-1
func F(n int, a1, a2 int) int {
	if n == 0 {
		return a1
	}
	return F(n-1, a2, a1+a2)
}

// 二分查找递归解法
func BinarySearch(array []int, target int, l, r int) int {
	if l > r {
		// 出界了，找不到
		return -1
	}
	// 从中间开始找
	mid := (l + r) / 2
	middleNum := array[mid]
	if middleNum == target {
		return mid // 找到了
	} else if middleNum > target {
		// 中间的数比目标还大，从左边找
		return BinarySearch(array, target, 1, mid-1)
	} else {
		// 中间的数比目标还小，从右边找
		return BinarySearch(array, target, mid+1, r)
	}
}

// 二分查找非递归解法
func BinarySearch2(array []int, target int, l, r int) int {
	ltemp := l
	rtemp := r
	for {
		if ltemp > rtemp {
			// 出界了，找不到
			return -1
		}
		// 从中间开始找
		mid := (ltemp + rtemp) / 2
		middleNum := array[mid]
		if middleNum == target {
			return mid // 找到了
		} else if middleNum > target {
			// 中间的数比目标还大，从左边找
			rtemp = mid - 1
		} else {
			// 中间的数比目标还小，从右边找
			ltemp = mid + 1
		}
	}
}
func main() {
	fmt.Println(Rescuvie(5))
	fmt.Println(RescuvieTail(5, 1))
	fmt.Println(F(5, 1, 1))
	array := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
	target := 500
	result := BinarySearch(array, target, 0, len(array)-1)
	fmt.Println(target, result)
	target = 189
	result = BinarySearch(array, target, 0, len(array)-1)
	fmt.Println(target, result)
}
