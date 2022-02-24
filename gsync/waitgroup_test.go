package gsync

import (
	"fmt"
	"sync"
	"testing"
)

/*
func (wg *WaitGroup) Add(delta int)
func (wg *WaitGroup) Done()
func (wg *WaitGroup) Wait()
*/
func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	// 开 N 个后台打印线程
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			fmt.Println("你好, 世界")

		}()
	}

	// 等待 N 个后台线程完成
	wg.Wait()
}
