package prototype

import "fmt"

type Prototype interface {
	Clone() Prototype
}

type Robot struct {
	ID        string
	Name      string
	processor string
}

func (r Robot) Clone() Prototype {
	return &Robot{
		Name:      r.Name,
		ID:        r.ID,
		processor: r.processor,
	}
}

func NewRobot(p Prototype) Prototype {
	return p.Clone()
}

func init() {
	r := Robot{
		ID:   "A",
		Name: "Edo",
	}
	r2 := r.Clone()

	robot2, ok := r2.(*Robot)
	if ok {
		robot2.Name = "John"
	}

	fmt.Println(r.Name, r2.(*Robot).Name, robot2.Name)
}
