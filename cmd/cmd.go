package cmd

//Command - interfacr for all commands
type Command interface {
	Execute([]string) string
}

//ExecByName - execute by name
func ExecByName(name string, args []string) string {
	// Register commands
	commands := map[string]Command{
		"sum":       &SumCommand{},
		"multiply":  &MultiplyCommand{},
		"prime":     &PrimeCommand{},
		"fibonacci": &FibonacciCommand{},
	}

	if command := commands[name]; command == nil {
		return "No such command found, throw error?"
	} else {
		return command.Execute(args)
	}
}
