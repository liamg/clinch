package loading

import (
	"fmt"
	"sync"

	"github.com/liamg/tml"

	"github.com/liamg/clinch/terminal"
)

type Bar struct {
	lock     sync.Mutex
	complete bool
}

// NewBar creates a new loading bar
func NewBar() *Bar {
	return &Bar{}
}

// SetPercent sets the current percentage completion of the bar, and renders it to the current line
func (bar *Bar) SetPercent(percent float64) {

	bar.lock.Lock()
	defer bar.lock.Unlock()
	if bar.complete {
		return
	}

	w, h := terminal.Size()
	terminal.HideCursor()
	terminal.MoveCursorTo(0, h-1)

	if percent > 100 {
		percent = 100
	} else if percent < 0 {
		percent = 0
	}

	barWidth := w - 5

	size := int(percent * float64(barWidth) / 100)
	remaining := int((100 - percent) * float64(barWidth) / 100)

	terminal.ClearLine()

	for i := 0; i < size; i++ {
		tml.Printf("<bg-blue> </bg-blue>")
	}
	for remaining > 0 {
		fmt.Printf(" ")
		remaining--
	}
	fmt.Printf(" %d%%", int(percent))
}

// Log logs a messages to the terminal without disrupting the rendered loading bar
func (bar *Bar) Log(msg string) {
	bar.lock.Lock()
	defer bar.lock.Unlock()
	terminal.ClearLine()
	fmt.Println(msg)
}

// Close hides the loading bar and tidies up
func (bar *Bar) Close() {
	bar.lock.Lock()
	defer bar.lock.Unlock()
	bar.complete = true
	_, h := terminal.Size()
	terminal.MoveCursorTo(0, h-1)
	terminal.ClearLine()
	terminal.ShowCursor()
}
