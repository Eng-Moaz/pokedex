package main

import (
	"encoding/json"
	"fmt"
	"os"

	pokecache "github.com/Eng-Moaz/pokedex/internal"
)

type CliCommand struct{
	name string
	description string
	callback func(*Config)error
}

type Config struct{
	Next *string
	Previous *string
	firstTime bool
	cache pokecache.Cache
}

var AllCommands = map[string]CliCommand{
	"exit": {
		name : "exit",
		description : "Exit the Pokedex",
		callback: commandExit,
	},
	"help": {
		name : "help",
		description: "Displays a help message",
		callback: commandHelp,
	},
	"map": {
		name : "map",
		description : "Displays 20 locations",
		callback:  commandMap,
	},
	"mapb": {
		name : "mapb",
		description: "Displays the previous 20 locations",
		callback: commandMapb,
	},
}

func commandExit(cfg *Config) error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error{
	fmt.Println("Welcome to the Pokedex!")
	return nil
}

func commandMap(cfg *Config) error{
	var locationArea LocationAreaResponse
	var apiUrl string
	if cfg.firstTime{
		apiUrl = "https://pokeapi.co/api/v2/location-area"
		cfg.firstTime = false
	}else{
		apiUrl = *cfg.Next
	}
	if value, ok := cfg.cache.Get(apiUrl);ok{
		err := json.Unmarshal(value, &locationArea)
		if err != nil{
			return fmt.Errorf("Couldn't Unmarshal: %v", err)
		}
	}else{
		var err error
		locationArea, err = ReqAndUnmarshal(apiUrl)
		if err != nil{
			return fmt.Errorf("Failed to get locationArea struct: %v", err)
		}
	}
	for i := range 20{
		fmt.Println(locationArea.Results[i].Name)
	}
	cfg.Next = locationArea.Next
	cfg.Previous = locationArea.Previous
	return nil
}


func commandMapb(cfg *Config) error{
	var locationArea LocationAreaResponse
	var apiUrl string
	if cfg.firstTime{
		apiUrl = "https://pokeapi.co/api/v2/location-area"
		cfg.firstTime = false
	}else{
		apiUrl = *cfg.Previous
	}
	if value, ok := cfg.cache.Get(apiUrl);ok{
		err := json.Unmarshal(value, &locationArea)
		if err != nil{
			return fmt.Errorf("Couldn't Unmarshal: %v", err)
		}
	}else{
		var err error
		locationArea, err = ReqAndUnmarshal(apiUrl)
		if err != nil{
			return fmt.Errorf("Failed to get locationArea struct: %v", err)
		}
	}
	for i := range 20{
		fmt.Println(locationArea.Results[i].Name)
	}
	cfg.Next = locationArea.Next
	cfg.Previous = locationArea.Previous
	return nil
}
