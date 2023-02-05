package correct

import "fmt"

type State interface {
	Press()
	CanTurnOnLight() bool
}

type LightSwitch struct {
	State State
}

func NewLightSwitch() *LightSwitch {
	ls := &LightSwitch{}
	ls.ChangeState(&Off{ls})
	return ls
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

func init() {
	ls := NewLightSwitch()
	ls.Press()
	ls.Press()
	fmt.Println(ls.CanTurnOnLight())
}
