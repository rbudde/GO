package main

import (
	"fmt"
	"strconv"
	"time"
	"math/rand"
)

func main() {
    rand.Seed(42)
	ch := make(chan int, 10)
	go reader(ch)
	go writer(ch)
	time.Sleep(time.Second * 20)
}

func reader (in chan int) {
	for i := 1; i < 100; i++ {
		select {
			case x := <- in:
				fmt.Println("reader got " + strconv.Itoa(x))
			case <- time.After(time.Second * 1):
				fmt.Println("timeout")
		}
	}
}

func writer (out chan int) {
	x := 1
	for i := 1; i < 80; i++ {
		time.Sleep(time.Millisecond * time.Duration(rand.Int31n(3000)))
		x++
		out <- x
	}
}
