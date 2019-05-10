package gui

import (
	"github.com/andlabs/ui"
)

type mainBox struct {
	vertical *ui.Box
	barBox   barBox
}

func newMainBox() mainBox {
	box := mainBox{vertical: func() *ui.Box {
		box := ui.NewVerticalBox()
		box.SetPadded(true)
		return box
	}(),
		barBox: newBarBox()}

	return box
}

func (m mainBox) settingsBar() {
	spinbox := ui.NewSpinbox(0, 100)
	slider := ui.NewSlider(0, 100)

	m.barBox.vertical.Append(spinbox, false)
	m.barBox.vertical.Append(slider, false)

	m.barBox.group.SetChild(m.barBox.vertical)

	m.vertical.Append(m.barBox.group, false)
}

type barBox struct {
	group    *ui.Group
	vertical *ui.Box
}

func newBarBox() barBox {
	b := barBox{
		group: func() *ui.Group {
			g := ui.NewGroup("setting")
			g.SetMargined(true)
			return g
		}(),
		vertical: func() *ui.Box {
			b := ui.NewVerticalBox()
			b.SetPadded(true)
			return b
		}()}

	return b
}
