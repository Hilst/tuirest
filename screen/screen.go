package screen

import (
	"github.com/Hilst/tuirest/suggestions"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const INVALID_INPUT = `\033[7;31""mInvalid Input\033[0m`

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

func (s *screen) makeInput() *screen {
	s.input = tview.NewInputField().
		SetLabel("> ").
		SetAutocompleteFunc(suggestions.Matches).
		Autocomplete().
		SetDoneFunc(func(key tcell.Key) {
			resultText := s.input.GetText()
			if !suggestions.Valid(resultText) {
				resultText = INVALID_INPUT
			}
			s.output.Write([]byte(resultText + "\n"))
			s.input.SetText("")
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
