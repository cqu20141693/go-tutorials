package gsync

import (
	"fmt"
	"sync"
	"testing"
)

var Total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		Total.Lock()
		Total.value += i
		Total.Unlock()
	}
}
func TestMux(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	//在 worker 的循环中，为了保证 Total.value += i 的原子性，我们通过 sync.Mutex 加锁和解锁来保证该语句在同一时刻只被一个线程访问
	go worker(&wg)
	go worker(&wg)
	wg.Wait()

	fmt.Println(Total.value)
}
