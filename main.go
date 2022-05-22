package main

import (
	"os"

	"gioui.org/app"         // Window handling.
	"gioui.org/font/gofont" // Special gioui font.
	"gioui.org/io/system"   // Events
	"gioui.org/layout"      // Dimensions, constraints, directions, flexbox.
	"gioui.org/op"
	"gioui.org/unit"            // implements device independent units and values. e.g. dp - device independent pixel, sp - scaled pixel - used for text sizes. and more.
	//"gioui.org/widget"          // UI component state tracking and event handling: Is the mouse hovering over the button? Is button pressed, and how many times?
	"gioui.org/widget/material" // theme

	"fake-klasse/ui"
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
	storage := storage.Storage{}
	storage.Init("School.db")
	defer storage.Close()
	
	// Filling DB:
	//storage.DeleteAllStudents()
	storage.AddStudent("Babaja", "Papaja")
	storage.AddStudent("Mambo", "Kurkuda")
	storage.AddStudent("Džuz", "Zeba")
	storage.AddStudent("Keto", "Vavai")

	// Rendering:
	var operations op.Ops
	theme := material.NewTheme(gofont.Collection())

	//currentLayout := ui.MainMenu(theme, &operations)// Declarating widgets and passing the drawing function as currentLayout. We're NOT drawing.
	currentLayout := ui.Students(theme, &operations, &storage)

	for event := range window.Events() {
		switch event := event.(type) {
		case system.FrameEvent:
			graphicalContext := layout.NewContext(&operations, event)

			// Drawing here:
			currentLayout(graphicalContext)
			
			event.Frame(graphicalContext.Ops)

		case system.DestroyEvent: // Sent when the app is closed.
			return event.Err // event.Err returns nil if app has been closed normally, and Err if something inappropriate caused a closure.
		}
	}

	return nil
}

// type Widget func(graphicalContext *layout.Context) layout.Dimensions