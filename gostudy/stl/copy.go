package stl

func main() {
	// 删除a[i]，可以用 copy 将i+1到末尾的值覆盖到i,然后末尾-1
	copy(a[i:], a[i+1:])
	a = a[:len(a)-1]

	// make创建长度，则通过索引赋值
	a := make([]int, n)
	a[n] = x
	// make长度为0，则通过append()赋值
	a := make([]int, 0)
	a = append(a, x)
}
