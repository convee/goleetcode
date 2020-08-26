package main

import "fmt"

//771.宝石与石头
//https://leetcode-cn.com/problems/jewels-and-stones/

func numJewelsInStones(J string, S string) int {
	n := 0
	h := make(map[rune]bool)
	for _, v := range J {
		h[v] = true
	}
	for _, v := range S {
		if _, ok := h[v]; ok {
			n++
		}
	}
	return n
}

func main() {
	J := "aA"
	S := "aaaBCDAaA"
	fmt.Println(numJewelsInStones(J, S))
}
