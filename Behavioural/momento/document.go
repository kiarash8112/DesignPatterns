package momento

import "fmt"

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

func (d Document) toString() string {
	return fmt.Sprint("content", d.content, "fontName",
		d.fontName, "fontSize", d.fontSize)
}
