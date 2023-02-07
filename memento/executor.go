package memento

import "fmt"

func Execute() {
	editor := &Editor{}
	editor.Text = "Text Pertama"

	Controller := Controller{
		Editor: editor,
	}

	Controller.CmdSave()
	fmt.Println(editor.Text)

	editor.Text = "Text pertama uda diupdate"
	Controller.CmdSave()
	fmt.Println(editor.Text)

	Controller.CmdUndo()
	fmt.Println(editor.Text)

	Controller.CmdRedo()
	fmt.Println(editor.Text)
}
