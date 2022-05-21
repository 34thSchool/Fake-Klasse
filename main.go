package main

import (
	"os"

	"gioui.org/app"         // Window handling.
	"gioui.org/font/gofont" // Special gioui font.
	"gioui.org/io/system"   // Events
	"gioui.org/layout"      // Dimensions, constraints, directions, flexbox.
	"gioui.org/op"
	"gioui.org/unit" // implements device independent units and values. e.g. dp - device independent pixel, sp - scaled pixel - used for text sizes. and more.

	//"gioui.org/widget"          // UI component state tracking and event handling: Is the mouse hovering over the button? Is button pressed, and how many times?
	"gioui.org/widget/material" // theme

	"fake-klasse/ui"
)

func main() {

	go func() {
		window := app.NewWindow(
			app.Title("Nach der Schule"),
			app.Size(unit.Dp(800), unit.Dp(800)),
			app.MaxSize(unit.Dp(800), unit.Dp(800)),
			app.MinSize(unit.Dp(800), unit.Dp(800)),
		)

		mainLoop(window)

		os.Exit(1)
	}()
	app.Main()
}

func mainLoop(window *app.Window) error {

	var operations op.Ops
	theme := material.NewTheme(gofont.Collection())

	currentLayout := ui.MainMenu(theme, &operations) // Declarating widgets and passing the drawing function as currentLayout. We're NOT drawing.

	for event := range window.Events() {
		switch event := event.(type) {
		case system.FrameEvent:
			graphicalContext := layout.NewContext(&operations, event)

			// Drawing here:
			currentLayout(graphicalContext) // Drawing;)

			event.Frame(graphicalContext.Ops)

		case system.DestroyEvent: // Sent when the app is closed.
			return event.Err // event.Err returns nil if app has been closed normally, and Err if something inappropriate caused a closure.
		}
	}

	return nil
}

// type Widget func(graphicalContext *layout.Context) layout.Dimensions
