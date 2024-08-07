package gocursors

import (
	"fmt"
	"os"

	"github.com/detectivekaktus/gocursors/internal/terminal"
	"golang.org/x/term"
)

type Window struct {
  CurX    int
  CurY    int
  Columns int
  Rows    int
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
  w := InitWindow(width, height)
  w.Home()
  return w
}

func EndCursors(root *Window) {
  EraseEntireScreen()
  root.Home()
}

func EraseEntireScreen() {
  fmt.Print("\033[2J")
}

func InitWindow(width, height int) *Window {
  return &Window{
    CurX: 1,
    CurY: 1,
    Columns: width + 1,
    Rows: height + 1,
  }
}

func (w *Window) GetChar() byte {
  return terminal.ReadByte()
}

func (w *Window) Home() {
  w.Move(1, 1)
}

func (w *Window) Move(x, y int) {
  if x > w.Columns || y > w.Rows || x < 0 || y < 0 {
    return
  }
  w.CurX = x
  w.CurY = y
  fmt.Printf("\033[%d;%dH", w.CurY, w.CurX)
}

func (w *Window) MoveX(x int) {
  if x > w.Columns || x < 0 {
    return
  }
  w.Move(x, w.CurY)
}

func (w *Window) MoveY(y int) {
  if y > w.Rows || y < 0 {
    return
  }
  w.Move(w.CurX, y)
}

func (w *Window) CurAdd(x, y int) {
  if (w.CurX + x > w.Columns || w.CurX + x < 0) || (w.CurY + y > w.Rows || w.CurY + y < 0) {
    return
  }
  w.CurX += x
  w.CurY += y
  fmt.Printf("\033[%d;%dH", w.CurY, w.CurX)
}

func (w *Window) CurAddX(x int) {
  if w.CurX + x > w.Columns || w.CurX + x < 0 {
    return
  }
  w.CurAdd(x, 0)
}

func (w *Window) CurAddY(y int) {
  if w.CurY + y > w.Rows || w.CurY + y < 0 {
    return
  }
  w.CurAdd(0, y)
}

func (w *Window) OutChar(b byte) {
  if w.CurX + 1 > w.Columns {
    w.CurAdd(-w.CurX + 1, 1)
  }
  fmt.Printf("%c", b)
  w.CurX++
}

func (w *Window) OutString(s string) {
  for i := 0; i < len(s); i++ {
    if s[i] == '\n' {
      w.CurAdd(-w.CurX + 1, 1)
    } else {
      w.OutChar(s[i])
    }
  }
}

func CbreakStart() *term.State {
  return terminal.MakeRaw()
}

func CbreakRestore(oldState *term.State) {
  terminal.ApplyState(oldState)
}
