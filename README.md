# GoCursors
GoCursors is an ncurses-like library written in Go for everyone who's tired of absurdly complex workflows, rules and interfaces that other popular Go TUI libraries have. It's important to notice that GoCursors is not the copy of ncurses with Go syntax but rather is a library inspired by ncurses and its simplicity.

Any knowledge of ncurses is welcome and can help a lot while exploring the features of the library.

In order to use this library successfully you need a terminal or a terminal emulator which supports ANSI escape codes. To install the library, use `go get github.com/detectivekaktus/gocursors`.

## Demonstration
- Hello world. The program that prints *Hello, World!* and waits for the user input and then exits normally.
```go
package main

import "github.com/detectivekaktus/gocursors/pkg/gocursors"

func main() {
  root := gocursors.GoCursors()
  gocursors.CbreakStart()

  defer gocursors.CbreakRestore()
  defer gocursors.EndCursors()

  root.OutString("Hello, World!")
  root.GetChar()
}
```

- Text editor. The program has a border around and can output and clear the characters. Once ESCAPE is pressed, the program exits normally.
```go
package main

import "github.com/detectivekaktus/gocursors/pkg/gocursors"

func main() {
  root := gocursors.GoCursors()
  gocursors.CbreakStart()

  defer gocursors.CbreakRestore()
  defer gocursors.EndCursors()

  root.Border()
  for {
    c := root.GetChar()
    if c == 27 { // 27 is ESCAPE character.
      break
    } else if c == 8 || c == 127 || c == '\b'{ // 8, 127 and '\b' are BACKSPACE characters.
      root.CurAdd(-1, 0)
      root.OutChar(' ')
      root.CurAdd(-1, 0)
    } else {
      root.OutChar(c)
    }
  }
}
```

## Contributing
In order to contribute to the project, you must follow the [guidelines](https://github.com/detectivekaktus/GoCursors/blob/main/CONTRIBUTING.md). Any deviations won't probably be merged.
