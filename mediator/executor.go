package mediator

func Execute() {
	t := &Teacher{}
	s1 := HighSchoolStudent{
		Name:      "Him",
		Moderator: t,
	}
	s2 := HighSchoolStudent{
		Name:      "Edo",
		Moderator: t,
	}

	t.P1 = s1
	t.P2 = s2

	s1.PressButton()
	s2.PressButton()
}
