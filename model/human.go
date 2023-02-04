package model

type person interface {
	speak()
}

type Human struct {
	Person person
}

func (p Human) DoSpeak() {
	p.Person.speak()
}
