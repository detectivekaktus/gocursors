/*
  Package that enables ncurses-like screen manipulations.

  The library is based on the ANSI escape codes and to perform
  the most basic operations ncurses library has. Note that the
  library is not an ncurses wrapper or binding, it's just
  inspired by it.
*/
package gocursors

import (
  "os"
  "fmt"

  "github.com/detectivekaktus/gocursors/internal/terminal"
  "golang.org/x/term"
)

var state *term.State

/*
  Resets the terminal state if set to raw mode, moves the
  cursor position to 1, 1 and exits the program with 1 exit code.

  You should always use this function whenever you want your program to
  exit with a followup crash message.
*/
func GoCrash(msg string, args ...any) {
  if state != nil {
    terminal.ApplyState(state)
  }
  if Root != nil {
    Root.Home()
  }
  EraseEntireScreen()
  fmt.Printf(msg, args...)
  os.Exit(1)
}

/*
  Resets the terminal state if set to raw mode, moves the
  cursor position to 1, 1 and exits the program with exit code specified.
  
  You should always use this function whenever you want you program to exit
  without any followup message.
*/
func GoExit(exitCode int) {
  if state != nil {
    terminal.ApplyState(state)
  }
  if Root != nil {
    Root.Home()
  }
  EraseEntireScreen()
  os.Exit(exitCode)
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
  Root.Home()
}

/*
  Makes the terminal to work in raw mode which enables to read characters
  immediately as they they are inputted.
  
  If you use this function, you must end the program with CbreakRestore
  function call, otherwise you may get unexpected behaviour in your terminal
  after your program has finished the exectuion.
*/
func CbreakStart() {
  state = terminal.MakeRaw()
}

/*
  Restores the terminal state to the normal one.
*/
func CbreakRestore() {
  terminal.ApplyState(state)
}
