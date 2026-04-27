package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/arjablc/pokedex/internals/api"
	"github.com/arjablc/pokedex/internals/pokecache"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(5 * time.Second)
	apiClient := api.ApiClient{Cache: cache}
	cfg := config{previousUrl: nil, nextUrl: nil, apiClient: &apiClient}
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}
		command := words[0]
		entry, exists := commandsMap[command]
		if exists {
			args := words[1:]
			err := entry.callback(&cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		fmt.Println("Unknown command")
	}

}
