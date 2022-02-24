package basic

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"testing"
)

/*
io 包里的 Readers 和 Writers 都是不带缓冲的，bufio 包里提供了对应的带缓冲的操作，在读写 UTF-8 编码的文本文件时它们尤其有用。

*/
func TestIOReader(t *testing.T) {
	var r io.Reader
	r = os.Stdin // see 12.1
	r = bufio.NewReader(r)

	r = new(bytes.Buffer)
	f, _ := os.Open("test.txt")
	r = bufio.NewReader(f)
}
