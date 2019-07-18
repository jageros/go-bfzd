package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

func producer(wg *sync.WaitGroup, l sync.Locker) {
	defer wg.Done()
	for i := 5; i > 0; i-- {
		l.Lock()
		l.Unlock()
		time.Sleep(1)
	}
}

func observer(wg *sync.WaitGroup, l sync.Locker) {
	defer wg.Done()
	l.Lock()
	defer l.Unlock()
	time.Sleep(1)
}

func test(count int, mutex, rwMutex sync.Locker) time.Duration {
	var wg sync.WaitGroup
	wg.Add(count+1)
	bTim := time.Now()
	go producer(&wg, mutex)

	for i := count; i > 0; i-- {
		go observer(&wg, rwMutex)
	}
	wg.Wait()
	return time.Since(bTim)
}

func main() {
	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutext\ttMutex\n")

	for i := 0; i < 22; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
			)
	}
}
