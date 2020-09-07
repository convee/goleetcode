package slice

func main() {
	//创建队列
	queue := make([]int, 0)
	//enqueue 入队
	queue = append(queue, 10)
	//dequeue 出队
	v := queue[0]
	queue := queue[1:]
	//长度为空
	len(queue) == 0

	//注意点
	//参数传递，只能修改，不能新增或者删除原始数据
	//默认 s=s[0:len(s)]，取下限不取上限，数学表示为：[)
}
