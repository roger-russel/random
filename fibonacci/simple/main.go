package main

import "fmt"

type Fibonacci interface {
	Get(n int) int
	Calc(n, currN, penultimo, previous int) int
}

type FiboImpl struct {}


func NewFibonacci() Fibonacci {
	return &FiboImpl{}
}

func main() {
	fib := NewFibonacci()
	v := fib.Get(3)
	fmt.Println(v)
}

/*
	0 : 0
	1 : 1
	2 : 1
	3 : 2
	4 : 3
	5 : 5
	6 : 8 = 23
	7 : 13
	8 : 21 = 3 x 7
	9 : 34 = 2 x 17
	10 : 55 = 5 x 11

*/
func (f *FiboImpl) Get(n int) int {
	if n == 0 { 
		return 0
	}

	if n == 1 {
		return 1
	}

	penultimo := 0
	previous := 1
	currN := 2
	return f.Calc(n, currN, penultimo, previous)
}

func (f *FiboImpl) Calc(n, currN, pnultimo, previous int) int {
	currValue := pnultimo + previous

	if currN == n {
		return currValue
	}

	currN++
	return f.Calc(n, currN, previous, currValue)
}