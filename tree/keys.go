package tree

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	CursorUp     key.Binding
	CursorDown   key.Binding
	CursorSelect key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		CursorUp: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		CursorDown: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		CursorSelect: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "select row"),
		),
	}
}
