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

		fmt.Printf("Your command was: %s\n", words[0])
	}
}
