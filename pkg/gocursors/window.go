package gocursors

import (
	"fmt"
	"os"

	"github.com/detectivekaktus/gocursors/internal/terminal"
	"golang.org/x/term"
)

var root *Window
var state *term.State

func GoCrash(msg string, args ...any) {
  if state != nil {
    terminal.ApplyState(state)
  }
  if root != nil {
    root.Home()
  }
  EraseEntireScreen()
  fmt.Printf(msg, args...)
  os.Exit(1)
}

type Window struct {
  StartX   int
  StartY   int

  CurX     int
  CurY     int

  Columns  int
  Rows     int

  Parent   *Window
  Children []*Window
}

func GoCursors() *Window {
  if !terminal.IsTerminal() {
    GoCrash("FATAL ERROR: Can't invoke GoCursors in non terminal context.\n")
  }
  columns, rows, err := terminal.GetTerminalSize()
  if err != nil {
    GoCrash("FATAL ERROR: Root window initialization failed.\n")
  }
  EraseEntireScreen()
  w := InitWindow(nil, columns, rows, 1, 1)
  w.Home()
  return w
}

func EndCursors(root *Window) {
  EraseEntireScreen()
  root.Home()
}

func InitWindow(parent *Window, columns, rows, startX, startY int) *Window {
  if parent == nil {
    if root != nil{
      GoCrash("FATAL ERROR: Can't create a root window when there's already an existing one.\n")
    }
  } else {
    if startX <= 0 || startY <= 0 {
      GoCrash("ERROR: Start coordinates begin from 1, 1.\n")
    }
    if columns <= 0 || rows <= 0 {
      GoCrash("ERROR: Window size cannot be equal or smaller than 0.\n")
    } else if columns >= parent.Columns || rows >= parent.Rows {
      GoCrash("ERROR: Child cannot be bigger or equal to its parent.\n")
    }
    w := &Window{
      Parent: parent,
      Children: make([]*Window, 0),

      StartX: parent.StartX + startX,
      StartY: parent.StartY + startY,

      Columns: columns,
      Rows: rows,

      CurX: 1,
      CurY: 1,
    }
    parent.Children = append(parent.Children, w)
    w.Home()
    return w
  }
  root = &Window{
    Parent: parent,
    Children: make([]*Window, 0),

    StartX: startX,
    StartY: startY,

    Columns: columns,
    Rows: rows,

    CurX: 1,
    CurY: 1,
  }
  root.Home()
  return root
}

func (w *Window) GetChar() byte {
  return terminal.ReadByte()
}

func (w *Window) Cursor() {
  w.Move(w.StartX + w.CurX - 1, w.StartY + w.CurY - 1)
}

func (w *Window) Home() {
  w.Move(w.StartX, w.StartY)
}

func (w *Window) Move(x, y int) {
  if x > w.StartX + w.Columns || y > w.StartY + w.Rows || x <= 0 || y <= 0 {
    return
  }
  w.CurX = x
  w.CurY = y
  MoveCursor(w.CurY, w.CurX)
}

func (w *Window) MoveX(x int) {
  if x > w.StartX + w.Columns || x <= 0 {
    return
  }
  w.Move(x, w.CurY)
}

func (w *Window) MoveY(y int) {
  if y > w.StartY + w.Rows || y <= 0 {
    return
  }
  w.Move(w.CurX, y)
}

func (w *Window) CurAdd(x, y int) {
  if (w.CurX + x > w.StartX + w.Columns || w.CurX + x <= 0) ||
    (w.CurY + y > w.StartY + w.Rows || w.CurY + y <= 0) {
    return
  }
  w.CurX += x
  w.CurY += y
  MoveCursor(w.CurY, w.CurX)
}

func (w *Window) CurAddX(x int) {
  if w.CurX + x > w.StartX + w.Columns || w.CurX + x <= 0 {
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

func (w *Window) OutChar(r rune) {
  if w.CurX + 1 > w.StartX + w.Columns {
    w.CurAdd(-w.Columns, 1)
  }
  fmt.Printf("%c", r)
  w.CurX++
}

func (w *Window) OutString(s string) {
  for i := 0; i < len(s); i++ {
    if s[i] == '\n' {
      w.CurAdd(-w.CurX + 1, 1)
    } else {
      w.OutChar(rune(s[i]))
    }
  }
}

func (w *Window) OutFormat(s string, args ...any) {
  w.OutString(fmt.Sprintf(s, args...))
}

func (w *Window) Border() {
  w.CustomBorder('┏', '┓', '┗', '┛', '━', '┃')
}

func (w *Window) CustomBorder(topLeft, topRight, bottomLeft, bottomRight, horizontal, vertical rune) {
  for y := 0; y < w.Rows; y++ {
    for x := 0; x < w.Columns; x++ {
      w.Move(w.StartX + x, w.StartY + y)
      if x == 0 && y == 0 {
        w.OutChar(topLeft)
      } else if x == w.Columns - 1 && y == 0 {
        w.OutChar(topRight)
      } else if x == 0 && y == w.Rows - 1 {
        w.OutChar(bottomLeft)
      } else if x == w.Columns - 1 && y == w.Rows - 1 {
        w.OutChar(bottomRight)
      } else if y == 0 || y == w.Rows - 1 {
        w.OutChar(horizontal)
      } else if x == 0 || x == w.Columns - 1 {
        w.OutChar(vertical)
      }
    }
  }
  w.Move(w.StartX + 1, w.StartY + 1)
}

func CbreakStart() {
  state = terminal.MakeRaw()
}

func CbreakRestore() {
  terminal.ApplyState(state)
}
