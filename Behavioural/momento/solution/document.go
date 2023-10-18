package main

/*
Our Document class has three attributes:
- content
- fontName
- fontSize
We should allow the user to undo the changes to any of these
attributes. In the future, we may add additional attributes in this
class and these attributes should also be undoable.
Implement the undo feature using the memento pattern.
*/

type Document struct {
	content  string
	fontName string
	fontSize int
}

func (d *Document) CreateMemento() DocumentState {
	return SetState(d.content, d.fontName, d.fontSize)
}

func (d *Document) Restore(sate DocumentState) {
	d.content = sate.content
	d.fontName = sate.fontName
	d.fontSize = sate.fontSize
}
