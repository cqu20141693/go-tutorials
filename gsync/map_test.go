package gsync

import (
	"fmt"
	"sync"
	"testing"
)

/*
使用func (m *Map) Store(key, value interface{})添加元素。
使用func (m *Map) Load(key interface{}) (value interface{}, ok bool) 检索元素。
使用func (m *Map) Delete(key interface{}) 删除元素。
使用func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)检索或添加之前不存在的元素。如果键之前在map中存在，则返回的布尔值为true。
使用func (m *Map) Range(f func(key, value interface{}) bool)遍历元素。遍历到条件不满足
*/
func TestMap(t *testing.T) {
	m := &sync.Map{}

	// 添加元素
	m.Store(1, "one")
	m.Store(2, "two")

	// 获取元素1
	value, contains := m.Load(1)
	if contains {
		fmt.Printf("%s\n", value.(string))
	}

	// 返回已存value，否则把指定的键值存储到map中
	value, loaded := m.LoadOrStore(3, "three")
	if !loaded {
		fmt.Printf("%s\n", value.(string))
	}

	m.Delete(3)

	// 迭代所有元素
	m.Range(func(key, value interface{}) bool {
		v := key.(int)
		fmt.Printf("%d: %s\n", v, value.(string))
		return true
	})
	// 遍历到条件不满足
	m.Range(func(key, value interface{}) bool {
		v := key.(int)
		fmt.Printf("%d: %s\n", v, value.(string))
		return v < 1
	})
}
