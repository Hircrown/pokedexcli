package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name 		string
	description string
	callback    func() error
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name: "help",
			description: "Display a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}
}


func commandHelp() error {
	helpMessage := "Welcome to the Pokedex!\nUsage:\n\n"

	for _, v := range getCommands() {
		helpMessage += fmt.Sprintf("%s: %s\n", v.name, v.description)
	}
	fmt.Println(helpMessage)
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}