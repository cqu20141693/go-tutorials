package chan_goroutine

import (
	"sync"
	"testing"
)

/*
在 Go 语言中，同一个 Goroutine 线程内部，顺序一致性内存模型是得到保证的。但是不同的 Goroutine 之间，并不满足顺序一致性内存模型，需要通过明确定义的同步事件来作为同步的参考。
如果两个事件不可排序，那么就说这两个事件是并发的。为了最大化并行，Go 语言的编译器和处理器在不影响上述规定的前提下可能会对执行语句重新排序（CPU 也会对一些指令进行乱序执行）。
*/
var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func TestName(t *testing.T) {

}

func TestGoroutine(t *testing.T) {
	go setup()
	for !done {
	}
	print(a)
}

/*
Channel 通信是在 Goroutine 之间进行同步的主要方法。在无缓存的 Channel 上的每一次发送操作都有与其对应的接收操作相配对，发送和接收操作通常发生在不同的
Goroutine 上（在同一个 Goroutine 上执行两个操作很容易导致死锁）。无缓存的 Channel 上的发送操作总在对应的接收操作完成前发生.
严谨的并发也应该是可以静态推导出结果的：根据线程内顺序一致性，结合 Channel 或 sync 同步事件的可排序性来推导，最终完成各个线程各段代码的偏序关系排序。
如果两个事件无法根据此规则来排序，那么它们就是并发的，也就是执行先后顺序不可靠的。
*/
func TestCSP(t *testing.T) {
	useChan()
	useMutex()
}

func useMutex() {
	var mu sync.Mutex

	mu.Lock()
	go func() {
		println("你好, 世界")
		mu.Unlock()
	}()

	mu.Lock()
}

func useChan() {
	done := make(chan int)
	var msg string
	go func() {
		msg = "你好, 世界"
		done <- 1
	}()
	<-done
	println(msg)
}
