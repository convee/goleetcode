package main

import "fmt"

//用数组也可以实现链表，比如定义一个数组 [5]Value，值类型为一个结构体 Value ：
func ArrayLink() {
	type Value struct {
		Data      string
		NextIndex int64
	}
	var array [5]Value          // 五个节点的数组
	array[0] = Value{"I", 3}    // 下一个节点的下标为3
	array[1] = Value{"Army", 4} // 下一个节点的下标为4
	array[2] = Value{"You", 1}  // 下一个节点的下标为1
	array[3] = Value{"Love", 2} // 下一个节点的下标为2
	array[4] = Value{"!", -1}   // -1表示没有下一个节点
	node := array[0]
	for {
		fmt.Println(node.Data)
		if node.NextIndex == -1 {
			break
		}
		node = array[node.NextIndex]
	}
}

//  打印出：
//  [0 0 0 0 0]
//  [8 9 7 0 0]
//  7

// 数组这一数据类型，是被编程语言高度抽象封装的结构，下标 会转换成 虚拟内存地址，然后操作系统会自动帮我们进行寻址，这个寻址过程是特别快的，所以往数组的某个下标取一个值和放一个值，时间复杂度都为 O(1)。
// 它是一种将 虚拟内存地址 和 数据元素 映射起来的内置语法结构，数据和数据之间是挨着，存放在一个连续的内存区域，每一个固定大小（8字节）的内存片段都有一个虚拟的地址编号。当然这个虚拟内存不是真正的内存，每个程序启动都会有一个虚拟内存空间来映射真正的内存，这是计算机组成的内容，和数据结构也有点关系
func main() {
	array := [5]int64{}
	fmt.Println(array)
	array[0] = 8
	array[1] = 9
	array[2] = 7
	fmt.Println(array)
	fmt.Println(array[2])
	ArrayLink()
}
