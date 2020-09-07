package stl

import "sort"

func main() {
	// int排序
	sort.Ints([]int{})
	// 字符串排序
	sort.Strings([]string{})
	// 自定义排序
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
}
