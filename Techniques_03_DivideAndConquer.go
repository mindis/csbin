package main

import (
	"fmt"
	"math/rand"
	"time"
)

func search(list []int, value, start, end int) int {
	if end-start < 1 {
		return -1
	}

	middle := start + ((end - start) / 2)
	if list[middle] == value {
		return middle
	} else if list[middle] < value {
		return search(list, value, middle+1, end)
	} else {
		return search(list, value, start, middle)
	}
}

func main() {
	fmt.Println("Exercise 1: 1 million random searches")
	for length := 10; length <= 1000000; length *= 10 {
		fmt.Printf("List Length % 8d: ", length)

		list := make([]int, length)
		for i := 0; i < length; i++ {
			list[i] = 2 * i
		}

		before := time.Now()
		for j := 0; j < 1000000; j++ {
			value := rand.Intn(len(list) * 2)
			index := search(list, value, 0, len(list))
			if index != -1 && list[index] != value {
				fmt.Printf("list[%d] == %d, != %d\n", index, list[index], value)
			}
		}
		millis := int((time.Now().Sub(before)) / time.Millisecond)
		fmt.Printf("% 4d ms\n", millis)
	}
}
