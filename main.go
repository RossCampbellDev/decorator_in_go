package main

import "fmt"

type Operation func(int) int // the Operatior type is a function that takes an int and returns an int

// this function contains the base line functionality - in this case, incrementing n
func FirstOp(n int) int {
	return n + 1
}

// this func will take in an operation and augment the behaviour before returning it
func SecondOp(op Operation) Operation {
	return func(n int) int {
		x := op(n)
		return x + 2
	}
}

// composite the baseline func with the decorator one and run them together
func main() {
	test := SecondOp(FirstOp)
	fmt.Println(test(5)) // prints out 8
}
