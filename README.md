# GoCursors
***WARNING! THIS PROJECT IS A WORK-IN-PROGRESS. DON'T EXPECT TOO MUCH FROM IT.***

GoCursors is an ncurses-like library written in Go for everyone who's tired of absurdly complex workflows, rules and interfaces that other popular Go TUI libraries have. It's important to notice that GoCursors is not the copy of ncurses with Go syntax but rather is a library inspired by ncurses and its simplicity.

If you have never used ncurses library or have never heard of it, check out [this on Wikipedia](https://en.wikipedia.org/wiki/Ncurses).

In order to use this library successfully you need a terminal or a terminal emulator which supports ANSI escape codes.

## Demonstration
Below you can find a program which enters into GoCursors mode, prints *Hello, World!* message at the coordinates 1, 1 (top left corner of the screen) and exits after any key press.
```go
package main

import "github.com/detectivekaktus/gocursors"

func main() {
  root := gocursors.GoCursors()
  oldState := gocursors.CbreakStart()

  defer gocursors.CbreakRestore(oldState)
  defer gocursors.EndCursors(root)

  root.OutString("Hello, World!")
  root.GetChar()
}
```

## Contributing
In order to contribute to the project, you must follow the [guidelines](https://github.com/detectivekaktus/GoCursors/blob/main/CONTRIBUTING.md). Any deviations won't probably be merged.
