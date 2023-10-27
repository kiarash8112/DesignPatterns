package mediator

/*
We’re using a third-party GUI framework for building an application.
We need to build a dialog box for a new user to sign up. On this
dialog box we need three UI elements:
- A text box to enter a username
- A text box to enter a password
- A check box to agree with the terms
- A sign up button
The sign up button is only enabled if both text boxes are filled out
and the check box is checked.
Use the mediator pattern with the observer pattern to implement the
coordination between these elements in a class called
SignUpDialogBox.
*/

type UIControl struct {
}

type TextBox struct {
	content string
}

type CheckBox struct {
	isChecked bool
}

type Button struct {
	isEnabled bool
}
