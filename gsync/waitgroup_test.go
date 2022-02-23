package gsync

import (
	"fmt"
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	// 开 N 个后台打印线程
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			fmt.Println("你好, 世界")
			wg.Done()
		}()
	}

	// 等待 N 个后台线程完成
	wg.Wait()
}
