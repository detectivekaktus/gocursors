package gocursors

import (
	"fmt"

	"github.com/detectivekaktus/gocursors/internal/terminal"
)

var Root *Window

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

  hasBorder   bool
  fgColor     Color
  bgColor     Color
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
    if Root != nil{
      GoCrash("FATAL ERROR: Can't create a root window when there's already an existing one.\n")
    }
  } else {
    if startX < 0 || startY < 0 {
      GoCrash("ERROR: Start coordinates begin from 1, 1.\n")
    }
    if columns <= 0 || rows <= 0 {
      GoCrash("ERROR: Window size cannot be equal or smaller than 0.\n")
    } else if columns > parent.Columns || rows > parent.Rows {
      GoCrash("ERROR: Child cannot be bigger than its parent.\n")
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

      hasBorder: false,
    }
    parent.Children = append(parent.Children, w)
    w.Home()
    return w
  }
  Root = &Window{
    Parent: parent,
    Children: make([]*Window, 0),

    StartX: startX,
    StartY: startY,

    Columns: columns,
    Rows: rows,

    CurX: 1,
    CurY: 1,

    hasBorder: false,
  }
  Root.Home()
  return Root
}

/*
  Reads an ASCII character from the user and returns it back.
*/
func (w *Window) GetChar() rune {
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
  if w.CurX == w.StartX + w.Columns && w.CurY == w.StartY + w.Rows {
    return
  }
  if w.bgColor != 0 {
    if w.bgColor < 256 {
      Apply8bitColor(w.bgColor, false)
    } else {
      ApplyRGBColor(w.bgColor, false)
    }
  }
  if w.fgColor != 0 {
    if w.fgColor < 256 {
      Apply8bitColor(w.fgColor, true)
    } else {
      ApplyRGBColor(w.fgColor, true)
    }
  }
  if w.hasBorder && (w.CurX + 1 == w.StartX + w.Columns) {
    w.CurAdd(-w.Columns + 2, 1)
  } else if w.CurX + 1 > w.StartX + w.Columns {
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
  to the coordinates 2, 2. Drawing a border erases all the characters
  that were previously present on the screen.
  If you want to create a custom border, consider using CustomBorder.
*/
func (w *Window) Border() {
  w.CustomBorder('┏', '┓', '┗', '┛', '━', '┃')
}

/*
  Creates a custom border around the window and moves the cursor to
  the coordinates 2, 2. Drawing a border erases all the characters
  that were previously present on the screen.

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
  w.hasBorder = true
  w.Move(w.StartX + 1, w.StartY + 1)
}

/*
  Clears the content of the Window.
*/
func (w *Window) Erase() {
  w.Home()
  w.hasBorder = false
  for y := 0; y < w.Rows; y++ {
    for x := 0; x < w.Columns; x++ {
      w.OutChar(' ')
    }
  }
  w.Home()
}

/*
  Sets the default background color for the Window and redraws it.
  Note that all the content previously rendered will disappear.

  For setting the terminal defined values, use 8-bit color values
  you can get here: https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797
*/
func (w *Window) SetBackgroundColor(clr Color) {
  w.bgColor = clr
  w.Erase()
}

/*
  Sets the default foreground color for the Window and redraws it.
  Note that all the content previously rendered will disappear.

  For setting the terminal defined values, use 8-bit color values
  you can get here: https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797
*/
func (w *Window) SetForegroundColor(clr Color) {
  w.fgColor = clr
  w.Erase()
}
