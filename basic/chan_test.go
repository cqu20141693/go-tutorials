package basic

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	event := make(chan int, 1)

	go func(e chan int) {
		select {
		case c, ok := <-event:
			fmt.Println("a receive event", c, ok)
		}
	}(event)
	go func(e chan int) {
		select {
		case c, ok := <-event:
			fmt.Println("b receive event", c, ok)
		}
	}(event)
	go func(e chan int) {
		select {
		case c, ok := <-event:
			fmt.Println("c receive event", c, ok)
		}
	}(event)
	go func(e chan int) {
		e <- 1
		close(e)
	}(event)

	select {
	case c, ok := <-event:
		fmt.Println("main receive event", c, ok)
	}
	time.Sleep(1 * time.Millisecond)
}
