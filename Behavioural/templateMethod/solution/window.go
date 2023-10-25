package solution

import "fmt"

type closeWindow interface {
	close()
	jobs()
}

type Window struct {
	exit closeWindow
}

func (w Window) DoJobAndClose() {
	w.exit.jobs()
	w.exit.close()
}

type CustomWindow struct {
	Window
}

func (CustomWindow) jobs() {
	fmt.Println("do some jobs before closing window")
}

func (CustomWindow) close() {
	fmt.Println("closing the window!")
}
