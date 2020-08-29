package main

//1534. 统计好三元组
//https://leetcode-cn.com/problems/count-good-triplets/
func countGoodTriplets(arr []int, a int, b int, c int) int {
	count := 0
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if absolute(arr[i]-arr[j]) <= a && absolute(arr[j]-arr[k]) <= b && absolute(arr[i]-arr[k]) <= c {
					count++
				}
			}
		}
	}
	return count
}
func absolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

}
