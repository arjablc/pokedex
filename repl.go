package main

import (
	"errors"
	"fmt"
	"github.com/arjablc/pokedex/internals/api"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*UrlConfig, *api.ApiClient) error
}

type UrlConfig struct {
	previousUrl *string
	nextUrl     *string
	areaName    *string
}

var commandsMap = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"map": {
		name:        "map",
		description: "Show the location areas or next location areas",
		callback:    commandMap,
	},
	"mapb": {
		name:        "map",
		description: "Show the previous location areas",
		callback:    commandMapb,
	},
	"explore": {
		name:        "explore",
		description: "Explore the given area, return the pokemons found",
		callback:    commandExplore,
	},
}

func cleanInput(text string) []string {
	out := strings.ToLower(text)
	words := strings.Fields(out)
	return words
}

func commandHelp(config *UrlConfig, _ *api.ApiClient) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandExit(config *UrlConfig, _ *api.ApiClient) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMapb(config *UrlConfig, apiClient *api.ApiClient) error {
	url := config.previousUrl
	if url == nil {
		println("You're on the first page")
		return nil
	}
	locationResponse := apiClient.RequestLocationArea(*url)
	config.nextUrl = locationResponse.Next
	config.previousUrl = locationResponse.Previous
	for _, location := range locationResponse.Results {
		println(location.Name)
	}
	return nil
}

func commandMap(config *UrlConfig, apiClient *api.ApiClient) error {
	url := config.nextUrl
	if url == nil {
		url = &api.LocationsUrl
	}

	locationResponse := apiClient.RequestLocationArea(*url)
	config.nextUrl = locationResponse.Next
	config.previousUrl = locationResponse.Previous
	for _, location := range locationResponse.Results {
		println(location.Name)
	}
	return nil
}

func commandExplore(config *UrlConfig, apiClient *api.ApiClient) error {
	if config.areaName == nil {
		return errors.New("No location provided")
	}
	locationName := config.areaName
	fmt.Println(*locationName)
	url := config.nextUrl
	if url == nil {
		url = &api.LocationsUrl
	}

	pokemonsResponse := apiClient.RequestPokemons(*locationName)
	for _, pokemon := range pokemonsResponse.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
