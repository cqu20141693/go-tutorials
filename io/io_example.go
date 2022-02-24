package main

/*
io 包里的 Readers 和 Writers 都是不带缓冲的，bufio 包里提供了对应的带缓冲的操作，在读写 UTF-8 编码的文本文件时它们尤其有用。

`bufio
 os
 fmt
 compress

defer 关键字对于在函数结束时关闭打开的文件非常有用
*/

func main() {

	// 将整个文件的内容读到一个字符串里，使用 io/ioutil 包里的 ioutil.ReadFile() 方法，
	//该方法第一个返回值的类型是 []byte，里面存放读取到的内容，第二个返回值是错误，如果没有错误发生，第二个返回值为 nil
}
