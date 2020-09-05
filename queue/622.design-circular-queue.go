package queue

//622. 设计循环队列
//https://leetcode-cn.com/problems/design-circular-queue/

type MyCircularQueue struct {
	data        []int
	front, rear int // front标识出队元素的下标，read标识队的下标
	maxSize     int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		data:    make([]int, k+1, k+1),
		front:   0,
		rear:    0,
		maxSize: k + 1, // 多加一个用于给rear指针站位
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.rear] = value
	this.rear++
	if this.rear == this.maxSize {
		this.rear = 0
	}

	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.front++
	if this.front == this.maxSize {
		this.front = 0
	}

	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.front]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	lastIndex := this.rear - 1
	// 注意边界问题
	if lastIndex < 0 {
		lastIndex = this.maxSize - 1
	}
	return this.data[lastIndex]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.front == this.rear
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	next := this.rear + 1
	if next == this.maxSize {
		next = 0
	}
	return next == this.front
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
