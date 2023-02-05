package correct

type Dim struct {
	LightSwitch *LightSwitch
}

func (d *Dim) Press() {
}

func (d *Dim) CanTurnOnLight() bool {
	return false
}
