package command

/*
This class represents the operations we can perform on a video,
such as adding a label to it or changing its contrast.
We need to allow the user to undo any changes he/she makes to a
video.
What pattern would you use to implement this feature? The
memento or the command pattern? Why?
Write the necessary code to implement the undo feature.
*/

//if we use memento pattern it takes snapshot of each stage
//then we will lost lot of memory space

type VideoEditor struct {
	contrast float32
	text     string
}
