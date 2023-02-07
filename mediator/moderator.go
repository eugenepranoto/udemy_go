package mediator

import "fmt"

type Moderator interface {
	Notify(s Participant)
}

type Teacher struct {
	P1, P2 Participant
}

func (t Teacher) Notify(s Participant) {
	var answer string
	if t.P1.GetName() == s.GetName() {
		answer = t.P1.Answer()
	} else {
		answer = t.P2.Answer()
	}
	fmt.Print(answer)
}
