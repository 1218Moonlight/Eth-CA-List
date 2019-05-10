package gui

import (
	"github.com/andlabs/ui"
	"fmt"
)

type mainBox struct {
	vertical *ui.Box
	barBox   barBox
	button   okBtn
}

func newMainBox() mainBox {
	box := mainBox{
		vertical: func() *ui.Box {
			box := ui.NewVerticalBox()
			box.SetPadded(true)
			return box
		}(),
		barBox: newBarBox(),
		button: newOkBtn()}

	return box
}

func (m mainBox) settingsBar() {
	spinbox := ui.NewSpinbox(0, 100)
	slider := ui.NewSlider(0, 100)

	m.barBox.vertical.Append(spinbox, false)
	m.barBox.vertical.Append(slider, false)
	m.barBox.vertical.Append(m.button.btn, false)

	spinbox.OnChanged(func(spinbox *ui.Spinbox) {
		slider.SetValue(spinbox.Value())
	})

	slider.OnChanged(func(slider *ui.Slider) {
		spinbox.SetValue(slider.Value())
	})

	m.button.btn.OnClicked(func(button *ui.Button) {
		m.button.number = slider.Value()
		fmt.Println("[] : ", m.button.number)
		// Todo
	})

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

type okBtn struct {
	number int
	btn *ui.Button
}

func newOkBtn() okBtn {
	o := okBtn{
		btn: func() *ui.Button {
			return ui.NewButton("Okay")
		}()}

	return o
}
