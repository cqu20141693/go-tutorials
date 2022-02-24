package main

import (
	"log"
	"runtime"
)

/*
	包 runtime 中的函数 Caller() 提供了当前调用栈信息
*/
func main() {
	testCaller()
}

func testCaller() {

	// 打印当前执行的位置
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	// some code
	where()
	// some more code
	where()

	// 日志文件
	log.SetFlags(log.Llongfile)
	where1 := log.Println

	where1()
	where1()
}
