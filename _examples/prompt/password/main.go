package main

import (
	"fmt"

	"github.com/liamg/clinch/prompt"
)

func main() {
	password := prompt.EnterPassword("Enter password:")
	fmt.Printf("You entered '%s'\n", password)
}
