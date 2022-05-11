package task

import (
	"fmt"
	"sync"
	"time"

	"github.com/liamg/clinch/terminal"
	"github.com/liamg/tml"
)

type Task struct {
	function          func(t *Task) error
	category          string
	description       string
	animationStopped  bool
	stopAnimationChan chan struct{}
	skipped           bool
	noOutput          bool
}

type Option func(*Task)

func WithNoOutput() Option {
	return func(t *Task) {
		t.noOutput = true
	}
}

// New creates a new task.
func New(category string, description string, f func(t *Task) error, opts ...Option) *Task {
	t := &Task{
		function:    f,
		category:    category,
		description: description,
	}

	for _, opt := range opts {
		opt(t)
	}

	return t
}

func (t *Task) Skip() {
	t.skipped = true
}

func (t *Task) StopAnimation() {
	if t.animationStopped {
		return
	}
	t.animationStopped = true
	t.stopAnimationChan <- struct{}{}
}

func (t *Task) printSkeleton() {
	terminal.ClearLine()
	tml.Printf("<lightblue>% 10s</lightblue> %s", t.category, t.description)
	terminal.MoveCursorToColumn(74)
	tml.Printf("<bold><darkgrey>[    ]</darkgrey></bold>")
}

// Run runs the task, providing animated output as it does so. If the task fails, the error from the task function will be returned here.
func (t *Task) Run() error {
	if t.noOutput {
		return t.function(t)
	}

	terminal.HideCursor()
	defer terminal.ShowCursor()
	t.printSkeleton()
	t.stopAnimationChan = make(chan struct{})
	wait := sync.WaitGroup{}
	wait.Add(1)
	go func() {

		ticker := time.NewTicker(time.Millisecond * 250)
		defer wait.Done()
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
			case <-t.stopAnimationChan:
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
	err := t.function(t)
	t.StopAnimation()
	wait.Wait()
	t.printSkeleton()
	terminal.MoveCursorToColumn(75)
	if err != nil {
		tml.Printf("<red><bold>FAIL</bold></red>")
	} else if t.skipped {
		tml.Printf("<yellow><bold>SKIP</bold></yellow>")
	} else {
		tml.Printf("<green><bold> OK </bold></green>")
	}
	terminal.MoveCursorToColumn(80)
	fmt.Printf("\n")
	return err
}
