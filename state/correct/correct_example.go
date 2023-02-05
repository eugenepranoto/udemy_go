package correct

import "fmt"

type State interface {
	Press()
	CanTurnOnLight() bool
}

type LightSwitch struct {
	State State
}

func (ls *LightSwitch) Press() {
	ls.State.Press()
}

func (ls *LightSwitch) CanTurnOnLight() bool {
	return ls.State.CanTurnOnLight()
}

func (ls *LightSwitch) ChangeState(state State) {
	ls.State = state
}

func (ls LightSwitch) IsElectricityOn() bool {
	return true
}

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

type Off struct {
	LightSwitch *LightSwitch
}

func (o *Off) Press() {
	if o.LightSwitch.IsElectricityOn() {
		return
	}

	o.LightSwitch.ChangeState(&On{LightSwitch: o.LightSwitch})
}

func (o *Off) CanTurnOnLight() bool {
	return false
}

type Dim struct {
	LightSwitch *LightSwitch
}

func (d *Dim) Press() {
}

func (d *Dim) CanTurnOnLight() bool {
	return false
}

func init() {
	ls := &LightSwitch{}
	ls.Press()
	ls.Press()
	fmt.Println(ls.CanTurnOnLight())
}
