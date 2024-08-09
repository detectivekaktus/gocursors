package terminal

import (
	"bufio"
	"os"
)

func ReadByte() rune {
  reader := bufio.NewReader(os.Stdin)
  r, _, err := reader.ReadRune()
  if err != nil {
    os.Exit(1)
  }
  return r
}
