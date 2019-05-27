package cmd

import (
	"fmt"
	"strconv"
)

//FibonacciCommand - sum command
type FibonacciCommand struct{}

//Execute - sum
func (p *FibonacciCommand) Execute(args []string) string {

	N, _ := strconv.Atoi(args[2])

	//print fib

	for i := 0; i <= N-1; i++ {
		fmt.Print(strconv.Itoa(p.FibonacciRecursion(i)) + " ")
	}

	return ""
}

//FibonacciRecursion -
func (p *FibonacciCommand) FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return p.FibonacciRecursion(n-1) + p.FibonacciRecursion(n-2)
}
