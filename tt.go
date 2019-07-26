package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 以前理解错误了，都这样写的
func selectSort1(arrays []int) {
	l := len(arrays)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if arrays[i] > arrays[j] {
				arrays[i], arrays[j] = arrays[j], arrays[i]
			}
		}
	}
}

// 正确的选择排序应该这样
func selectSort2(arrays []int) {
	l := len(arrays)
	for i := 0; i < l-1; i++ {
		index := i
		for j := i + 1; j < l; j++ {
			if arrays[index] > arrays[j] {
				index = j
			}
		}
		arrays[i], arrays[index] = arrays[index], arrays[i]
	}
}

// 冒泡排序
func bubbleSort(arrays []int) {
	l := len(arrays)
	for i := 0; i < l; i++ {
		for j := 0; j < l-1-i; j++ {
			if arrays[j] > arrays[j+1] {
				arrays[j], arrays[j+1] = arrays[j+1], arrays[j]
			}
		}
	}
}

func copyArray(arrays []int) []int {
	result := make([]int, len(arrays))
	copy(result, arrays)
	return result
}

func newRandArrays(n int) []int {
	var arrays []int
	rand.Seed(time.Now().Unix())
	for i := n; i > 0; i-- {
		arrays = append(arrays, rand.Intn(100))
	}
	return arrays
}

func main() {
	arrays := newRandArrays(10000)
	fmt.Printf("array-src: %v\n", arrays)
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		array := copyArray(arrays)
		start := time.Now()
		selectSort1(array)
		useTime := time.Now().Sub(start)
		fmt.Printf("sort1: %v\nTime=%v\n", array, useTime)
	}()

	go func() {
		defer wg.Done()
		array := copyArray(arrays)
		start := time.Now()
		selectSort2(array)
		useTime := time.Now().Sub(start)
		fmt.Printf("sort2: %v\nTime=%v\n", array, useTime)
	}()

	go func() {
		defer wg.Done()
		array := copyArray(arrays)
		start := time.Now()
		bubbleSort(array)
		useTime := time.Now().Sub(start)
		fmt.Printf("bubbleSort: %v\nTime=%v\n", array, useTime)
	}()

	wg.Wait()
	fmt.Println("Run End !")
}
