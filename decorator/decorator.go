package decorator

import (
	"fmt"
	"math/rand"
)

func NewElectronicKeyDoor(opener Opener) Opener {
	return &EletronicKey{
		Opener: opener,
	}
}

func NewMagicKeyDoor(opener Opener) Opener {
	return &MagicalKey{
		Opener: opener,
	}
}

type EletronicKey struct {
	Opener Opener
}

func (e EletronicKey) Open() {
	fmt.Println("this is electronic key")
	e.ConnectToWifi()
	e.Opener.Open()
}

func (e EletronicKey) ConnectToWifi() {
	fmt.Println("Connecting to wifi")
}

type MagicalKey struct {
	Opener Opener
}

func (e MagicalKey) Open() {
	fmt.Println("magical key is dangerous")
	if e.CanWeUseMagic() {
		fmt.Println("Magic is working")
	} else {
		fmt.Println("cant user magic. User ur hand")
	}
	e.Opener.Open()
}

func (e MagicalKey) CanWeUseMagic() bool {
	return rand.Int() > 0
}
