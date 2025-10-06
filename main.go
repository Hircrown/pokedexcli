package main

const cliCustomCmd = "Pokedex > "

func main() {
	cfg := &config{}
	startRepl(cliCustomCmd, cfg)
}
