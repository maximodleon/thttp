package main

import (
	//	"fmt"
	"github.com/jroimartin/gocui"
	//	"log"
)

type Modal struct {
	*gocui.Gui
	name string
	*Attributes
	*Position
}

func NewModal(gui *gocui.Gui, x, y, w int) *Modal {
	p := &Position{
		X: x,
		Y: y,
		W: w,
		H: y + 3,
	}

	return &Modal{
		Gui:  gui,
		name: "modal",
		Attributes: &Attributes{
			textColor:   gocui.ColorWhite,
			textBgColor: gocui.ColorBlue,
		},
		Position: p,
	}
}

func (m *Modal) Draw() {
	if v, err := m.Gui.SetView(m.name, m.X, m.Y, m.W, m.H); err != nil {
		if err != gocui.ErrUnknownView {
			panic(err)
		}

		v.Frame = true
		v.FgColor = m.textColor
		v.BgColor = m.textBgColor

		// set modal on top
		if _, err := m.Gui.SetViewOnTop(m.name); err != nil {
			if err != gocui.ErrUnknownView {
				panic(err)
			}
		}

	}
}
