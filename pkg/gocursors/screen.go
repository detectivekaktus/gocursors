package gocursors

import (
	"os"

	"github.com/detectivekaktus/gocursors/internal/terminal"
)

type ResizeFunc func(*Window)

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
