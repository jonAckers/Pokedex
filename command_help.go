package main

import "fmt"

func commandHelp(_ *config, _ []string) error {
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
