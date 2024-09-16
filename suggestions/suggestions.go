package suggestions

import "strings"

var commands = []string{
	"open",
}

func Matches(currentText string) (entries []string) {
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
}

func Valid(text string) bool {
	if text == "" {
		return false
	}
	for _, cmd := range commands {
		if cmd == text {
			return true
		}
	}
	return false
}
