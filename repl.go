package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		userCommand := words[0]
		if command, ok := getCommands()[userCommand]; !ok {
			fmt.Println("Command unknown")
			continue
		} else {
			if err := command.callback(); err != nil {
				fmt.Println(err)
			}
			continue
		}
		
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	return words
}