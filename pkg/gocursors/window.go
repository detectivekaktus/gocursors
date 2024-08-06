package gocursors

import (
	"fmt"
	"os"

	"github.com/detectivekaktus/gocursors/internal/terminal"
)

type Window struct {
  CurX          int
  CurY          int
  Width         int
  Height        int

  KeyPadEnabled bool
  EchoEnabled   bool
}

func GoCursors() *Window {
  if !terminal.IsTerminal() {
    fmt.Println("FATAL ERROR: Cannot invoke GoCursors in non terminal context.")
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
    KeyPadEnabled: false,
    EchoEnabled: false,
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
    KeyPadEnabled: false,
    EchoEnabled: false,
  }
}

func (w *Window) EnableKeyPad(b bool) {
  w.KeyPadEnabled = b
}

func (w *Window) EnableEcho(b bool) {
  w.EchoEnabled = b
}
