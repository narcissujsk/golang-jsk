package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 59999; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main1() {
	go say("world")
	say("hello")
}
