package cmd

import (
	"fmt"
	"math"
	"strconv"
)

//PrimeCommand - sum command
type PrimeCommand struct{}

//Execute - sum
func (p *PrimeCommand) Execute(args []string) string {

	N, _ := strconv.Atoi(args[2])
	counter := 0

	start := 3

	if N >= 1 {
		fmt.Println("first Prime Number : " + "2")
	}

	for counter != N {

		if IsPrime(start) {
			fmt.Println(start)
			counter++
		}

		start++
	}

	return ""

}

//IsPrime -
func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
