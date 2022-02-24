package gsync

import (
	"fmt"
	"sync"
	"testing"
)

/*
sync.Once是一个简单而强大的原语，可确保一个函数仅执行一次。
*/

func TestOnce(t *testing.T) {

	for i := 0; i < 4; i++ {

		go onceTask(i)
	}
}

var once1 = &sync.Once{}

func onceTask(i int) {

	// func 中的代码只会被执行一次，尽管方法被多次调用
	once1.Do(func() {
		fmt.Printf("first %d\n", i)
	})
}
