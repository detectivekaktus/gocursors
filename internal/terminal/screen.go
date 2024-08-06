package terminal

import (
	"os"

	"golang.org/x/term"
)

func IsTerminal() bool {
  return term.IsTerminal(int(os.Stdout.Fd()))
}

func GetTerminalSize() (int, int, error) {
  return term.GetSize(int(os.Stdout.Fd()))
}
