package solution

type AbstractArtefact interface {
	render() string
}

var _ AbstractArtefact = &Artefact{}

type Artefact struct {
	name string
}

func (a Artefact) render() string {
	return a.name
}

type ErrorDecorator struct {
	AbstractArtefact
}

func (e ErrorDecorator) render() string {
	return "[Error]" + e.render()
}

type MainDecorator struct {
	AbstractArtefact
}

func (m MainDecorator) render() string {
	return "[Main]" + m.render()
}

type editor struct {
}

func (editor) openProject(path string) {
	artefacts := []AbstractArtefact{
		Artefact{name: "Main"},
		Artefact{name: "Demo"},
		Artefact{name: "EmailClient"},
		Artefact{name: "EmailProvider"},
	}
	artefacts[0] = MainDecorator{AbstractArtefact: ErrorDecorator{AbstractArtefact: artefacts[0]}}
	artefacts[2] = ErrorDecorator{AbstractArtefact: artefacts[2]}
	for _, arartefact := range artefacts {
		arartefact.render()
	}
}
