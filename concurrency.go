package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go say("world")
	say("hello")

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 deadlock
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch)

	c2 := make(chan int, 10)
	go fibonacci(cap(c2), c2)

	for i := range c2 {
		fmt.Println(i)
	}

	c3 := make(chan int)
	quit := make(chan int)

	//lambda func
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c3)
		}
		quit <- 0
	}()
	fibonacci2(c3, quit)

	safeCounter := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go safeCounter.Increase("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(safeCounter.Value("somekey"))

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		//switch like system for channels
		//select statement is more focused on coordinating channel operations
		//and handling multiple channels concurrently.
		//If multiple cases are ready, one is chosen at random
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

type SafeCounter struct {
	mutex sync.Mutex
	v     map[string]int
}

func (safeCounter *SafeCounter) Increase(key string) {
	// Lock so only one goroutine at a time can access the map safeCounter.v.
	safeCounter.mutex.Lock()
	safeCounter.v[key]++
	safeCounter.mutex.Unlock()
}

func (safeCounter *SafeCounter) Value(key string) int {
	safeCounter.mutex.Lock()
	//we need to unlock just before it returns the value
	//so, we need to use defer here to achieve that
	defer safeCounter.mutex.Unlock()
	return safeCounter.v[key]
}
