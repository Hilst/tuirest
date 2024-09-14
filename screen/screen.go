package screen

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type IScreen interface {
	Run() error
}

type screen struct {
	app    *tview.Application
	input  *tview.InputField
	output *tview.TextView
	flex   *tview.Flex
}

func (s *screen) Run() error {
	return s.app.SetRoot(s.flex, true).Run()
}

func MakeScreen() IScreen {
	return start().
		makeApp().
		makeOutput().
		makeInput().
		makeFlex()
}

func start() *screen {
	return &screen{}
}

func (s *screen) makeApp() *screen {
	s.app = tview.NewApplication()
	return s
}

func (s *screen) makeOutput() *screen {
	s.output = tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetWordWrap(true).
		SetChangedFunc(func() { s.app.Draw() })
	return s
}

var commands = []string{"help", "exit", "list", "clear", "show", "delete", "add", "update"}

func (s *screen) makeInput() *screen {
	s.input = tview.NewInputField().
		SetLabel("> ").
		SetAutocompleteFunc(func(currentText string) (entries []string) {
			if currentText == "" {
				entries = append(entries, commands...)
				return
			}
			for _, cmd := range commands {
				if strings.HasPrefix(cmd, currentText) {
					entries = append(entries, cmd)
				}
			}
			return
		}).
		Autocomplete().
		SetDoneFunc(func(key tcell.Key) {
			command := s.input.GetText()
			if command != "" {
				s.output.Write([]byte(command + "\n"))
				s.input.SetText("")
			}
		})
	return s
}

func (s *screen) makeFlex() *screen {
	s.flex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(s.output, 0, 2, false).
		AddItem(s.input, 1, 1, true)
	return s
}
