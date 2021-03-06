package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

var DefaultEditor Editor = EditorFunc(simpleEditor)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", rune(113), gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", rune(97), gocui.ModNone, addChannel); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", 0, 0, maxX/10, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, []rune("qa"))
	}
	return nil
}

func addChannel(g *gocui.Gui, v *gocui.View) error {
	g.Views()[0].Clear()
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
