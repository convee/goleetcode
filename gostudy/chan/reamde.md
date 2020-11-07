## chan 数据结构：
```
type hchan struct {
	qcount   uint           // 当前队列中剩余元素个数
	dataqsiz uint           // 环形队列长度，即可以存放的元素个数
	buf      unsafe.Pointer // 环形队列指针
	elemsize uint16         // 每个元素的大小
	closed   uint32         // 标识关闭状态
	elemtype *_type         // 元素类型
	sendx    uint           // 队列下标，指示元素写入时存放到队列中的位置
	recvx    uint           // 队列下标，指示元素从队列的该位置读出
	recvq    waitq          // 等待读消息的goroutine队列
	sendq    waitq          // 等待写消息的goroutine队列
	lock     mutex          // 互斥锁，chan不允许并发读写
}
```

## 环形队列
* chan内部实现了一个环形队列作为其缓冲区，队列的长度是创建chan时指定的
* make(chan int, 6)
* dataqsiz指示了队列长度为6，即可缓存6个元素；
* buf指向队列的内存，队列中还剩余两个元素；
* qcount表示队列中还有两个元素；
* sendx指示后续写入的数据存储的位置，取值[0, 6)；
* recvx指示从该位置读取数据, 取值[0, 6)；

## 等待队列
* 从channel读数据，如果channel缓冲区为空或者没有缓冲区，当前goroutine会被阻塞。向channel写数据，如果channel缓冲区已满或者没有缓冲区，当前goroutine会被阻塞。
* 被阻塞的goroutine将会挂在channel的等待队列中：
** 因读阻塞的goroutine会被向channel写入数据的goroutine唤醒；
** 因写阻塞的goroutine会被从channel读数据的goroutine唤醒；
* 注意，一般情况下recvq和sendq至少有一个为空。只有一个例外，那就是同一个goroutine使用select语句向channel一边写数据，一边读数据。

## 类型信息
* 一个channel只能传递一种类型的值，类型信息存储在hchan数据结构中。
* elemtype代表类型，用于数据传递过程中的赋值；
* elemsize代表类型大小，用于在buf中定位元素位置。

## 锁
* 一个channel同时仅允许被一个goroutine读写，为简单起见，本章后续部分说明读写过程时不再涉及加锁和解锁。

## 创建 chan
* 创建channel的过程实际上是初始化hchan结构。其中类型信息和缓冲区长度由make语句传入，buf的大小则与元素大小和缓冲区长度共同决定
```
func makechan(t *chantype, size int) *hchan {
	var c *hchan
	c = new(hchan)
	c.buf = malloc(元素类型大小 * size)
	c.elemsize = 元素类型大小
	c.elemtype = 元素类型
	c.dataqsiz = size
	return c
}
```
## 向channel写数据
* 如果等待接收队列recvq不为空，说明缓冲区中没有数据或者没有缓冲区，此时直接从recvq取出G,并把数据写入，最后把该G唤醒，结束发送过程；
* 如果缓冲区中有空余位置，将数据写入缓冲区，结束发送过程；
* 如果缓冲区中没有空余位置，将待发送数据写入G，将当前G加入sendq，进入睡眠，等待被读goroutine唤醒；

## 从channel读数据
* 如果等待发送队列sendq不为空，且没有缓冲区，直接从sendq中取出G，把G中数据读出，最后把G唤醒，结束读取过程；
* 如果等待发送队列sendq不为空，此时说明缓冲区已满，从缓冲区中首部读出数据，把G中数据写入缓冲区尾部，把G唤醒，结束读取过程；
* 如果缓冲区中有数据，则从缓冲区取出数据，结束读取过程；
* 将当前goroutine加入recvq，进入睡眠，等待被写goroutine唤醒

## 关闭channel
* 关闭channel时会把recvq中的G全部唤醒，本该写入G的数据位置为nil。把sendq中的G全部唤醒，但这些G会panic。
* 除此之外，panic出现的常见场景还有：
** 关闭值为nil的channel
** 关闭已经被关闭的channel
** 向已经关闭的channel写数据