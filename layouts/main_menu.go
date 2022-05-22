//the file for main menu layout drawing and event handling
package layouts

import (
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"fake-klasse/storage"
	"fake-klasse/ui"
)

func MainMenu(theme *material.Theme, operations *op.Ops, shouldQuit *bool, storage *storage.Storage) ui.Screen {

	// Widget declaration:
	var (
		studentsButton widget.Clickable
		quitButton     widget.Clickable
	)
	
	return func(graphicalContext layout.Context) (ui.Screen, func(graphicalContext layout.Context)) {
		// Widget drawing:
		layout := func(graphicalContext layout.Context) {
			// Rendering:
			// Drawing background:
			ui.DrawBackground(operations, ui.BackgroundColor)
			// Flexbox with Top alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceEnd, // Top
				}.Layout(graphicalContext,
					layout.Rigid(
						ui.DrawTitle(theme, operations, 70, "Fake-Klasse", ui.TitleColor, ui.Rect{Right: 0, Left: 0, Top: 0, Bottom: 0}),
					),
				)
				// Flexbox with Middle alignment:
				layout.Flex{
					Axis:    layout.Vertical,
					Spacing: layout.SpaceAround, // Middle
					}.Layout(graphicalContext,
						layout.Rigid(
							ui.DrawButtonWithMargins(theme, &studentsButton, "Students", 15, ui.Rect{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
						),
					)
					// Flexbox with Bottom alignment:
					layout.Flex{
						Axis:    layout.Vertical,
						Spacing: layout.SpaceStart, // Bottom
						}.Layout(graphicalContext,
							layout.Rigid(
					ui.DrawButtonWithMargins(theme, &quitButton, "Quit", 15, ui.Rect{Right: 200, Left: 200, Top: 0, Bottom: 35}, ui.ButtonColor),
				),
			)
		}
		// Event handling:
		if studentsButton.Clicked() {
			return Students(theme, operations, storage, shouldQuit), layout
		}
		if quitButton.Clicked() {
			*shouldQuit = true
		}
		return nil, layout
	}
	

}
