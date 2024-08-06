package gocursors

import (
	"fmt"
	"os"

	"github.com/detectivekaktus/gocursors/internal/terminal"
	"golang.org/x/term"
)

type Window struct {
  CurX   int
  CurY   int
  Width  int
  Height int
}

func GoCursors() *Window {
  if !terminal.IsTerminal() {
    fmt.Println("FATAL ERROR: Can't invoke GoCursors in non terminal context.")
    os.Exit(1)
  }
  width, height, err := terminal.GetTerminalSize()
  if err != nil {
    fmt.Println("FATAL ERROR: Root window initialization failed.")
    os.Exit(1)
  }
  EraseEntireScreen()
  return  &Window{
    CurX: 0,
    CurY: 0,
    Width: width,
    Height: height,
  }
}

func EraseEntireScreen() {
  fmt.Print("\033[2J")
}

func InitWindow(width, height int) *Window {
  return &Window{
    CurX: 0,
    CurY: 0,
    Width: width,
    Height: height,
  }
}

func (w *Window) GetChar() byte {
  return terminal.ReadByte()
}

func CbreakStart() *term.State {
  return terminal.MakeRaw()
}

func CbreakRestore(oldState *term.State) {
  terminal.ApplyState(oldState)
}
