package terminal

import (
	"fmt"
	"os"
)

func ReadByte() byte {
  buf := make([]byte, 1)
  _, err := os.Stdin.Read(buf)
  if err != nil {
    fmt.Println("ERROR: Couldn't read the byte at stdin stream to the buffer.")
    os.Exit(1)
  }
  return buf[0]
}
