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

func EraseEntireScreen() {
  fmt.Print("\033[2J")
}

func InitWindow(parent *Window, columns, rows, startX, startY int) *Window {
  if parent == nil {
    if root != nil{
      GoCrash("FATAL ERROR: Can't create a root window when there's already an existing one.\n")
    }
  } else {
    if startX <= 0 || startY <= 0 {
      GoCrash("ERROR: Start coordinates begin from 1, 1.\n")
    } else if startX > parent.Columns + columns || startY > parent.Rows + rows {
      GoCrash("ERROR: Start coordinates must be within the window.\n")
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
  }
  w := &Window{
    Parent: parent,
    Children: make([]*Window, 0),

    StartX: startX,
    StartY: startY,

    Columns: columns,
    Rows: rows,

    CurX: 1,
    CurY: 1,
  }
  root = w
  return w
}

func (w *Window) GetChar() byte {
  return terminal.ReadByte()
}

func (w *Window) Cursor() {
  w.Move(w.StartX + w.CurX - 1, w.StartY + w.CurY - 1)
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
  w.Cursor()
  for i := 0; i < len(s); i++ {
    if s[i] == '\n' {
      w.CurAdd(-w.CurX + 1, 1)
    } else {
      w.OutChar(s[i])
    }
  }
}

func (w *Window) OutFormat(s string, args ...any) {
  w.Cursor()
  w.OutString(fmt.Sprintf(s, args...))
}

func CbreakStart() {
  state = terminal.MakeRaw()
}

func CbreakRestore() {
  terminal.ApplyState(state)
}
