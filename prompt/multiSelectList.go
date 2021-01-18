package prompt

import (
	"fmt"
	"sort"

	"github.com/liamg/clinch/terminal"
	"github.com/liamg/keyboard"
	"github.com/liamg/tml"
)

const (
	ROW_OFFSET     = 2
	DEFAULT_COLUMN = 0
)

type listItem struct {
	index    int
	value    string
	selected bool
	colour   string
}

var ErrNoOptionsProvided = fmt.Errorf("no options were provided")

func (item *listItem) toString() string {
	check := " "
	if item.selected {
		check = "X"
	}
	return fmt.Sprintf(" <darkgrey>[</darkgrey>%v<darkgrey>]</darkgrey> <%s>%v\n\r", check, item.colour, item.value)
}

func ChooseFromMultiList(message string, options []string) ([]int, []string, error) {
	if len(options) == 0 {
		return nil, nil, ErrNoOptionsProvided
	}
	sort.Strings(options)
	var items []*listItem
	colours := []string{"lightblue", "lightgreen", "lightyellow", "white"}
	for index, option := range options {
		col := colours[index%len(colours)]
		items = append(items, &listItem{index: index, value: option, colour: col})
	}
	return getListSelection(message, items)
}

func getListSelection(message string, items []*listItem) ([]int, []string, error) {
	fmt.Printf("\n %s\n\r", message)
	currentPos := 0
	drawItems(items, currentPos, false)

	if err := keyboard.Open(); err != nil {
		return nil, nil, err
	}
	defer func() {
		_ = keyboard.Close()
	}()

keyInput:
	for {
		_, keyCode, err := keyboard.GetSingleKey()
		if err != nil {
			return nil, nil, err
		}
		switch keyCode {
		case keyboard.KeyArrowDown:
			if currentPos < len(items)-1 {
				terminal.MoveCursorDown(1)
				currentPos += 1
			}
		case keyboard.KeyArrowUp:
			if currentPos > 0 {
				terminal.MoveCursorUp(1)
				currentPos -= 1
			}
		case keyboard.KeySpace:
			items[currentPos].selected = !items[currentPos].selected
			drawItems(items, currentPos, true)
		case keyboard.KeyEsc:
			resetPrompt(len(items) - currentPos)
			return []int{}, []string{}, ErrUserCancelled
		case keyboard.KeyEnter:
			break keyInput
		}
	}
	resetPrompt(len(items) - currentPos)

	var selectedIndexes []int
	var selectedValues []string

	for _, item := range items {
		if !item.selected {
			continue
		}
		selectedIndexes = append(selectedIndexes, item.index)
		selectedValues = append(selectedValues, item.value)
	}
	return selectedIndexes, selectedValues, nil
}

func resetPrompt(rowPosition int) {
	terminal.MoveCursorDown(rowPosition + ROW_OFFSET - 1)
	terminal.ClearLine()
	terminal.MoveCursorToColumn(DEFAULT_COLUMN)
}

func drawItems(items []*listItem, currentPos int, isRedraw bool) {
	if isRedraw {
		terminal.MoveCursorUp(currentPos)
	}
	fmt.Print("\r")

	for _, item := range items {
		_ = tml.Printf(item.toString())
	}
	fmt.Println("")
	fmt.Println(" space to toggle, return to accept. (Esc to cancel): ")
	terminal.MoveCursorUp(len(items) - currentPos + ROW_OFFSET)
	terminal.MoveCursorToColumn(1)
}
