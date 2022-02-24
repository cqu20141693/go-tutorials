package gsync

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sync"
	"testing"
)

type Connection struct {
	address  string
	user     string
	password string
}

func NewConnection(address string, user string, password string) *Connection {
	return &Connection{address: address, user: user, password: password}
}

/*
一组安全的对象池
使用func (p *Pool) Put(x interface{}) 将对象回收
使用func (p *Pool) Get() interface{} 获取对象，可能存在对象池中不存在对象,可以利用New方法进行无线创建
*/
func TestPool(t *testing.T) {
	pool := &sync.Pool{}

	pool.Put(NewConnection("127.0.0.1", "root", "cc"))
	pool.Put(NewConnection("127.0.0.1", "root", "ccc"))
	pool.Put(NewConnection("127.0.0.1", "root", "cccc"))

	connection := pool.Get().(*Connection)
	fmt.Printf("%s\n", connection.password)
	connection = pool.Get().(*Connection)
	fmt.Printf("%s\n", connection.password)
	connection = pool.Get().(*Connection)
	fmt.Printf("%s\n", connection.password)

	get := pool.Get()
	if get == nil {
		fmt.Println("current pool object is nil")
	}
	usePoolBuffer()

}

func usePoolBuffer() {
	pool := &sync.Pool{
		New: func() interface{} {
			return bytes.NewBufferString("pool buffer")
		},
	}

	usePool(pool)
	usePool(pool)
	get1 := pool.Get()
	get2 := pool.Get()
	fmt.Println(get1, get2)
}

func usePool(pool *sync.Pool) {
	buf := pool.Get().(*bytes.Buffer)
	fmt.Println(buf)
	defer pool.Put(buf)

	// Reset 缓存区，不然会连接上次调用时保存在缓存区里的字符串foo
	buf.Reset()

	buf.WriteString("foo")
	ioutil.WriteFile("filename", buf.Bytes(), 0644)
}
