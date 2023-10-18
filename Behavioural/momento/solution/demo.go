package main

import "fmt"

func main() {
	document := Document{content: "hello", fontName: "name", fontSize: 10}
	history := DocumentSateHistory{size: 0, history: [10]DocumentState{}}
	history.Push(document.CreateMemento())

	document.fontName = "name1"
	history.Push(document.CreateMemento())

	document.Restore(history.Pop())
	fmt.Printf("document: %v\n", document)
}
