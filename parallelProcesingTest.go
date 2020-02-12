package main

import (
	"fmt"
	"sync"
	"time"
)

func forRoop() {
	for i := 0; i < 1000000000; i += 2 {
		i--
	}
}

func main() {
	start := time.Now()
	forRoop()
	forRoop()
	end := time.Now()
	fmt.Println(end.Sub(start).Seconds())

	// Method 1
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	start = time.Now()
	go func() {
		forRoop()
		ch1 <- true
	}()
	go func() {
		forRoop()
		ch2 <- true
	}()
	<-ch1
	<-ch2
	end = time.Now()
	fmt.Println(end.Sub(start).Seconds())

	// Method 2
	wg := sync.WaitGroup{}
	wg.Add(2)
	start = time.Now()
	go func() {
		forRoop()
		wg.Done()
	}()
	go func() {
		forRoop()
		wg.Done()
	}()
	wg.Wait()
	end = time.Now()
	fmt.Println(end.Sub(start).Seconds())
}
