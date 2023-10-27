package main

import "fmt"

func main() {
	editor := VideoEditor{contrast: 12, text: "init"}
	history := history{commands: [20]undoableCommand{}}

	setText := SetTextCommand{editor: &editor, history: &history, text: "newText"}
	setText.execute()

	setContrast := SetContrastCommand{editor: &editor, history: &history, contrast: 15}
	setContrast.execute()

	undo := UndoCommand{history: &history}
	undo.execute()
	undo.execute()

	fmt.Printf("editor: %v\n", editor)
}
