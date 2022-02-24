package main

import (
	"bufio"
	"fmt"
	"github.com/cqu20141693/go-tutorials/io/tool"
	"os"
	"strings"
	"sync"
	"testing"
)

var inputReader *bufio.Reader
var input string
var err error

const (
	lineStart = "> "
)

func main() {

	//	testReader()
	TestScanner()
}

func testReader() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	for {
		input, err = inputReader.ReadString('\n')
		if err == nil {
			fmt.Printf("The input was: %s", input)
		}
		if strings.EqualFold(strings.TrimSpace(input), "q") {
			break
		}
	}
}

func TestScanner() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Print(lineStart)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() != "quit" {
				args := strings.Split(scanner.Text(), " ")
				fmt.Printf("exec cmd %v \n", args)
			} else {
				break
			}
			fmt.Print(lineStart)
		}
		wg.Done()
	}()
	wg.Wait()
}

func TestReaderFile(t *testing.T) {
	tool.ReadFileDataLineByLine("D:\\go-project\\go-demo\\src\\lib\\os\\os_example.go")
}
