package main

import (
	"os"

	"gioui.org/app"       // Window handling.
	"gioui.org/io/system" // Events
	"gioui.org/layout"    // Dimensions, constraints, directions, flexbox.
	"gioui.org/op"
	"gioui.org/unit" // implements device independent units and values. e.g. dp - device independent pixel, sp - scaled pixel - used for text sizes. and more.

	//"gioui.org/widget"          // UI component state tracking and event handling: Is the mouse hovering over the button? Is button pressed, and how many times?

	"fake-klasse/layouts"
	"fake-klasse/state"
	"fake-klasse/storage"
)

func main() {
	go func() {
		window := app.NewWindow(
			app.Title("Fake-Klasse"),
			app.Size(unit.Dp(800), unit.Dp(800)),
			app.MinSize(unit.Dp(500), unit.Dp(500)),
		)

		mainLoop(window)

		os.Exit(0)
	}()
	app.Main()
}

func mainLoop(window *app.Window) error {

	// Initializing DB:
	storage.Singleton.Init("school.db")
	defer storage.Singleton.Close()

	

	currentLayout := layouts.MainMenu() // Declarating widgets and passing the drawing function as currentLayout. We're NOT drawing.
	//currentLayout := layouts.Students(theme, &operations, &storage)

	for event := range window.Events() {
		switch event := event.(type) {
		case system.FrameEvent:
			graphicalContext := layout.NewContext(&op.Ops{}, event)

			// Drawing here:
			nextLayout, drawLayout := currentLayout(graphicalContext)
			drawLayout(graphicalContext)
			if nextLayout != nil {
				currentLayout = nextLayout
			}

			// Checking whether or not we should quit:
			if state.ShouldQuit {
				window.Perform(system.ActionClose)
			}

			event.Frame(graphicalContext.Ops)

		case system.DestroyEvent: // Sent when the app is closed.
			return event.Err // event.Err returns nil if app has been closed normally, and Err if something inappropriate caused a closure.
		}
	}

	return nil
}
