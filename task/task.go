package task

import (
	"fmt"
	"time"

	"github.com/liamg/clinch/terminal"
	"github.com/liamg/tml"
)

type Task struct {
	function    func() error
	category    string
	description string
}

// New creates a new task
func New(category string, description string, f func() error) *Task {
	return &Task{
		function:    f,
		category:    category,
		description: description,
	}
}

// Run runs the task, providing animated output as it does so. If the task fails, the error from the task function will be returned here.
func (t *Task) Run() error {
	terminal.HideCursor()
	defer terminal.ShowCursor()
	terminal.ClearLine()
	tml.Printf("<lightblue>% 10s</lightblue> %s", t.category, t.description)
	terminal.MoveCursorToColumn(74)
	tml.Printf("<bold><darkgrey>[    ]</darkgrey></bold>")
	stopAnimationChan := make(chan struct{})
	go func() {

		ticker := time.NewTicker(time.Millisecond * 250)
		defer ticker.Stop()

		frameIndex := 0

		frames := []string{
			"*   ",
			" *  ",
			"  * ",
			"   *",
			"   *",
			"  * ",
			" *  ",
			"*   ",
		}

		for {
			select {
			case <-stopAnimationChan:
				return
			case <-ticker.C:
				frame := frames[(frameIndex)%len(frames)]
				frameIndex++
				terminal.SaveCursor()
				terminal.HideCursor()
				terminal.MoveCursorToColumn(75)
				tml.Printf("<yellow>%s</yellow>", frame)
				terminal.RestoreCursor()
			}
		}
	}()
	err := t.function()
	stopAnimationChan <- struct{}{}
	terminal.MoveCursorToColumn(75)
	if err != nil {
		tml.Printf("<red><bold>FAIL</bold></red>")
	} else {
		tml.Printf("<green><bold> OK </bold></green>")
	}
	terminal.MoveCursorToColumn(80)
	fmt.Printf("\n")
	return err
}
