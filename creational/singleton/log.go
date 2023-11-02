package singleton

import "fmt"

/*
With the current implementation, we can create multiple loggers writing to the
same log le in parallel.
Use the singleton pattern to ensure only a single logger can be instantiated
for a given le.

*/

type Logger struct {
	filename string
}

func (Logger) write(msg string) {
	fmt.Println("writing message to the log")
}
