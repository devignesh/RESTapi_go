package main

import (
	"fmt"
	"sync"
)

type element struct {
	p int
	c int
}

func main() {
	var num int
	// fmt.Scanf("Enter how many numbers you want", &num)
	num = 10
	ch := make(chan element, 100)
	wg := new(sync.WaitGroup)
	var ele = element{
		p: 0,
		c: 1,
	}
	ch <- ele
	for i := 0; i < num; i++ {
		wg.Add(1)
		el := <-ch
		go fib(el, ch)
	}
	wg.Wait()
	close(ch)

}

func fib(ele element, ch chan element) {
	// ele := <-ch
	n := ele.c + ele.p
	ele.p = ele.c
	ele.c = n
	fmt.Println(ele.c)
	ch <- ele
}
