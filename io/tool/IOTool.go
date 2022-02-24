package tool

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*
	逐行读取文件数据
*/
func ReadFileDataLineByLine(path string) bool {
	open, err := os.Open(path)
	if err != nil {
		fmt.Println("file not exist")
	} else {
		defer func(open *os.File) {
			err := open.Close()
			if err != nil {

			}
		}(open)
		err := open.Chmod(0664)
		if err != nil {
			return false
		}
		fmt.Println("file exist,name=", open.Name())
	}

	// 读取文件
	inputReader := bufio.NewReader(open)
	for {
		//inputReader.ReadLine()
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return true
		}
	}
}

/*
读取整个文件
*/
func ReadTheEntireFile(path string) string {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	return string(buf)
}

/*
	拷贝文件到另外一个文件
*/
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer func(src *os.File) {
		err := src.Close()
		if err != nil {
		}
	}(src)

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
		}
	}(dst)

	return io.Copy(dst, src)
}
