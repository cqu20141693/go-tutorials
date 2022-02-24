package basic

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

/*
	// defer 和追踪 ，参考defer_test.go
	关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数
	类似于面向对象编程语言 Java 中finally 语句块
	当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
	关闭文件流：defer file.Close()
	解锁一个加锁的资源：mu.Lock()  defer mu.Unlock()
	打印最终报告： printFooter()
	关闭数据库链接： defer disconnectFromDB()

*/
func TestDefer(t *testing.T) {
	f, err := os.Open("./defer_test.go")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println("file close failed")
		}
	}(f)
}

/*
	return xxx这一条语句并不是一条原子指令!
	先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中
	但是defer函数对返回值进行操作，会改变返回值
*/
func TestDeferReturnValue(t *testing.T) {

	log.Println("f=", f())
	log.Println("f2=", f2())
	log.Println("f5=", f5())
	log.Println("f3=", f3())
	log.Println("f4=", f4())
	log.Println("f6=", f6())
	fmt.Println()
}

func TestRecordParamAndReturn(t *testing.T) {
	func1("Go")
}

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}

func f6() (r int) {
	// 这里是对参数的操作
	defer func(r *int) {
		*r = (*r) + 5
	}(&r)
	return 1
}
func f4() (result int) {
	result = 0     //return语句不是一条原子调用，return xxx其实是赋值＋ret指令
	defer func() { //defer被插入到return之后执行，也就是赋返回值和ret指令之间
		result++
	}()
	return
}
func f3() (r int) {
	// 这里是对参数的操作
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
func f5() (r int) {
	t := 5
	defer func() {
		r = r + 5
	}()
	// 先将t赋值r后再执行的defer
	return t
}
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	// 先将t赋值r后再执行的defer
	return t
}
func f() (result int) {
	defer func() {
		result++
	}()
	return
}
