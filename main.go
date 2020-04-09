package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"io/ioutil"
	"log"
)

// TODO: Add shortcut to copy body results
// to clipboard
func main() {

	g, err := gocui.NewGui(gocui.Output256)

	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	url := NewURLWidget(g, "url") //"https://jsonplaceholder.typicode.com/todos")
	body := NewBodyWidget(g, "body")
	helper := NewHelpBar(g, "help")
	box := NewBoxWidget(g, "box")
	box.Draw()
	helper.Draw()
	body.Draw()
	url.Draw()
	g.Cursor = true
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen
	//	g.SetViewOnBottom("body")

	// TODO: add other key binding to quit functionality
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, toggleView); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyF2, gocui.ModNone, createModal); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyF5, gocui.ModNone, displayRequestResults); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}

func createModal(g *gocui.Gui, v *gocui.View) error {
	for _, v := range g.Views() {
		_, err := g.SetViewOnBottom(v.Name())

		if err != nil {
			log.Panicln(err)
		}

		if v.Name() == "modal" {
			g.DeleteView(v.Name())
			return nil
		}
	}

	modal := NewModal(g, 10, 20, 30)
	modal.Draw()
	g.SetViewOnTop("modal")
	return nil
}

// TODO: add other views to toggle between when using
// the keyboard
func toggleView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == "body" {
		_, err := g.SetCurrentView("url")
		g.Cursor = true
		return err
	}
	_, err := g.SetCurrentView("body")
	g.Cursor = false
	return err
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// TODO: Return correct value
func displayRequestResults(g *gocui.Gui, v *gocui.View) error {
	// TODO: handle error returnded by View function
	g.SetCurrentView("body")
	// TODO: create helper to get view and not repeate these lines?
	bodyView, err := g.View("body")
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	urlView, err := g.View("url")
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	method := "GET" //TODO: dynamically get method

	bodyView.Clear()
	// TODO: send body for POST, PATCH and PUT requests
	err, response := GetResponse(method, urlView.Buffer(), "")

	if err != nil {
		fmt.Fprint(bodyView, err)
		return nil
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	fmt.Fprint(bodyView, string(body))
	//TODO: return correct value
	return nil
}
