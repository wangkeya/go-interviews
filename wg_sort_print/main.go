package main

import (
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wgTest2()

}

func wgTest1() {
	// result: 3 1 2
	var wg sync.WaitGroup
	wg.Add(3)

	go func(n int) {
		println(n)
		wg.Done()
	}(1)
	go func(n int) {
		println(n)
		wg.Done()
	}(2)
	go func(n int) {
		println(n)
		wg.Done()
	}(3)
	wg.Wait()
}

func wgTest2() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(i)
		go func(n int) {
			println(n)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
