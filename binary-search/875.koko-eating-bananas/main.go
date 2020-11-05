package main

// 875. 爱吃香蕉的珂珂
// https://leetcode-cn.com/problems/koko-eating-bananas/

// 输入: piles = [3,6,7,11], H = 8
// 输出: 4
func minEatingSpeed(piles []int, H int) int {
	// 最少为 1，最大为 max(piles)，因为一小时最多只能吃一堆香蕉
	left, right := 1, getMax(piles)+1
	for left < right {
		mid := (left + right) / 2
		if canFinish(piles, H, mid) { // 以 speed 是否能在 H 小时内吃完香蕉
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func getMax(piles []int) (maxPile int) {
	maxPile = piles[0]
	for _, x := range piles {
		if x > maxPile {
			maxPile = x
		}
	}
	return
}

func canFinish(piles []int, H, speed int) bool {
	time := 0
	for _, pile := range piles {
		time += timeOf(pile, speed)
	}
	return time <= H
}

func timeOf(pile, speed int) int { // 每堆的耗时
	if pile%speed == 0 && pile > speed {
		return pile / speed
	} else if pile <= speed {
		return 1
	} else {
		return pile/speed + 1
	}
}
