package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

/*
hash 包：实现了 adler32、crc32、crc64 和 fnv 校验；
*/

func main() {
	hasher := sha1.New()
	// 写入数据到hasher
	io.WriteString(hasher, "test")
	b := []byte{}
	// hash %v
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	// hash %d
	fmt.Printf("Result: %d\n", hasher.Sum(b))
	// 重置hash
	hasher.Reset()
	data := []byte("We shall overcome!")
	// 写数据
	n, err := hasher.Write(data)
	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v / %v", n, err)
	}
	// hash
	checksum := hasher.Sum(b)
	fmt.Printf("Result: %x\n", checksum)

	Md5()
}

func Md5() {
	hasher := md5.New()
	b := []byte{}
	io.WriteString(hasher, "test")
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %x\n", hasher.Sum(b))
}
