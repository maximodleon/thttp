package main

import (
	"github.com/jroimartin/gocui"
)

// URL Widget
type BoxWidget struct {
	*gocui.Gui
	label    string
	handlers Handlers
	*Position
	*Attributes
}

func NewBoxWidget(gui *gocui.Gui, label string) *BoxWidget {

	return &BoxWidget{
		Gui:   gui,
		label: label,
		Position: &Position{
			X: 2,
			Y: 2,
			H: 4,
			W: 13,
		},
		Attributes: &Attributes{
			textColor: gocui.ColorWhite | gocui.AttrBold,
			hlColor:   gocui.ColorBlue | gocui.AttrBold,
			hlBgColor: gocui.ColorWhite,
		},
		handlers: make(Handlers),
	}
}

func (u *BoxWidget) Draw() {
	v, err := u.Gui.SetView(u.label, u.X, u.Y, u.W, u.H)

	if err != nil {
		if err != gocui.ErrUnknownView {
			panic(err)
		}

		v.Frame = true
		v.Editable = true
		v.Title = "Method"
		v.Write([]byte("GET"))

		if _, err := u.Gui.SetCurrentView(u.label); err != nil {
			panic(err)
		}
	}
}
