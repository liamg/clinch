package prompt

import (
	"bufio"
	"os"

	"github.com/liamg/clinch/terminal"
)

// EnterInput requests input from the user with the given message, and returns any user input that was gathered until a newline was entered
func EnterInput(msg string) string {
	terminal.ClearLine()
	terminal.PrintImportantf(msg)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil || len(input) <= 1 {
		return ""
	}
	s := input[:len(input)-1]
	if s[len(s)-1] == '\r' {
		s = input[:len(input)-1]
	}
	return s
}
