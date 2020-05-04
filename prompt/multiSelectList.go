package prompt

import (
	"fmt"
	"github.com/liamg/clinch/terminal"
	"github.com/liamg/tml"
	"github.com/pkg/term"
)

const (
	SPACE  = 32
	UP     = 38
	DOWN   = 40
	ESCAPE = 27
)

type listItem struct {
	index    int
	value    string
	selected bool
	colour   string
}

func (item *listItem) toString() string {
	check := " "
	if item.selected {
		check = "X"
	}
	return fmt.Sprintf(" <darkgrey>[</darkgrey>%v<darkgrey>]</darkgrey> <%s>%v\n", check, item.colour, item.value)
}

func ChooseFromMultiList(message string, options []string) ([]int, []string, error) {
	var items []*listItem
	colours := []string{"lightblue", "lightgreen", "lightyellow", "white"}
	for index, option := range options {
		col := colours[index%len(colours)]
		items = append(items, &listItem{index: index, value: option, colour: col})
	}
	return getListSelection(message, items)
}

func getListSelection(message string, items []*listItem) ([]int, []string, error) {
	fmt.Printf("\n %s\n\n", message)
	currentPos := 0
	fmt.Println(" Press escape to cancel, space to toggle, return to accept.")
	fmt.Println("")
	drawItems(items, currentPos, false)

keyInput:
	for {
		keyCode, err := getKeyInput()
		if err != nil {
			panic(err)
		}
		switch keyCode {
		case DOWN:
			if currentPos < len(items)-1 {
				terminal.MoveCursorDown(1)
				currentPos += 1
			}
		case UP:
			if currentPos > 0 {
				terminal.MoveCursorUp(1)
				currentPos -= 1
			}
		case SPACE:
			items[currentPos].selected = !items[currentPos].selected
			drawItems(items, currentPos, true)
		case ESCAPE:
			terminal.MoveCursorDown(len(items) - currentPos + 1)
			return []int{}, []string{}, ErrUserCancelled
		default:
			break keyInput
		}
	}
	terminal.MoveCursorDown(len(items) - currentPos + 1)

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

func drawItems(items []*listItem, currentPos int, isRedraw bool) {
	if isRedraw {
		terminal.MoveCursorUp(currentPos)
		terminal.MoveCursorToColumn(-2)
	}

	for _, item := range items {
		_ = tml.Printf(item.toString())
	}
	fmt.Println("")
	terminal.MoveCursorUp(len(items) + 1 - currentPos)
	terminal.MoveCursorToColumn(1)
}

func getKeyInput() (keyCode int, err error) {
	t, err := term.Open("/dev/tty")
	if err != nil {
		panic(err)
	}
	err = term.RawMode(t)
	if err != nil {
		panic(err)
	}
	bytes := make([]byte, 3)

	var numRead int
	numRead, err = t.Read(bytes)
	if err != nil {
		return
	}
	if numRead == 3 && bytes[0] == 27 && bytes[1] == 91 {
		if bytes[2] == 65 {
			keyCode = UP
		} else if bytes[2] == 66 {
			keyCode = DOWN
		}
	} else if numRead == 1 {
		if bytes[0] == 27 {
			keyCode = ESCAPE
		}
		if int(bytes[0]) == ' ' {
			keyCode = SPACE
		}
	} else {
		print(int(bytes[2]))
	}
	t.Restore()
	t.Close()
	return
}
