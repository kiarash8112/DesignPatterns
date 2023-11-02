package prototype

/*
You’re building a video editor similar to Adobe Premier. The editor contains a
timeline of various types of components such as text, clips, audio, and so on.
The user should be able to duplicate any component. The duplicated
component should be added to the timeline.

What are the problems in the current implementation?
Refactor the code using the prototype pattern. What have you achieved?
*/

type Component interface {
}

type Text struct {
	content string
}

type Audio struct {
}

type Clip struct {
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
	if source, ok := component.(Text); ok {
		newText := Text{content: source.content}
		c.timeline.add(newText)
	} else if _, ok := component.(Clip); ok {
		// Logic for duplicating a Clip object
	} else if _, ok := component.(Audio); ok {
		// Logic for duplicating a Audio object
	}
}
