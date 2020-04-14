package main

import (
	"fmt"
	"testing"
	"time"
)

func Chann(ch chan int, stopCh chan bool) {
	var _ int
	_ = 10
	for j := 0; j < 10; j++ {
		ch <- j
		time.Sleep(time.Second)
	}
	stopCh <- true
}

func TestSelect(t *testing.T) {

	ch := make(chan int)
	c := 0
	stopCh := make(chan bool)

	go Chann(ch, stopCh)

	for {
		select {
		case c = <-ch:
			fmt.Println("Recvice", c)
			fmt.Println("channel")
		case s := <-ch:
			fmt.Println("Receive", s)
		case _ = <-stopCh:
			goto end
		}
	}
end:
}
