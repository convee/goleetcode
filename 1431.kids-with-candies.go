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

func main() {
	candies := []int{1, 2, 3, 4, 5}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies, extraCandies))
}
