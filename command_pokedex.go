package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("you haven't caught any Pokemon yet!")
		return nil
	}

	fmt.Println("you've caught:")
	for name := range cfg.caughtPokemon {
		fmt.Println("  -", name)
	}

	return nil
}
