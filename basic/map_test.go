package basic

import (
	"fmt"
	"sort"
	"testing"
)

/*
	map 是引用类型，可以使用如下声明：
	var map1 map[keytype]valuetype
	未初始化的 map 的值是 nil,
	key 可以是任意可以用 == 或者！= 操作符比较的类型，比如 string、int、float,切片和数组不能作为key，结构体作为 key 可以提供 Key() 和 Hash() 方法
	map 传递给函数的代价很小：在 32 位机器上占 4 个字节，64 位机器上占 8 个字节
	初始化方式：
	{key1: val1, key2: val2} 的描述方法来初始化，就像数组和结构体一样
	var map1 = make(map[keytype]valuetype)

	map 增长到容量上限的时候，如果再增加新的 key-value 对，map 的大小会自动加 1。
	所以出于性能的考虑，对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明
	map的操作时非线程安全的

	delete(m, "two") ： map删除
	m["key"] : 获取map值
	m["key"]=1 : 设置map值

	map for for range : map[] 可以修改map数据，for range是对数据的拷贝
*/

func TestUseMap(t *testing.T) {

	// map 定义,没有初始化，不能使用
	var mapLit map[string]int
	//var mapCreated map[string]float32
	var mapAssigned map[string]int
	fmt.Println("map var ", mapLit, len(mapLit))
	if mapLit == nil {
		fmt.Println("mapList not  init ", mapLit)
	}
	mapLit = map[string]int{"one": 1, "two": 2}
	fmt.Println("map {} init ", mapLit, len(mapLit))
	mapCreated := make(map[string]float32)
	fmt.Println("map make init ", mapCreated, len(mapCreated))
	mapAssigned = make(map[string]int, 10)
	fmt.Println("map make init ", mapAssigned, len(mapAssigned))
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3
	fmt.Println("map [] set ", mapCreated, len(mapCreated))
	fmt.Println("map [] set ", mapAssigned, len(mapAssigned))

	testPresentAndDelete(mapAssigned)
	fmt.Println("testMapForRange")
	mapAssigned["age"] = 25
	testMapForRange(mapAssigned)
	fmt.Println("testSliceMap")
	testSliceMap()
	fmt.Println("testSortMap")
	testSortMap()
}

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

//map 默认是无序的，不管是按照 key 还是按照 value 默认都不排序
//为 map 排序，需要将 key（或者 value）拷贝到一个切片，再对切片排序，然后可以使用切片的 for-range 方法打印出所有的 key 和 value
func testSortMap() {
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}
}

func testSliceMap() {
	// Version A:
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	// Version B: NOT GOOD!
	items2 := make([]map[int]int, 5)
	//for range 是值copy
	for _, item := range items2 {
		item = make(map[int]int, 1) // item is only a copy of the slice element.
		item[1] = 2                 // This 'item' will be lost on the next iteration.
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)

	// test map slice
	m := make(map[string][]int, 1)
	m["employee"] = []int{1, 2}
	// range slice 是引用
	for _, ints := range m {
		ints[1] = 3
	}
	fmt.Println("map slice modify ", m["employee"])

	//
}

func testMapForRange(m map[string]int) {
	for key := range m {
		fmt.Printf("key=%s,value=%d", key, m[key])
	}
	fmt.Println()

	for key, value := range m {
		fmt.Printf("key=%s,value=%d", key, value)
	}
	fmt.Println()
}

func testPresentAndDelete(m map[string]int) {
	if _, ok := m["key"]; ok {
		fmt.Println("map key is exist", m["key"])
	}
	if _, ok := m["two"]; ok {
		fmt.Println("map two is exist", m["two"])
	}
	// delete 方法没有返回值
	delete(m, "two")
	fmt.Println("map delete key two")
	if _, ok := m["two"]; !ok {
		fmt.Println("map two is not exist", m["two"])
	}
}

type Foo map[string]string

func TestMap(t *testing.T) {

	// 使用 new(Foo) 返回的是一个指向 nil 的指针，它尚未被分配内存
	f6 := new(Foo)
	//(*f6)["name"]="wq"
	fmt.Println("map new ", f6)
	// 对指针重新赋值后操作
	*f6 = make(map[string]string, 0)
	(*f6)["company"] = "cc"
	fmt.Println("map new init ", f6)
	// 使用make
	foo := make(Foo)
	foo["name"] = "gowb"
	foo["age"] = "23"
	fmt.Println("map make ", foo)
}
