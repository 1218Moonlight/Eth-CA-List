package gui

import (
	"github.com/andlabs/ui"
)

type mainBox struct {
	vertical *ui.Box
}

func newMainBox() mainBox {
	box := mainBox{vertical: func() *ui.Box {
		box := ui.NewVerticalBox()
		box.SetPadded(true)
		return box
	}()}

	return box
}
