package gui

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"log"
)

type window struct {
	main      *ui.Window
	configure windowConfigure
}

type windowConfigure struct {
	title      string
	width      int
	height     int
	hasMenubar bool
}

func newWindow() window {
	w := window{main: nil,
		configure: windowConfigure{title: title, width: width, height: height, hasMenubar: hasMenubar},
	}

	w.main = ui.NewWindow(w.configure.title, w.configure.width, w.configure.height, w.configure.hasMenubar)
	return w
}

func (w window) onClosingAndQuit() {
	w.main.OnClosing(func(window *ui.Window) bool {
		ui.Quit()
		log.Println("OnClosing")
		return true
	})

	ui.OnShouldQuit(func() bool {
		w.main.Destroy()
		log.Println("OnShouldQuit")
		return true
	})
}

func (w window) show() {
	w.main.Show()
}
