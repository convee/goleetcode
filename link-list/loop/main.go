package main

//循环链表
type Ring struct {
	next, prev *Ring       //前驱和后驱节点
	Value      interface{} //数据
}

//初始化一个空的循环链表
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

//创建一个指定大小 N 的循环链表，值全为空：
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}
func main() {
	r := new(Ring)
	r.init()
}
