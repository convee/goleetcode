package main

import "fmt"

//拥有最多糖果的孩子
//https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies/
func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	result := make([]bool, n)
	if n <= 1 {
		result[0] = true
		return result
	}
	maxValue := 0
	for _, v := range candies {
		if maxValue < v {
			maxValue = v
		}
	}
	for j, v := range candies {
		result[j] = (v+extraCandies >= maxValue)
	}
	return result
}

func kidsWithCandies2(candies []int, extraCandies int) []bool {
	n := len(candies)
	maxCandies := 0
	for i := 0; i < n; i++ {
		maxCandies = max(maxCandies, candies[i])
	}
	ret := make([]bool, n)
	for i := 0; i < n; i++ {
		ret[i] = candies[i]+extraCandies >= maxCandies
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	candies := []int{1, 2, 3, 4, 5}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies, extraCandies))
}
