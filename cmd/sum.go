package cmd

import "strconv"

//SumCommand - sum command
type SumCommand struct{}

//Execute - sum
func (p *SumCommand) Execute(args []string) string {

	x, _ := strconv.Atoi(args[2])

	y, _ := strconv.Atoi(args[3])

	return strconv.Itoa(x + y)
}
