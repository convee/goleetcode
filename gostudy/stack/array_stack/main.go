package main

import (
	"fmt"
	"sync"
)

// 数组栈，后进先出
type ArrayStack struct {
	array []string   // 底层切片
	size  int        // 栈的元素数量
	lock  sync.Mutex // 为了并发安全使用的锁
}

// 入栈
func (stack *ArrayStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	// 放入切片中，后进的元素放在数组最后面
	stack.array = append(stack.array, v)
	// 栈中元素数量+1
	stack.size = stack.size + 1
}

func (stack *ArrayStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}
	// 栈顶元素
	v := stack.array[stack.size-1]
	// 切片收缩，但可能占用空间越来越大
	//stack.array = stack.array[0 : stack.size-1]
	// 创建新的数组，空间占用不会越来越大，但可能移动元素次数过多
	newArray := make([]string, stack.size-1, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newArray[i] = stack.array[i]
	}
	stack.array = newArray
	// 栈中元素数量-1
	stack.size = stack.size - 1
	return v
}

// 获取栈顶元素
func (stack *ArrayStack) Peek() string {
	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}
	// 栈顶元素值
	v := stack.array[stack.size-1]
	return v
}

// 栈大小
func (stack *ArrayStack) Size() int {
	return stack.size
}

// 栈是否为空
func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}
func main() {
	arrayStack := new(ArrayStack)
	arrayStack.Push("cat")
	arrayStack.Push("dog")
	arrayStack.Push("hen")
	fmt.Println("size:", arrayStack.Size())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("size:", arrayStack.Size())
	arrayStack.Push("drag")
	fmt.Println("pop:", arrayStack.Pop())
}
