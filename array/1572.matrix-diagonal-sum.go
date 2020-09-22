package main

//1572. 矩阵对角线元素的和
//https://leetcode-cn.com/problems/matrix-diagonal-sum/

//方法一：无论奇偶都累计相加，最后减去重复
func diagonalSum(mat [][]int) int {
	l := len(mat)
	sum := 0

	for i, j := 0, l-1; i < l; i, j = i+1, j-1 {
		sum += mat[i][i] + mat[i][j]
	}
	if l%2 == 0 {
		return sum
	}
	m := (l - 1) / 2
	return sum - mat[m][m]
}

func diagonalSum2(mat [][]int) int {
	res, n := 0, len(mat)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || i+j == n-1 {
				res += mat[i][j]
			}
		}
	}
	return res
}
