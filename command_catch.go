package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("no Pokemon specified")
	}

	if len(args) > 1 {
		return errors.New("too many Pokemon")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("throwing a pokeball at %s...\n", pokemon.Name)
	if chance := rand.Intn(pokemon.BaseExperience); chance > 40 {
		fmt.Printf("caught %s!\n", pokemon.Name)
		fmt.Println("you may now inspect it with the 'inspect' command")
	} else {
		fmt.Printf("%s escaped.\n", pokemon.Name)
	}

	cfg.caughtPokemon[pokemon.Name] = pokemon

	return nil
}
