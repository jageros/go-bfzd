package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var value int

func writes(i int, wg *sync.WaitGroup, rwM *sync.RWMutex) {
	defer wg.Done()

	for {
		rand.Seed(time.Now().Unix())
		num := rand.Intn(100)
		rwM.Lock()
		value = num
		fmt.Printf("writes goroutine:%d, Write num=%d\n", i, num)
		rwM.Unlock()
		time.Sleep(time.Second * time.Duration(i+1))
	}
}

func readRLocker(i int, wg *sync.WaitGroup, rwM *sync.RWMutex) {
	defer wg.Done()
	w := rwM.RLocker()

	for {
		w.Lock()
		num := value
		fmt.Printf("readRLocker goroutine:%d, Reading num=%d\n", i, num)
		time.Sleep(time.Second * 2)
		w.Unlock()
		time.Sleep(time.Second * 1)
	}
}

func read2(i int, wg *sync.WaitGroup, rwM *sync.RWMutex) {
	defer wg.Done()

	for {
		rwM.RLock()
		num := value
		fmt.Printf("read2 goroutine:%d, Reading num=%d\n", i, num)
		time.Sleep(time.Second * 2)
		rwM.RUnlock()
		time.Sleep(time.Second * 1)
	}
}

func main() {
	rwMutex := &sync.RWMutex{}
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go writes(i, wg, rwMutex)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go readRLocker(i, wg, rwMutex)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read2(i, wg, rwMutex)
	}

	wg.Wait()
}
