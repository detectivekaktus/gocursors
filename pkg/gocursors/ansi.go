package gocursors

import "fmt"

type Style uint8
const (
  BOLD          Style = 1
  DIM           Style = 2
  ITALIC        Style = 3
  UNDERLINE     Style = 4
  BLINK         Style = 5
  INVERSE       Style = 7
  HIDE          Style = 8
  STRIKETHROUGH Style = 9
)

func EraseEntireScreen() {
  fmt.Print("\033[2J")
}

func MoveCursor(y, x int) {
  fmt.Printf("\033[%d;%dH", y, x)
}

func ApplyStyle(stl Style) {
  fmt.Printf("\033[%dm", stl)
}

func ResetAll() {
  fmt.Print("\033[0m")
}
