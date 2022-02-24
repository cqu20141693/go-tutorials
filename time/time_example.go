package main

import (
	"fmt"
	"reflect"
	"time"
)

/*
	时间日期工具
	Time 时间结构体

	时间日期功能func： Day(),Minute(),Second(),Month(),Year()

	Duration 类型表示两个连续时刻所相差的纳秒数，类型为 int64。
	Location 类型映射某个时区的时间，UTC 表示通用协调世界时间

	时间格式化函数：
	func (t Time) Format(layout string) string
	定义的格式，如：time.ANSIC 或 time.RFC822

	time.Ticker 结构体，对象以指定的时间间隔重复的向通道 C 发送时间值
	时间间隔的单位是 ns（纳秒，int64），在工厂函数 time.NewTicker 中以 Duration 类型的参数传入：func Newticker(dur) *Ticker


	time 函数
	time.Now()
	time.Tick(1e8)
	time.After(1e8)
	time.Sleep(1e8)


*/
var week time.Duration

func main() {
	t := time.Now()
	second := t.Unix()
	milli := t.UnixMilli()
	fmt.Printf("second=%d,millil=%d,mirco=%d,nano=%d \n", second, milli, t.UnixMicro(), t.UnixNano())
	fmt.Println("timestamp convert time ", time.Unix(second, 0), time.UnixMilli(milli))

	fmt.Println("t type=", reflect.TypeOf(t))
	fmt.Println(t) // e.g. Wed Dec 21 09:52:14 +0100 RST 2011
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	// 21.12.2011
	t = time.Now().UTC()
	fmt.Println(t)          // Wed Dec 21 08:52:14 +0000 UTC 2011
	fmt.Println(time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011
	// calculating times:
	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(week)
	fmt.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011
	// formatting times:
	fmt.Println(t.Format(time.RFC822))         // 21 Dec 11 0852 UTC
	fmt.Println(t.Format(time.ANSIC))          // Wed Dec 21 08:56:34 2011
	fmt.Println(t.Format("21 Dec 2011 08:52")) // 21 Dec 2011 08:52
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
	// Wed Dec 21 08:52:14 +0000 UTC 2011 => 20111221

	testTick()

	// 计算时间
	calculateTime()

}

func calculateTime() {
	start := time.Now()
	time.Sleep(3e9)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func testTick() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom.")
			return
		default:
			fmt.Println("....")
			time.Sleep(5e7)
		}
	}

}
