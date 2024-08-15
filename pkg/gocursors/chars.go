package gocursors

import (
  "unicode"
)

/*
  Returns true if the got character is equal to expected character.
  The equality is checked even for the uppercase characters, so
  Compare('a', 'A') will be evaluated to true.
*/
func Compare(got, expected rune) bool {
  return unicode.ToUpper(got) == unicode.ToUpper(expected)
}
