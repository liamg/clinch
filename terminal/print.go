package terminal

import "github.com/liamg/tml"

// PrintErrorf prints a string to stdout in bold, red text
func PrintErrorf(message string, args ...interface{}) {
	tml.Printf("<red><bold>%s", append([]interface{}{message}, args...)...)
}

// PrintImportantf prints a string to stdout in bold, light blue text
func PrintImportantf(message string, args ...interface{}) {
	tml.Printf("<lightblue><bold>%s", append([]interface{}{message}, args...)...)
}

// PrintSuccessf prints a string to stdout in bold, green text
func PrintSuccessf(message string, args ...interface{}) {
	tml.Printf("<green><bold>%s", append([]interface{}{message}, args...)...)
}
