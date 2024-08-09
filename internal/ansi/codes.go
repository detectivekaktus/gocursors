package ansi

import "fmt"

func EraseEntireScreen() {
  fmt.Print("\033[2J")
}

func MoveCursor(y, x int) {
  fmt.Printf("\033[%d;%dH", y, x)
}
