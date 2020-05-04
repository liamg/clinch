package main

import (
	"fmt"
	"strings"

	"github.com/liamg/clinch/prompt"
)

func main() {

	options := []string{
		"apple",
		"banana",
		"orange",
		"pear",
		"grapes",
		"grapefruit",
	}

	_, selected, err := prompt.ChooseFromMultiList("Choose a fruit for the salad!", options)
	if err != nil   {
		println(err.Error())
		return
	}

	fmt.Printf("The salad will contain delicious %s!\n", strings.Join(selected, ", "))
}
