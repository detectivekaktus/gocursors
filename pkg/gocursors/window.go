package gocursors

import (
	"fmt"
	"os"

	"github.com/detectivekaktus/gocursors/internal/terminal"
	"golang.org/x/term"
)

var root *Window
var state *term.State

/*
  Resets the terminal state if set to raw mode, moves the
  cursor position to 1, 1 and exits the program with 1 exit code.

  You should always use this function whenever you want you program
  exit abnormally, otherwise you may get unexpected behaviour once
  the program has finished the exectuion.
*/
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

/*
  The Window struct represents a portion of a terminal buffer that can
  be used for input and output operations. It defines a specific region within
  the terminal where text and other data can be displayed or manipulated.

  Note that StartX and StartY fields are relative to the root window.
*/
type Window struct {
  StartX      int
  StartY      int

  CurX        int
  CurY        int

  Columns     int
  Rows        int

  Parent      *Window
  Children  []*Window
}

/*
  Enables the GoCursors mode which erases the entire screen and creates
  the root window which is returned after the function call.

  The root window is the main window of the GoCursors program which holds
  all other windows and handles the screen resize events. The root window
  start coordinates are at 1, 1 (top left corner of the screen).

  To finish the program execution normally, see the EndCursors
  function.
  To finish the program execution abnormally, see the GoCrash
  function.
*/
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

/*
  Disables the GoCursors mode, erases the entire screen and moves the cursor
  at 1, 1 (top left corner of the screen).
*/
func EndCursors() {
  EraseEntireScreen()
  root.Home()
}

/*
  Initializes a new Window on the terminal buffer and returns it.

  Any errors that occur during the window initializations are going to crash
  the program and display the error message.

  If there's no windows created, the root window must be passed as parent
  instead.
*/
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

/*
  Reads an ASCII character from the user and returns it back.
*/
func (w *Window) GetChar() byte {
  return terminal.ReadByte()
}

/*
  Moves the cursor to the window's CurX and CurY.
*/
func (w *Window) Cursor() {
  w.Move(w.StartX + w.CurX - 1, w.StartY + w.CurY - 1)
}

/*
  Moves the cursor to the start of the window.
*/
func (w *Window) Home() {
  w.Move(w.StartX, w.StartY)
}

/*
  Moves the cursor to a specified position within the window.
  If you want to increment or decrement x or y values, use CurAdd instead.
*/
func (w *Window) Move(x, y int) {
  if x > w.StartX + w.Columns || y > w.StartY + w.Rows || x <= 0 || y <= 0 {
    return
  }
  w.CurX = x
  w.CurY = y
  MoveCursor(w.CurY, w.CurX)
}

/*
  The same as Move. Expects only the x value.
*/
func (w *Window) MoveX(x int) {
  if x > w.StartX + w.Columns || x <= 0 {
    return
  }
  w.Move(x, w.CurY)
}

/*
  The same as Move. Expects only the y value.
*/
func (w *Window) MoveY(y int) {
  if y > w.StartY + w.Rows || y <= 0 {
    return
  }
  w.Move(w.CurX, y)
}

/*
  Adds values to the current cursor position within the Window.
  If you want to jump to a specific position on the screen, use Move.
*/
func (w *Window) CurAdd(x, y int) {
  if (w.CurX + x > w.StartX + w.Columns || w.CurX + x <= 0) ||
    (w.CurY + y > w.StartY + w.Rows || w.CurY + y <= 0) {
    return
  }
  w.CurX += x
  w.CurY += y
  MoveCursor(w.CurY, w.CurX)
}

/*
  The same as CurAdd. Expects only the x value.
*/
func (w *Window) CurAddX(x int) {
  if w.CurX + x > w.StartX + w.Columns || w.CurX + x <= 0 {
    return
  }
  w.CurAdd(x, 0)
}

/*
  The same as CurAdd. Expects only the y value.
*/
func (w *Window) CurAddY(y int) {
  if w.CurY + y > w.Rows || w.CurY + y < 0 {
    return
  }
  w.CurAdd(0, y)
}

/*
  Outputs a character to the screen. If the current cursor
  position is at the end of the line, jumps to the next line.
  If the cursor position is at the end of the window, nothing
  is printed.
*/
func (w *Window) OutChar(r rune) {
  if w.CurX + 1 > w.StartX + w.Columns {
    w.CurAdd(-w.Columns, 1)
  }
  fmt.Printf("%c", r)
  w.CurX++
}

/*
  Outputs a string to the screen. If the string can't fit
  into a single line, it will be split in different lines.
  If the cursor position is at the end of the window, nothing
  is printed.
*/
func (w *Window) OutString(s string) {
  for i := 0; i < len(s); i++ {
    if s[i] == '\n' {
      w.CurAdd(-w.CurX + 1, 1)
    } else {
      w.OutChar(rune(s[i]))
    }
  }
}

/*
  Outputs a formatted string to the screen. If the string can't
  fit into a single line, it will be split in different lines.
  If the cursor position is at the end of the window, nothing
  is printed.
*/
func (w *Window) OutFormat(s string, args ...any) {
  w.OutString(fmt.Sprintf(s, args...))
}

/*
  Creates a standard border around the window and moves the cursor
  to the coordinates 2, 2.
  If you want to create a custom border, consider using CustomBorder.
*/
func (w *Window) Border() {
  w.CustomBorder('┏', '┓', '┗', '┛', '━', '┃')
}

/*
  Creates a custom border around the window and moves the cursor to
  the coordinates 2, 2.
*/
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

/*
  Makes the terminal to work in raw mode which enables to read characters
  immediately as they they are inputted.
  
  If you use this function, you must end the program with CbreakRestore
  function call, otherwise you may get unexpected behaviour in your terminal
  after your program finished the exectuion.
*/
func CbreakStart() {
  state = terminal.MakeRaw()
}

/*
  Restores the terminal state to normal one.
*/
func CbreakRestore() {
  terminal.ApplyState(state)
}
