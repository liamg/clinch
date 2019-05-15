package main

import (
	"fmt"

	"github.com/liamg/clinch/prompt"
)

func main() {
	password := prompt.EnterPassword("Enter password:")
	fmt.Printf("\nYou entered '%s'\n", password)
}
