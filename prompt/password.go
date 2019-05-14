package prompt

import (
	"bufio"
	"os"
	"strings"

	"github.com/liamg/clinch/terminal"
	"github.com/liamg/tml"
)

// EnterPassword requests input from the user with the given message, hiding that input, and returns any user input that was gathered until a newline was entered
func EnterPassword(msg string) string {

	terminal.ClearLine()
	terminal.PrintImportantf(msg)

	parser := tml.NewParser(os.Stdout)
	parser.IncludeTrailingResets = false
	parser.Parse(strings.NewReader("<hidden>"))
	terminal.HideCursor()
	defer func() {
		parser.Parse(strings.NewReader("</hidden>"))
		terminal.ShowCursor()
	}()

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
