// jika misal ada suatu object A yang punya pungsi dan
// membutuhkan input dengan format tertentu
// tapi kita punya inputnya beda format
// kita bisa bikin object adapter. dimana
// adapteenya ada object A, lalu buat fungsi yang sama
// pada adapter seperti fungsi di object A
// tapi tipe paramnya sesuai dengan tipe data yang kita punya

package adapter

import "fmt"

type Audio interface {
	Play(m Mp3)
}

type Mp3 struct {
	Data []byte
}

func (m Mp3) PlayAudio() {
	fmt.Println("Play mp3" + string(m.Data))
}

type Kaset struct {
	PitaMusik string
}

type Walkman struct {
}

func (w Walkman) Play(c Kaset) {
	fmt.Println(c.PitaMusik)
}

type Mp3ToKasetAdapter struct {
	Adaptee Walkman
}

func (a Mp3ToKasetAdapter) Play(m Mp3) {
	k := Kaset{}
	k.PitaMusik = string(m.Data)
	a.Adaptee.Play(k)
}

func init() {
	mp3 := Mp3{Data: []byte(`This is Test Taylor swift songs`)}
	walkman := Walkman{}
	adapter := Mp3ToKasetAdapter{Adaptee: walkman}
	adapter.Play(mp3)
}
