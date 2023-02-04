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

	ed := NewElectronicKeyDoor(d)
	ed.Open()

	md := NewMagicKeyDoor(ed)
	md.Open()
}
