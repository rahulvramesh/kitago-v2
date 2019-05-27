package cmd

import "strconv"

//MultiplyCommand - sum command
type MultiplyCommand struct{}

//Execute - sum
func (p *MultiplyCommand) Execute(args []string) string {

	x, _ := strconv.Atoi(args[2])

	y, _ := strconv.Atoi(args[3])

	return strconv.Itoa(x * y)
}
