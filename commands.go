package main

import (
	"fmt"
	"os"
)

type CliCommand struct{
	name string
	description string
	callback func()error
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
}

func commandExit() error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error{
	fmt.Println("Welcome to the Pokedex!")
	return nil
}
