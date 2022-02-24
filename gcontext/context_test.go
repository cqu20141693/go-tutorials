package gcontext

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

/*
context 源于 google，于 1.7 版本加入标准库，按照官方文档的说法，它是一个请求的全局上下文，携带了截止时间、手动取消等信号，并包含一个并发安全的 map 用于携带数据。
*/

/*
测试构造器
*/
func TestConstruct(t *testing.T) {
	// 不可以为 nil
	var ctx context.Context
	var cancelFunc context.CancelFunc
	ctx = context.Background()
	//如果拿捏不准是否需要一个全局的 context，可以使用下面这个函数构造
	ctx = context.TODO()
	ctx, cancelFunc = context.WithCancel(ctx)
	ctx, cancelFunc = context.WithTimeout(ctx, 2*time.Second)
	ctx, cancelFunc = context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	cancelFunc()
	ctx.Done()
}

/*
1. 请求链路传值
WithValue(parent Context, key, val interface{}) Context
Value(key interface{}) interface{}
*/
func TestLinkValue(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "k1", "v1")
	fmt.Println(ctx.Value("k1").(string))
}

/*
取消耗时操作，及时释放资源,网络交互场景,经常通过 SetReadDeadline、SetWriteDeadline、SetDeadline 进行超时取消
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

Done() <-chan struct{}
Err() error
*/
func TestCancelAndDeadline(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	timeoutWithContext(ctx)
	useContextCancel()

	useSelect()
}

func timeoutWithContext(ctx context.Context) {
	hctx, hcancel := context.WithTimeout(ctx, time.Second*4)
	defer hcancel()

	resp := make(chan struct{}, 1)
	// 处理逻辑
	go func() {
		// 处理耗时
		time.Sleep(time.Second * 10)
		resp <- struct{}{}
	}()

	// 超时机制
	select {
	//	case <-ctx.Done():
	//		fmt.Println("ctx timeout")
	//		fmt.Println(ctx.Err())
	case <-hctx.Done():
		fmt.Println("hctx timeout")
		fmt.Println(hctx.Err())
	case v := <-resp:
		fmt.Println("test2 function handle done")
		fmt.Printf("result: %v\n", v)
	}
	fmt.Println("test2 finish")
	return
}

func useContextCancel() {
	wg := new(sync.WaitGroup)
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go cancelWithContext(ctx, wg)
	time.Sleep(time.Second * 2)
	// 触发取消
	cancel()
	// 等待goroutine退出
	wg.Wait()
}

func cancelWithContext(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	respC := make(chan int)
	// 处理逻辑
	go func() {
		time.Sleep(time.Second * 5)
		respC <- 10
	}()
	// 取消机制
	select {
	case <-ctx.Done():
		fmt.Println("cancel")
		return errors.New("cancel")
	case r := <-respC:
		fmt.Println(r)
		return nil
	}
}

func useSelect() {
	respC := make(chan int)
	i := rand.Intn(10)
	// 处理逻辑
	go func() {
		time.Sleep(time.Duration(i) * time.Second)
		respC <- 10
		close(respC)
	}()

	// 超时逻辑
	select {
	case r := <-respC:
		fmt.Printf("Resp: %d\n", r)

	case <-time.After(time.Second * 2):
		fmt.Println("catch timeout")
	}
}
