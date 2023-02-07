package memento

type Controller struct {
	Editor    *Editor
	Histories []Memento
	ptr       int // pointer for current snapshoot
}

func (c *Controller) CmdSave() {
	m := c.Editor.Save()
	c.Histories = append(c.Histories, m)
	l := len(c.Histories)
	c.ptr = l - 1
}

func (c *Controller) CmdUndo() {
	if c.ptr == 0 {
		return
	}

	c.ptr--
	m := c.Histories[c.ptr]
	c.Editor.Restore(m)
}

func (c *Controller) CmdRedo() {
	l := len(c.Histories)
	if c.ptr+1 >= l {
		return
	}

	c.ptr++
	m := c.Histories[c.ptr]
	c.Editor.Restore(m)
}
