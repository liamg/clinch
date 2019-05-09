package prompt

import (
	"fmt"
	"strconv"

	"github.com/liamg/tml"
)

var ErrUserCancelled = fmt.Errorf("User cancelled")
var ErrUserChoiceInvalid = fmt.Errorf("User cancelled")

func ChooseFromList(message string, options []string) (int, string, error) {
	fmt.Printf("\n %s\n\n", message)
	colours := []string{"lightblue", "lightgreen", "lightyellow", "white"}

	for i, option := range options {
		col := colours[i%len(colours)]
		pad := ""
		if i+1 < 10 { // ocd padding
			pad = " "
		}
		tml.Printf(
			fmt.Sprintf(" %%s<darkgrey>[</darkgrey><%s>%%d<darkgrey>]</darkgrey> <%s>%%s\n", col, col),
			pad,
			i+1,
			option,
		)
	}
	fmt.Println("")
	choice := EnterInput("Enter choice (blank to cancel): ")
	fmt.Println("")

	if choice == "" {
		return -1, "", ErrUserCancelled
	}

	for i, opt := range options {
		if opt == choice {
			return i, opt, nil
		}
	}

	choiceIndex, err := strconv.Atoi(choice)
	if err != nil || choiceIndex-1 >= len(options) || choiceIndex <= 0 {
		return -1, "", ErrUserCancelled
	}

	return choiceIndex - 1, options[choiceIndex-1], nil
}
