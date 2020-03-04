package main

import (
	"fmt"
	"log"
	"math/rand"
	"sorted/pkg/sort"
	"time"
)

var (
	logger *log.Logger
)

const size int = 100000

func main() {

	arr := make([]int, size)
	rand.Seed(time.Now().UnixNano())

	for i:=0; i<size; i++ {
		arr[i] = rand.Intn(size * 5) + 1
	}

	start := time.Now()
	sort.Sort(arr)
	end := time.Now()

	d := end.Sub(start)

	fmt.Printf("Loop size: %d - Run time: %v", size, d)


}

