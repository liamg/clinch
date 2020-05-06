package main

import (
	"fmt"
	"github.com/liamg/clinch/prompt"
	"strings"
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

	_, selected, err := prompt.ChooseFromMultiList("Choose the fruit for the salad!", options)
	if err != nil {
		fmt.Printf("%v\n",err.Error())
		return
	}
	if len(selected) == 0 {
		fmt.Println("That will be a boring fruit salad!!")
	} else {
		fmt.Printf("The salad will contain delicious %s!\n", strings.Join(selected, ", "))
	}
}
