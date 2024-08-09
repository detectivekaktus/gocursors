package gocursors

import "fmt"

/*
  Style type represents all the possible formatting
  styles that can be applied to the text.
*/
type Style uint8

/*
  The list of all the possible styles that can be applied
  with ApplyStyle.
*/
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

/*
  Deletes all the characters ever outputted to the screen buffer.
*/
func EraseEntireScreen() {
  fmt.Print("\033[2J")
}

/*
  Moves the cursor to a specific position on the screen.
  Consider using Window methods Move, MoveX, and MoveY instead
  as they provide more safety over the cursor position.
*/
func MoveCursor(y, x int) {
  fmt.Printf("\033[%d;%dH", y, x)
}

/*
  Applies a Style passed as an argument.
  Call to the ResetAll function to disable all the styles applied
  before.
*/
func ApplyStyle(stl Style) {
  fmt.Printf("\033[%dm", stl)
}

/*
  Resets all the formatting Styles and all the Colors ever applied
  before calling the function.
*/
func ResetAll() {
  fmt.Print("\033[0m")
}
