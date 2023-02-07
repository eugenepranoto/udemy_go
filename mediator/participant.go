package mediator

import "fmt"

type Participant interface {
	Answer() string
	PressButton()
	GetName() string
}

type HighSchoolStudent struct {
	Moderator Moderator
	Name      string
}

func (h HighSchoolStudent) Answer() string {
	return fmt.Sprintf("%s Menjawab pertanyaan", h.Name)
}

func (h HighSchoolStudent) PressButton() {
	h.Moderator.Notify(h)
}

func (h HighSchoolStudent) GetName() string {
	return h.Name
}
