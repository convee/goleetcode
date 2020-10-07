package main

import (
	"fmt"
	"unsafe"
)

//因为数组大小是固定的，当数据元素特别多时，固定的数组无法储存这么多的值，所以可变长数组出现了，这也是一种数据结构。在 Golang语言中，可变长数组被内置在语言里面：切片 slice
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

//1.指向底层数组的指针。( Golang 语言是没有操作原始内存的指针的，所以 unsafe 包提供相关的对内存指针的操作，一般情况下非专业人员勿用)
//2.切片的真正长度，也就是实际元素占用的大小。
//3.切片的容量，底层固定数组的长度。

// 每次可以初始化一个固定容量的切片，切片内部维护一个固定大小的数组。当 append 新元素时，固定大小的数组不够时会自动扩容，如：
func main() {
	// 创建一个容量为2的切片
	array := make([]int, 0, 2)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	// 虽然 append 但是没有赋予原来的变量 array
	_ = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	_ = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	_ = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	fmt.Println("-------")
	// 赋予回原来的变量
	array = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	array = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	array = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	array = append(array, 1, 1, 1, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	array = append(array, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
}
