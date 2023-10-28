package main

type observer interface {
	update()
}

type UIControl struct {
	observers []observer
}

func (u *UIControl) addObserver(observer observer) {
	u.observers = append(u.observers, observer)
}

func (u UIControl) notifyObservers() {
	for _, observer := range u.observers {
		observer.update()
	}
}

type TextBox struct {
	UIControl
	content string
}

func (t *TextBox) setContent(content string) {
	t.content = content
	t.notifyObservers()
}

type CheckBox struct {
	UIControl
	isChecked bool
}

func (c *CheckBox) setEnable(enable bool) {
	c.isChecked = enable
	c.notifyObservers()
}

type Button struct {
	UIControl
	isEnabled bool
}

func (b *Button) setEnable(enable bool) {
	b.isEnabled = enable
	b.notifyObservers()
}

type loginDialog struct {
	signupBotton *Button
	agreeToTerms *CheckBox
	username     *TextBox
	password     *TextBox
}

func NewLoginDialog() loginDialog {

	signupBotton := Button{UIControl: UIControl{observers: []observer{}}}
	agreeToTerms := CheckBox{UIControl: UIControl{observers: []observer{}}}
	username := TextBox{UIControl: UIControl{observers: []observer{}}}
	password := TextBox{UIControl: UIControl{observers: []observer{}}}

	newLoginDialog := loginDialog{signupBotton: &signupBotton, agreeToTerms: &agreeToTerms,
		username: &username, password: &password}

	c := controlChanged{loginDialog: newLoginDialog}
	newLoginDialog.username.addObserver(&c)
	newLoginDialog.password.addObserver(&c)
	newLoginDialog.agreeToTerms.addObserver(&c)

	return newLoginDialog
}

type controlChanged struct {
	loginDialog
}

func (c *controlChanged) update() {
	c.signupBotton.setEnable(c.isFormValid())
}

func (c *controlChanged) isFormValid() bool {
	return c.username.content != "" && c.password.content != "" && c.agreeToTerms.isChecked
}
