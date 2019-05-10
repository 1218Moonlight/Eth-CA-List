package gui

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func Start() {
	ui.Main(func() {
		w := newWindow()
		w.onClosingAndQuit()

		mBox := newMainBox()

		w.setChild(mBox.vertical)
		w.show()
	})
}
