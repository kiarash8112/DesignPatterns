package templatemethod

import "fmt"

/*
This class represents a window in a GUI framework. Application
developers can use this framework to build desktop applications.
The Window class has a method for closing a window. Certain
windows may need to execute some code before or after a window
is closed.
We cannot hardcode this behaviour in the Window class because
the code that needs to be executed is different from one window to
another.
Use the template method pattern to solve this problem.
*/

type Window struct {
}

func (Window) close() {
	fmt.Println("Removing the window from the screen")
}
