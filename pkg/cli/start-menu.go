package cli

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type startMenu struct {
	options []string
	cursor  int
	// selected map[int]struct{}
	actions map[int]interface{}
}

func initialStartMenu() startMenu {
	sm := startMenu{
		options: []string{
			"Select directory to map",
			"Map the current directory",
		},

		actions: make(map[int]interface{}),
	}

	sm.actions[0] = startInput
	sm.actions[1] = startInput

	return sm
}

func (sm startMenu) Init() tea.Cmd {
	return nil
}

func (sm startMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return sm, tea.Quit

		case "up", "k":

			if sm.cursor > 0 {
				sm.cursor--
			}

		case "down", "j":
			if sm.cursor < len(sm.options)-1 {
				sm.cursor++
			}

		case "enter", " ":
			fn := sm.actions[sm.cursor]
			if err := fn.(func() error)(); err != nil {
				log.Fatal(err)
			}

		}
	}

	return sm, nil
}

func (sm startMenu) View() string {

	s := "Welcome to the Just For Fun Foundations Codemapper, what would you like to do?\n\n"

	for i, option := range sm.options {

		cursor := " "
		if sm.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, option)
	}

	s += "\nPress q to quit\n"

	return s
}
