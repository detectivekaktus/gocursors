package gocursors

import "fmt"

func EraseEntireScreen() {
  fmt.Print("\033[2J")
}

func MoveCursor(y, x int) {
  fmt.Printf("\033[%d;%dH", y, x)
}

func Bold() {
  fmt.Print("\033[1m")
}

func Dim() {
  fmt.Print("\033[2m")
}

func Italic() {
  fmt.Print("\033[3m")
}

func Underline() {
  fmt.Print("\033[4m")
}

func Blink() {
  fmt.Print("\033[5m")
}

func Inverse() {
  fmt.Print("\033[7m")
}

func Hide() {
  fmt.Print("\033[8m")
}

func Strikethrough() {
  fmt.Print("\033[9m")
}

func ResetAll() {
  fmt.Print("\033[0m")
}
