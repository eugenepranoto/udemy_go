package factory

import "udemy/go/prototype"

type Factory struct {
}

func (f Factory) CopyRobot(r *prototype.Robot) *prototype.Robot {
	return r.Clone().(*prototype.Robot)
}
