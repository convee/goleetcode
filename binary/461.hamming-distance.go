package binary

//461. 汉明距离

//两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
//
//给出两个整数 x 和 y，计算它们之间的汉明距离。
//
//注意：
//0 ≤ x, y < 231.


func hammingDistance(x int, y int) int {
	z:=x^y
	sum:=0
	for z != 0 {
		sum += z&1
		z = z>>1
	}
	return sum
}