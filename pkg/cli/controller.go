package cli

const (
	PNG  int = 0
	JPEG int = 1
	PSD  int = 2
	SVG  int = 3
	PDF  int = 4
	TXT  int = 5
	JSON int = 6
	DOT  int = 7
)

type clicontroller struct {
	dir    string
	output string
}

func NewController() clicontroller {
	return clicontroller{
		dir:    "",
		output: "",
	}
}

func (c clicontroller) getDir() {}
