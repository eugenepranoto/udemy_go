package correct

type On struct {
	LightSwitch *LightSwitch
	Test        string
}

func (o *On) Press() {
	o.LightSwitch.ChangeState(&Off{LightSwitch: o.LightSwitch})
}

func (o *On) CanTurnOnLight() bool {
	return true
}
