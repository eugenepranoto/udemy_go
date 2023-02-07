package memento

// buat undo previous state kyk ctrl+z
// buat controller yang menyimpan snapshot history
// jalankan perintah di controller
// untuk mengupdate/mengambil state

type Editor struct {
	Text string
}

func (e *Editor) Save() Memento {
	return Snapshot{Text: e.Text}
}

func (e *Editor) Restore(m Memento) {
	e.Text = m.GetText()
}
