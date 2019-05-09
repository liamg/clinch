package main

import (
	"fmt"

	"github.com/liamg/clinch/prompt"
)

func main() {
	input := prompt.EnterInput("Enter your name: ")
	fmt.Printf("You entered '%s'\n", input)
}
