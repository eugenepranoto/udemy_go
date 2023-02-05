package factory

import (
	"fmt"
	"udemy/go/prototype"
)

type Factory struct {
}

func (f Factory) CopyRobot(r *prototype.Robot) (*prototype.Robot, error) {
	res, ok := r.Clone().(*prototype.Robot)
	if !ok {
		return nil, fmt.Errorf("failed")
	}
	return res, nil
}
