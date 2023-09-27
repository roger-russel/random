package main

import "fmt"

type Fibonacci interface {
	Get(n int) int
}

type FiboImpl struct { }

func main() {
	fibo := &FiboImpl{}
	v := fibo.Get(3)
	fmt.Println(v)
}

// Time  O(n)
// Space O(1)
func (f *FiboImpl) Get(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	penultimoValue := 0
	previousValue := 1
	var currValue int

	for currN := 2; currN <= n; currN++ {
		currValue = penultimoValue + previousValue
		penultimoValue = previousValue
		previousValue = currValue
	}

	return currValue
}