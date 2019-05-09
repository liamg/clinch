package main

import (
	"fmt"

	"github.com/liamg/clinch/prompt"
)

func main() {

	options := []string{
		"apple",
		"banana",
		"orange",
		"pear",
	}

	_, fruit, err := prompt.ChooseFromList("Choose a fruit!", options)
	if err != nil { // user cancelled
		panic(err)
	}

	fmt.Printf("A delicious %s!\n", fruit)
}
