package main

import (
	"fmt"
	"sync"
)

// 集合

/*Collection

 */
func Collection() {

	maps()
}

type ConcurrentMap struct {
	sync.RWMutex
	m map[string]int
}

func (receiver *ConcurrentMap) get(key string) int {
	receiver.RWMutex.RLock()
	ret := receiver.m[key]
	receiver.RUnlock()
	return ret
}

func (receiver *ConcurrentMap) set(key string, value int) {
	receiver.RWMutex.Lock()
	receiver.m[key] = value
	receiver.Unlock()

}

/*maps
键值对：引用数据类型
初始化： {},make(map[key-type]val-type).
set:  name[key] = val
get:  name[key] ： 如果key不存在为零值

len: 长度
delete(map,key): 删除键

*/
func maps() {

	var m map[string]int
	fmt.Println("define:", m)
	// make
	m = make(map[string]int)

	// set
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	//get
	v1 := m["k1"]
	v1, _ = m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("default value: ", m["default"])

	// len
	fmt.Println("len:", len(m))

	// delete
	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// init
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	var concurrentMap = ConcurrentMap{m: make(map[string]int)}
	concurrentMap.set("age", 10)
	name := concurrentMap.get("age")
	fmt.Println("age", name)
}
