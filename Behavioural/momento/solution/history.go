package main

type DocumentState struct {
	content  string
	fontName string
	fontSize int
}

func SetState(content, fontName string, fontSize int) DocumentState {
	return DocumentState{
		content:  content,
		fontName: fontName,
		fontSize: fontSize,
	}
}

type DocumentSateHistory struct {
	size    int
	history [10]DocumentState
}

func (d *DocumentSateHistory) Push(document DocumentState) {
	d.size++
	d.history[d.size] = document
}

func (d *DocumentSateHistory) Pop() DocumentState {
	if d.size <= 0 {
		return DocumentState{}
	}
	d.size--
	return d.history[d.size]
}
