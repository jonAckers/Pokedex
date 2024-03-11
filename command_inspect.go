package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("no Pokemon specified")
	}

	if len(args) > 1 {
		return errors.New("too many Pokemon")
	}

	pokemon, ok := cfg.caughtPokemon[args[0]]
	if !ok {
		return errors.New("you haven't caught that Pokemon")
	}

	fmt.Printf("name: %s\n", pokemon.Name)
	fmt.Printf("height: %d\n", pokemon.Height)
	fmt.Printf("weight: %d\n", pokemon.Weight)
	fmt.Println("stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("types:")
	for _, t := range pokemon.Types {
		fmt.Println("  -", t.Type.Name)
	}

	return nil
}
