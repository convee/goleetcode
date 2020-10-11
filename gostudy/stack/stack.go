package main

func main() {
	//创建栈
	stack := make([]int, 0)
	//push 押入
	stack = append(stack, 10)
	//pop 弹出
	v := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	//检查栈空
	len(stack) == 0
}
