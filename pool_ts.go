package main

import (
	"fmt"
	"sync"
	"time"
)

type st struct {
	i       int
	isPay   bool
	i2count map[int]int
}

func (s *st) init(i int, isPay bool, i2c map[int]int) {
	s.i = i
	s.isPay = isPay
	s.i2count = i2c
}

func main() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated += 1
			st := &st{}
			return st
		},
	}

	//for i := 0; i < 5; i++ {
	//	calcPool.Put(calcPool.New())
	//}

	const numWorkers = 1000
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := numWorkers; i > 0; i-- {
		go func(j int) {
			defer wg.Done()
			time.Sleep(time.Millisecond * time.Duration(j))
			fn := calcPool.Get().(*st)
			fn.init(j, true, map[int]int{j: j * 100})
			fmt.Printf("%d--st: %v\n", j, fn)
			defer calcPool.Put(fn)
		}(i)
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.\n", numCalcsCreated)
}
