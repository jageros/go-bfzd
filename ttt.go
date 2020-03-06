package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	path := RandNFromList([]interface{}{2, 3, 4, 5, 6, 7, 7, 8, 9, 0}, 5)
	fmt.Printf("path=%v\n", path)
	path = RandNFromList([]interface{}{2, 3, 4, 5, 6, 7, 7, 8, 9, 0}, 5)
	fmt.Printf("path=%v\n", path)
}

func RandNFromList(list []interface{}, n int) []interface{} {
	if n <= 0 {
		return []interface{}{}
	}
	listLen := len(list)
	if n >= listLen {
		return list
	}

	var result []interface{}
	index := rand.Perm(listLen)
	for i := 0; i < n; i++ {
		result = append(result, list[index[i]])
	}
	return result
}
