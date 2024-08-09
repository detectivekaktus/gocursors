package gocursors

import "fmt"

type Color uint8
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

func ApplyColor(clr Color) {
  fmt.Printf("\033[%dm", clr)
}

func Apply8bitColor(val uint8, foreground bool) {
  if foreground {
    fmt.Printf("\033[38;5;%dm", val)
  } else {
    fmt.Printf("\033[48;5;%dm", val)
  }
}

func ApplyRGBColor(val uint32, foreground bool) {
  if foreground {
    fmt.Printf("\033[38;2;%d;%d;%dm", (val >> 16) & 0xFF, (val >> 8) & 0xFF, val & 0xFF)
  } else {
    fmt.Printf("\033[48;2;%d;%d;%dm", (val >> 16) & 0xFF, (val >> 8) & 0xFF, val & 0xFF)
  }
}
