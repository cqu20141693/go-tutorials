package concurrent

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

type Data struct {
	sum     int32
	counter int32
}

func TestAtomicInt(t *testing.T) {
	data := Data{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&data.sum, 1)
			data.counter += 1
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("data.sum=", atomic.LoadInt32(&data.sum)) // 1000
	fmt.Println("data.counter", data.counter)             // 1000
}

type CasData struct {
	sign int32
}

func TestCASBool(t *testing.T) {
	casData := CasData{}
	for i := 0; i < 1000; i++ {
		go func() {
			if atomic.CompareAndSwapInt32(&casData.sign, 0, -1) {
				fmt.Printf("goroutine i=%d cas success", i)
			}
		}()

	}

}
