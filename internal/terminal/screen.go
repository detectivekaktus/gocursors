package terminal

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func IsTerminal() bool {
  return term.IsTerminal(int(os.Stdout.Fd()))
}

func GetTerminalSize() (int, int, error) {
  return term.GetSize(int(os.Stdout.Fd()))
}

func MakeRaw() *term.State {
  oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
  if err != nil {
    fmt.Println("FATAL ERROR: Couldn't enter into raw mode on the current terminal.")
    os.Exit(1)
  }
  return oldState
}

func ApplyState(state *term.State) {
  term.Restore(int(os.Stdin.Fd()), state)
}
