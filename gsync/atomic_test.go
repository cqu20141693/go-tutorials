package gsync

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

/*
sync/atomic 包对基本的数值类型及复杂对象的读写都提供了原子操作的支持。bool,int,float,unit
atomic.Value 原子对象提供了 Load 和 Store 两个原子方法，分别用于加载和保存数据，返回值和参数都是 interface{} 类型，因此可以用于任意的自定义复杂类型。
*/
var total uint64

func TestAtomic(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	go atomicWorker(&wg)
	go atomicWorker(&wg)
	wg.Wait()
}

/*
func AddInt32(addr *int32, delta int32) (new int32)
func LoadInt32(addr *int32) (val int32)
*/
func TestAtomicAddAndLoad(t *testing.T) {
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

func TestAtomicValue(t *testing.T) {
	// 初始化配置信息
	config.Store(loadConfig())
	go func() {
		for {
			time.Sleep(time.Second)
			config.Store(loadConfig())
		}
	}()
	// 用于处理请求的工作者线程始终采用最新的配置信息
	for i := 0; i < 10; i++ {
		go func() {
			for i, r := range requests() {
				fmt.Println(config.Load(), i, r)
				// ...
			}
		}()
	}

	time.Sleep(time.Second)
}

/*
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
*/
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

func atomicWorker(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i <= 100; i++ {
		//atomic.AddUint64 函数调用保证了 total 的读取、更新和保存是一个原子操作
		atomic.AddUint64(&total, i)
	}
}

//原子操作配合互斥锁可以实现非常高效的单件模式。互斥锁的代价比普通整数的原子读写高很多，在性能敏感的地方可以增加一个数字型的标志位，通过原子检测标志位状态降低互斥锁的使用次数来提高性能。
type singleton struct{}

var (
	instance     *singleton
	onceInstance *singleton
	initialized  uint32
	mu           sync.Mutex
	once         sync.Once
)

// 单例
func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		fmt.Println("singleton instance")
		instance = &singleton{}
	}
	return instance
}
func InstanceByOnce() *singleton {
	once.Do(func() {
		fmt.Println("singleton by onceo")
		onceInstance = &singleton{}
	})
	return onceInstance
}
func TestAtomicSingleton(t *testing.T) {

	go func() {
		fmt.Println(Instance())
	}()
	go func() {
		fmt.Println(Instance())
	}()

	go func() {
		fmt.Println(InstanceByOnce())
	}()
	go func() {
		fmt.Println(InstanceByOnce())
	}()
	time.Sleep(time.Second)
}

var config atomic.Value // 保存当前配置信息

type AppConfig struct {
	maxConn int
}

func requests() []string {
	return []string{"get,post,delete,put,head,option"}
}

func loadConfig() interface{} {
	fmt.Println("load config from remote database")
	return &AppConfig{10}
}

type Data struct {
	sum     int32
	counter int32
}

type CasData struct {
	sign int32
}
