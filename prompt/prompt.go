package prompt

import (
	"bufio"
	"os"

	"github.com/liamg/clinch/terminal"
)

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
