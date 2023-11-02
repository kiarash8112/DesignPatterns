package solution

type Component interface {
	clone() Component
}

type Text struct {
	content string
}

func (t *Text) clone() Component {
	return &Text{content: t.content}
}

type Audio struct {
}

func (Audio) clone() Component {
	return Audio{}
}

type Clip struct {
}

func (Clip) clone() Component {
	return Clip{}
}

type TimeLine struct {
	components []Component
}

func (t *TimeLine) add(component Component) {
	t.components = append(t.components, component)
}

type ContextMenu struct {
	timeline TimeLine
}

func (c *ContextMenu) duplicate(component Component) {
	newComponent := component.clone()
	c.timeline.add(newComponent)
}
