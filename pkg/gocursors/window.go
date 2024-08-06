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
  w := &Window{
    CurX: 0,
    CurY: 0,
    Width: width,
    Height: height,
  }
  w.Home()
  return w
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

func (w *Window) Home() {
  w.Move(0, 0)
}

func (w *Window) Move(x, y int) {
  if x > w.Width || y > w.Height || x < 0 || y < 0 {
    return
  }
  w.CurX = x
  w.CurY = y
  fmt.Printf("\033[%d;%dH", w.CurY, w.CurX)
}

func (w *Window) MoveX(x int) {
  if x > w.Width || x < 0 {
    return
  }
  w.CurX = x
  fmt.Printf("\033[%d;%dH", w.CurY, w.CurX)
}

func (w *Window) MoveY(y int) {
  if y > w.Height || y < 0 {
    return
  }
  w.CurY = y
  fmt.Printf("\033[%d;%dH", w.CurY, w.CurX)
}

func (w *Window) CurAdd(x, y int) {
  if (w.CurX + x > w.Width || w.CurX + x < 0) || (w.CurY + y > w.Height || w.CurY + y < 0) {
    return
  }
  w.CurX += x
  w.CurY += y
  fmt.Printf("\033[%d;%dH", w.CurY, w.CurX)
}

func (w *Window) CurAddX(x int) {
  if w.CurX + x > w.Width || w.CurX + x < 0 {
    return
  }
  w.CurX += x
  fmt.Printf("\033[%d;%dH", w.CurY, w.CurX)
}

func (w *Window) CurAddY(y int) {
  if w.CurY + y > w.Height || w.CurY + y < 0 {
    return
  }
  w.CurY += y
  fmt.Printf("\033[%d;%dH", w.CurY, w.CurX)
}

func (w *Window) OutChar(b byte) {
  if w.CurX + 1 > w.Width {
    return
  }
  fmt.Printf("%c", b)
  w.CurX++
}

func (w *Window) OutString(s string) {
  fmt.Printf("%s", s)
  w.CurY += len(s) / w.Width
  w.CurX += len(s) - ((len(s) / w.Width) * w.Width)
}

func CbreakStart() *term.State {
  return terminal.MakeRaw()
}

func CbreakRestore(oldState *term.State) {
  terminal.ApplyState(oldState)
}
