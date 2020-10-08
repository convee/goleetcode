package main

import "fmt"

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

//获取下一个节点
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

//获取上一个节点
func (r *Ring) Prev() *Ring {
	if r.prev == nil {
		return r.init()
	}
	return r.prev
}

//获取第n个节点
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:

		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r

}

//添加节点
//往节点A，链接一个节点，并且返回之前节点A的后驱节点
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

//删除节点
func (r *Ring) Unlink(n int) *Ring {
	if n < 0 {
		return nil

	}
	return r.Link(r.Move(n + 1))
}

// 查看循环链表长度
func (r *Ring) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}
func linkNewTest() {
	// 第一个节点
	r := &Ring{Value: -1}
	// 链接新的五个节点
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})
	node := r
	for {
		// 打印节点值
		fmt.Println(node.Value)
		// 移到下一个节点
		node = node.Next()
		//  如果节点回到了起点，结束
		if node == r {
			return
		}
	}
}
func main() {
	linkNewTest()
}
