package main

import "fmt"

func main() {
	nums := []uint{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	for _, n := range nums {
		f := Fibonacci(n)
		fmt.Println("Fibonacci sequencie value for",n,":", f)
	}
}

func Fibonacci(n uint) uint {
	if n < 1 {
		return 0
	}

	_, v := fastDoubling(n-1)
	return v
}

func fastDoubling(n uint)(uint, uint){
	if n == 0 {
		return 0, 1
	}

	a, b := fastDoubling(n >> 1)
	c := a * ((b<<1)-a)
	d := a*a + b*b

	if n&1 == 0 {
		return c, d
	}

	return d, c+d	
}