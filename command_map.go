package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, _ []string) error {
	if cfg.nextLocationsUrl == nil && cfg.prevLocationsUrl != nil {
		return errors.New("end of locations reached")
	}

	locResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = locResp.Next
	cfg.prevLocationsUrl = locResp.Previous

	for _, loc := range locResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config, _ []string) error {
	if cfg.prevLocationsUrl == nil {
		return errors.New("start of locations reached")
	}

	locResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = locResp.Next
	cfg.prevLocationsUrl = locResp.Previous

	for _, loc := range locResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
