package decorator

/*

We’re going to build a code editor like IntelliJ or VSCode. When we open a
project, we should see all the artefacts (items) in that project.
Every artefact should have an icon and the icon can be different from one
artefact to another.

Similarly, if an artefact includes errors, its icon should include an error marker:

look at the current implementation of these concepts in the decorator
package.
Read my comments in the Artefact struct about the problems with this
implementation.
Refactor the solution using the decorator pattern.
*/

type Artefact struct {
	// The current implementation is not easily extensible. If tomorrow we need
	// to support other special markers, we have to come back and modify this struct.
	//
	// For example, if the project is under source control, we need to highlight
	// the artefacts that are changed but not committed to the repository with a
	// special marker.
	//
	// Similarly, if an artefact is excluded from the project, we may want to
	// make the icon look disabled or semi-transparent.
	//
	// Every time we need to support a new marker, we have to come back and modify
	// this struct. Over time, the code in this package gets more and more convoluted
	// with several if statements and additional fields.
	name     string
	hasError bool
	isMain   bool
}

func (a *Artefact) render() string {
	res := ""
	if a.hasError {
		res += "hasError"
	}
	if a.isMain {
		res += "isMain"
	}
	return res
}

type editor struct {
}

func (editor) openProject(path string) {
	artefacts := []Artefact{
		{name: "Main"},
		{name: "Demo"},
		{name: "EmailClient"},
		{name: "EmailProvider"},
	}

	artefacts[0].isMain = true
	artefacts[2].hasError = true

	for _, arartefact := range artefacts {
		arartefact.render()
	}
}
