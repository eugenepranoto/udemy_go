package incorrect

// jika light state banyak, maka else if akan makin banyak
// makin susah ditest
// tidak scalable

type LightState string

const (
	On  LightState = "On"
	Off LightState = "Off"
	Dim LightState = "Dim"
)

type LightSwitch struct {
	State LightState
}

func (l *LightSwitch) Press() {
	if l.State == Off {
		l.State = On
	} else if l.State == On {
		l.State = Off
	} else {
		l.State = Dim
	}
}

func (l *LightSwitch) CanTurnOnLight() bool {
	return l.State == On
}
