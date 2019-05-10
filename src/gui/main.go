package gui

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func Start(){
	ui.Main(func() {
		w := newWindow()
		w.onClosingAndQuit()

		// todo

		w.show()
	})
}