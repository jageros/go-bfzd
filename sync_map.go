package main

import (
	"fmt"
	"sync"
)

func main() {
	sMap := sync.Map{}

	fmt.Println("=====================Store=======================")

	fmt.Println("store: key=1, value=jay")
	sMap.Store(1, "jay")
	fmt.Println("store: key=2, value=lhj")
	sMap.Store(2, "lhj")
	fmt.Println("store: key=2, value=lhj-hhh")
	sMap.Store(2, "lhj-hhh")


	fmt.Println("=====================LoadOrStore=======================")

	result, succeed := sMap.LoadOrStore(1, "jays")
	fmt.Printf("key=1 store: value=jays; result: load_succeed=%t, value=%v\n", succeed, result)

	result, succeed = sMap.LoadOrStore(2, "lhj-s")
	fmt.Printf("key=2 store: value=lhj-s; result: load_succeed=%t, value=%v\n", succeed, result)

	result, succeed = sMap.LoadOrStore(3, "jay-lhj")
	fmt.Printf("key=3 store: value=jay-lhj; result: load_succeed=%t, value=%v\n", succeed, result)



	fmt.Println("=====================Load key:1~4=======================")

	if val, ok := sMap.Load(1); ok {
		fmt.Printf("key=1, value=%v\n", val)
	}else {
		fmt.Printf("key=1, load failed\n")
	}

	if val, ok := sMap.Load(2); ok {
		fmt.Printf("key=2, value=%v\n", val)
	}else {
		fmt.Printf("key=2, load failed\n")
	}

	if val, ok := sMap.Load(3); ok {
		fmt.Printf("key=3, value=%v\n", val)
	}else {
		fmt.Printf("key=3, load failed\n")
	}

	if val, ok := sMap.Load(4); ok {
		fmt.Printf("key=4, value=%v\n", val)
	}else {
		fmt.Printf("key=4, load failed\n")
	}

	fmt.Println("======================Delete======================")

	fmt.Println("delete key=1")
	sMap.Delete(1)

	fmt.Println("=====================Range=======================")
	sMap.Range(func(key, value interface{}) bool {
		fmt.Printf("key=%v value=%v\n", key, value)
		return true
	})

}
