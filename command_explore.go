package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("no location specified")
	}

	if len(args) > 1 {
		return errors.New("too many locations")
	}

	locResp, err := cfg.pokeapiClient.GetLocation(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("exploring %s...\n", locResp.Name)
	fmt.Println("found Pokemon:")
	for _, p := range locResp.PokemonEncounters {
		fmt.Println("  -", p.Pokemon.Name)
	}

	return nil
}
