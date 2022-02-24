package chan_goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
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

/**
测试并发限制
*/
func TestConcurrencyLimit(t *testing.T) {
	var limit = make(chan int, 3)

	work := func(index int32) {
		time.Sleep(time.Second)
		fmt.Println("index=", index)
	}
	works := []func(i int32){work, work, work, work, work, work}
	index := int32(1)
	for _, w := range works {
		loadInt32 := atomic.LoadInt32(&index)
		go func() {
			limit <- 1
			w(loadInt32)
			<-limit
		}()
		atomic.AddInt32(&index, 1)
	}
	select {
	case <-time.After(3 * time.Second):
		return // 超时
	}
}

/*
Go 语言中不同 Goroutine 之间主要依靠管道进行通信和同步。要同时处理多个管道的发送或接收操作，我们需要使用 select 关键字（这个关键字和网络编程中的 select 函数的行为类似）。
当 select 有多个分支时，会随机选择一个可用的管道分支，如果没有可用的管道分支则选择 default 分支，否则会一直保存阻塞状态。
*/
func TestSelect(t *testing.T) {
	cancel := make(chan bool)
	defer close(cancel)
	for i := 0; i < 10; i++ {
		if i&8 == 1 {
			cancel <- true
		}
		// 通过chan 通知多个goroutine取消任务
		//其实我们可以通过 close 关闭一个管道来实现广播的效果，所有从关闭管道接收的操作均会收到一个零值和一个可选的失败标志
		go worker(cancel)
	}
	select {
	case <-time.After(1 * time.Second):
		return
	}
}

/*
测试goroutine 广播优雅退出
*/
func TestGoroutineBroadcast(t *testing.T) {
	cancel := make(chan bool)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go workerWithShutdown(&wg, cancel)
	}
	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
	fmt.Println("main shutdown")
}
func workerWithShutdown(wg *sync.WaitGroup, cancel chan bool) {
	defer wg.Done()
	for {
		var exit bool
		select {
		case c := <-cancel:
			fmt.Println("exit", c)
			exit = true
		default:
			fmt.Println("default")
			// 正常工作
		}
		if exit {
			break
		}
	}
	fmt.Println("goroutine shutdown")
}

func worker(cancel chan bool) {
	for {
		select {
		case c := <-cancel:
			fmt.Println("exit", c)
			return
		default:
			fmt.Println("default")
			// 正常工作
		}
	}
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

func TestDeadLock(t *testing.T) {
	select {}
}
