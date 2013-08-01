package main

import "fmt"

func factorial1(n int) int {
	if n < 2 {
		return 1
	} else {
		return n * factorial1(n-1)
	}
}

func factorial2(n int) int {
	acc := 1
	for ; n > 1; n-- {
		acc *= n
	}
	return acc
}

func fact3Inner(acc int, x int) int {
	if x < 2 {
		return acc
	} else {
		return fact3Inner(acc*x, x-1)
	}
}

func factorial3(n int) int {
	return fact3Inner(1, n)
}

func main() {
	fmt.Printf("factorial1(10) = %d\n", factorial1(10))
	fmt.Printf("factorial2(10) = %d\n", factorial2(10))
	fmt.Printf("factorial3(10) = %d\n", factorial3(10))
}
