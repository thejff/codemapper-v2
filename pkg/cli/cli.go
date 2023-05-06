package cli

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func StartCLI() {
	p := tea.NewProgram(initialStartMenu())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}

}
