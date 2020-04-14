package main

import (
	"fmt"

	"github.com/liamg/clinch/prompt"
)

func main() {
	input := prompt.EnterInput("Enter your name: ")
	fmt.Printf("You entered '%s'\n", input)

	searchEngineUrl := prompt.EnterInputWithDefault("Set search engine url", "https://www.google.com?q=")
	fmt.Printf("Search engine URL set to '%s'\n", searchEngineUrl)
}
