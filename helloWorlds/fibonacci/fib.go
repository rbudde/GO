package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Initialize two big ints with the first two numbers in the sequence.
	ai := int64(0)
	a := big.NewInt(ai)
	b := big.NewInt(1)

	// Initialize limit as 10^199, the smallest integer with 200 digits.
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(199), nil)

	// Loop while a is smaller than 1e200.
	for a.Cmp(&limit) < 0 {
		// Compute the next Fibonacci number, storing it in a.
		a.Add(a, b)
		// Swap a and b so that b is the next number in the sequence.
		a, b = b, a
	}
	fmt.Println(a) // 200-digit Fibonacci numb
}
