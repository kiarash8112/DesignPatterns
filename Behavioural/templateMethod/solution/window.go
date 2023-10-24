package solution

type closeWindow interface {
	close()
	jobs()
}

type Window struct {
	exit closeWindow
}

func (w Window) close() {
	w.exit.jobs()
	w.exit.close()
}
