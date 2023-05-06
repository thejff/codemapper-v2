package cli

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type inputModel struct {
	textInput textinput.Model
	err       error
}

type (
	errMsg error
)

func startInput() error {
	p := tea.NewProgram(initialInputModel())

	d, err := p.Run()

	if err != nil {
		return err
	}

	fmt.Printf("output %v", d)

	return nil
}

func initialInputModel() inputModel {
	ti := textinput.New()

	ti.Placeholder = "~/mycode"
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 20

	return inputModel{
		textInput: ti,
		err:       nil,
	}
}

func (im inputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (im inputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:

			fmt.Println("Enter or something pressed")
			return im, tea.Quit
		}

	case errMsg:
		im.err = msg
		return im, nil

	}

	im.textInput, cmd = im.textInput.Update(msg)

	return im, cmd
}

func (im inputModel) View() string {
	return fmt.Sprintf(
		"What directory would you like to map?\n\n%s\n\n%s",
		im.textInput.View(),
		"(esc to go back)",
	) + "\n"
}
