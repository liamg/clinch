package terminal

import "fmt"

func Reset() {
	fmt.Println("\x1b[3J")
}

func Clear() {
	fmt.Printf("\x1b[2J")
}

func SetAltBuffer() {
	fmt.Printf("\x1b[?1049h")
}

func SetMainBuffer() {
	fmt.Printf("\x1b[?1049l")
}

func SaveCursor() {
	fmt.Printf("\033[s")
}

func RestoreCursor() {
	fmt.Printf("\033[u")
}

func ShowCursor() {
	fmt.Printf("\033[?25h")
}

func HideCursor() {
	fmt.Printf("\033[?25l")
}

// zero indexed
func MoveCursorToColumn(column int) {
	fmt.Printf("\r\033[%dC", column+1)
}

func MoveCursorTo(column int, row int) {
	fmt.Printf("\033[%d;%dH", row+1, column+1)
}

func ClearLine() {
	fmt.Printf("\033[2K\r")
}
