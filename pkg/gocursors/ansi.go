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

/*
  Color type represents all the possible color values
  of the current terminal.
  It also stores all the values that are accessable with
  the terminals that allow using 8bit colors.
*/
type Color uint32 

/*
  Terminal defined foreground colors.
*/
const (
  FG_BLACK    Color = 30
  FG_RED      Color = 31
  FG_GREEN    Color = 32
  FG_YELLOW   Color = 33
  FG_BLUE     Color = 34
  FG_MAGENT   Color = 35
  FG_CYAN     Color = 36
  FG_WHITE    Color = 37
  FG_DEFAULT  Color = 39
)

/*
  Terminal defined background colors.
*/
const (
  BG_BLACK    Color = 40
  BG_RED      Color = 41
  BG_GREEN    Color = 42
  BG_YELLOW   Color = 43
  BG_BLUE     Color = 44
  BG_MAGENT   Color = 45
  BG_CYAN     Color = 46
  BG_WHITE    Color = 47
  BG_DEFAULT  Color = 49
)

/*
  Applies a color to the screen. Call to ResetAll to disable
  the colors after outputting the characters.
*/
func ApplyColor(clr Color) {
  fmt.Printf("\033[%dm", clr)
}

/*
  Applies one of the 256 possible colors. Call to ResetAll to
  disable the colors after outputting the characters.
*/
func Apply8bitColor(clr Color, foreground bool) {
  if foreground {
    fmt.Printf("\033[38;5;%dm", clr)
  } else {
    fmt.Printf("\033[48;5;%dm", clr)
  }
}

/*
  Applies an RGB color to the screen. Call to ResetAll to
  disable the colors after outputting the characters.
  The color must be passed a hex value.
*/
func ApplyRGBColor(clr Color, foreground bool) {
  if foreground {
    fmt.Printf("\033[38;2;%d;%d;%dm", (clr >> 16) & 0xFF, (clr >> 8) & 0xFF, clr & 0xFF)
  } else {
    fmt.Printf("\033[48;2;%d;%d;%dm", (clr >> 16) & 0xFF, (clr >> 8) & 0xFF, clr & 0xFF)
  }
}
