package main

type VideoEditor struct {
	contrast float32
	text     string
}

type command interface {
	execute()
}

type undoableCommand interface {
	command
	unexecute()
}

type history struct {
	index    int
	commands [20]undoableCommand
}

func (h *history) push(command undoableCommand) {
	h.commands[h.index] = command
	h.index++
	//also we should check if index > 20(size of h.commands) and return error
}

func (h *history) pop() undoableCommand {
	h.index--
	res := h.commands[h.index]
	return res
}

type SetContrastCommand struct {
	editor      *VideoEditor
	history     *history
	contrast    float32
	preContrast float32
}

func (s *SetContrastCommand) execute() {
	s.preContrast = s.editor.contrast
	s.editor.contrast = s.contrast
	s.history.push(s)
}

func (s *SetContrastCommand) unexecute() {
	s.editor.contrast = s.preContrast
}

type SetTextCommand struct {
	editor  *VideoEditor
	history *history
	text    string
	preText string
}

func (s *SetTextCommand) execute() {
	s.preText = s.editor.text
	s.editor.text = s.text
	s.history.push(s)
}

func (s *SetTextCommand) unexecute() {
	s.editor.text = s.preText
}

type UndoCommand struct {
	history *history
}

func (u *UndoCommand) execute() {
	if u.history.index > 0 {
		u.history.pop().unexecute()
	}
}
