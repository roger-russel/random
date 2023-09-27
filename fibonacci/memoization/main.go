package main

import (
	"fmt"
)

// function implementing memoization for Fibonaccin sequence number generation
// function take in the variables - nth term (type unsigned integer) and memo (type map) for caching
// function returns an unsignd integer value, representing the Fib Sequence value for the nth term
func memoized_fib(n uint, memo map[uint]uint) uint {
	// if nth term is already cached, return value
	if val, ok := memo[n]; ok {
		return val
	}
	
	// if nth term is 0 or 1, return the same value (0 or 1)
	if n == 0 || n == 1 {
		return n
	}
	
	// recursively generate the Fib Sequence values while caching results per iteration
	memo[n] = memoized_fib(n-1, memo) + memoized_fib(n-2, memo)
	return memo[n]
}

func main() {
	mem := make(map[uint]uint) // create a map (dict in python, object in JavaScript etc) to be used to cache results
	
	numSlice := []uint{10,15,20,25,30,40,50} // a slice (array/list in JS/Python) with nth terms to compute Fibonacci sequence numbers
	
	// undertake a loop, ranging over the slice items and passing each item to the memoized_fib function
	for _, j := range numSlice {
		fmt.Printf("Fib seq value for %v: %v\n", j, memoized_fib(uint(j), mem)) // print result for each slice entry
	}
		
}