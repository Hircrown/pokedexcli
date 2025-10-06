package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}

func startRepl(placeholder string) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(placeholder)
		scanner.Scan()

		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}

		cmdName := words[0]

		cmd, exists := getCommands()[cmdName]
		if !exists {
			fmt.Println("Unknown command")
		} else {
			if err := cmd.callback(); err != nil {
				fmt.Errorf("Callback error: %w", err)
			}
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Give an overview of cli commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
