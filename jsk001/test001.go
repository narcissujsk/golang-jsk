package jsk001

import "fmt"

var b bool = false

func test(bool2 bool) bool {
	b = bool2
	return b
}

var d = b

func maintest001() {
	var a *int
	c := true
	println(c)
	println(*&a)
	println(d)
	for i := 0; i < 99; i++ {
		println(i)
	}
	var nums []int = []int{1, 2, 3}
	for i, x := range nums {
		fmt.Printf("%d %d \n", i, x)
	}
}
