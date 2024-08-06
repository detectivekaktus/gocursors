# GoCursors
***WARNING! THIS PROJECT IS A WORK-IN-PROGRESS. DON'T EXPECT TOO MUCH FROM IT.***

GoCursors is an ncurses-like library written in Go for everyone who's tired of absurdly complex workflows, rules and interfaces that other popular Go TUI libraries have. It's important to notice that GoCursors is not the copy of ncurses with Go syntax but rather is a library inspired by ncurses and its simplicity.

If you have never used ncurses library or have never heard of it, check out [this on Wikipedia](https://en.wikipedia.org/wiki/Ncurses).

## Demonstration
Below you can find a program which enters into gocurses mode, prints *Hello, World!* message and exits after any key press.
```go
package main

import "github.com/detectivekaktus/gocursors/pkg/gocursors"

func main() {
  root := gocursors.GoCursors()
  oldState := gocursors.CbreakStart()

  defer gocursors.EndCursors(root)
  defer gocursors.CbreakRestore(oldState)

  root.OutString("Hello, World!")
  root.GetChar()
}
```

## Contributing
In order to contribute to the project, you must follow the [guidelines](https://github.com/detectivekaktus/GoCursors/blob/main/CONTRIBUTING.md). Any deviations won't probably be merged.
