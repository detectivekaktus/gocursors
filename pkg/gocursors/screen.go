package gocursors

import (
	"os"

	"github.com/detectivekaktus/gocursors/internal/terminal"
)

/*
  A function that handles the resize logic must
  the following signature.
*/
type ResizeFunc func(*Window)

/*
  Updates the root window and calls to the ResizeFunc to
  handle the terminal size changes.
*/
func Resize(root *Window, f ResizeFunc) {
  columns, rows, err := terminal.GetTerminalSize()
  if err != nil {
    GoCrash("Unable to get the terminal size. Quiting...")
    os.Exit(1)
  }
  if root.Columns != columns || root.Rows != rows {
    root.Columns = columns
    root.Rows = rows
    f(root)
  }
}
