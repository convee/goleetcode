package main

func subtractProductAndSum(n int) int {
	j := 1
	h := 0
	for n > 0 {
		y := n % 10
		n = n / 10
		j = j * y
		h = h + y
	}
	return j - h
}
