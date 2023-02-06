package decorator

import "fmt"

type Opener interface {
	Open()
}

type Door struct {
}

func (d Door) Open() {
	fmt.Println("Door is open")
}

func init() {
	d := &Door{}
	d.Open()
	// door is open

	ed := NewElectronicKeyDoor(d)
	ed.Open()
	// new eletric
	// connect wifi
	// door is open

	md := NewMagicKeyDoor(ed)
	md.Open()
	// magical key is dangerous
	// Magic is working
	// new eletric
	// connect wifi
	// door is open
}
