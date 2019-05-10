package gui

import (
	"github.com/Kioryu/Eth-CA-List/contract"
	"github.com/andlabs/ui"
	"fmt"
)

type mainBox struct {
	vertical  *ui.Box
	barBox    barBox
	btn       *ui.Button
	outPutBox outPutBox
}

func newMainBox() mainBox {
	box := mainBox{
		vertical: func() *ui.Box {
			box := ui.NewVerticalBox()
			box.SetPadded(true)
			return box
		}(),
		barBox: newBarBox(),
		btn: func() *ui.Button {
			return ui.NewButton("Okay")
		}(),
		outPutBox: newOutPutBox()}

	return box
}

func (m mainBox) settingsBar() {
	spinbox := ui.NewSpinbox(0, 100)
	slider := ui.NewSlider(0, 100)

	m.barBox.vertical.Append(spinbox, false)
	m.barBox.vertical.Append(slider, false)
	m.barBox.vertical.Append(m.btn, false)

	spinbox.OnChanged(func(spinbox *ui.Spinbox) {
		slider.SetValue(spinbox.Value())
	})

	slider.OnChanged(func(slider *ui.Slider) {
		spinbox.SetValue(slider.Value())
	})

	m.btn.OnClicked(func(button *ui.Button) {
		m.caList(slider.Value())
	})

	m.barBox.group.SetChild(m.barBox.vertical)

	m.vertical.Append(m.barBox.group, false)
}

func (m mainBox) caList(nonce int) {
	//Todo
	ca := contract.CreateAddress("0x6f090f6cb125f77396d4b8f52fdabf7d5c1b53d4", uint64(nonce))
	m.outPutBox.line.Append(fmt.Sprintf("%s\n", ca))
}

func (m mainBox) setLogBox() {
	m.outPutBox.vertical.Append(m.outPutBox.line, true)

	m.outPutBox.group.SetChild(m.outPutBox.vertical)

	m.vertical.Append(m.outPutBox.group, true)
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

type outPutBox struct {
	group    *ui.Group
	vertical *ui.Box
	line     *ui.MultilineEntry
}

func newOutPutBox() outPutBox {
	o := outPutBox{
		group: func() *ui.Group {
			g := ui.NewGroup("Log")
			g.SetMargined(true)
			return g
		}(),
		vertical: func() *ui.Box {
			b := ui.NewVerticalBox()
			b.SetPadded(true)
			return b
		}(),
		line: func() *ui.MultilineEntry {
			n := ui.NewMultilineEntry()
			n.SetReadOnly(true)
			return n
		}()}

	return o
}
