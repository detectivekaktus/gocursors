package gocursors

import (
  "unicode"
)

const (
  CTRL_A = iota + 1
  CTRL_B
  CTRL_C
  CTRL_D
  CTRL_E
  CTRL_F
  CTRL_G
  CTRL_H
  CTRL_I
  CTRL_J
  CTRL_K
  CTRL_L
  CTRL_M
  CTRL_N
  CTRL_O
  CTRL_P
  CTRL_Q
  CTRL_R
  CTRL_S
  CTRL_T
  CTRL_U
  CTRL_V
  CTRL_W
  CTRL_X
  CTRL_Y
  CTRL_Z
)

/*
  Returns true if the got character is equal to expected character.
  The equality is checked even for the uppercase characters, so
  Compare('a', 'A') will be evaluated to true.
*/
func Compare(got, expected rune) bool {
  return unicode.ToUpper(got) == unicode.ToUpper(expected)
}
