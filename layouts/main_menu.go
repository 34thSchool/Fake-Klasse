//the file for main menu layout drawing and event handling
package layouts

import (
	"gioui.org/layout"
	"gioui.org/widget"

	"fake-klasse/state"
	"fake-klasse/storage"
	"fake-klasse/ui"
)

func MainMenu(state *state.State, s *storage.Storage) ui.Screen {

	// Widget declaration:
	var (
		studentsButton widget.Clickable
		classesButton  widget.Clickable
		quitButton     widget.Clickable
	)

	return func(graphicalContext layout.Context) (ui.Screen, func(graphicalContext layout.Context)) {
		// Widget drawing:
		layout := func(graphicalContext layout.Context) {
			// Rendering:
			// Drawing background:
			ui.DrawBackground(graphicalContext.Ops, ui.BackgroundColor)
			// Flexbox with Top alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceEnd, // Top
			}.Layout(graphicalContext,
				layout.Rigid(
					ui.DrawTitle(state, 70, "Fake-Klasse", ui.TitleColor, ui.Margins{Right: 0, Left: 0, Top: 0, Bottom: 0}),
				),
			)
			// Flexbox with Middle alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceSides, // Middle
			}.Layout(graphicalContext,
				layout.Rigid(
					ui.DrawButtonWithMargins(state, &studentsButton, "Students", 15, ui.Margins{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
				),
				layout.Rigid(
					ui.DrawButtonWithMargins(state, &classesButton, "Classes", 15, ui.Margins{Right: 175, Left: 175, Top: 0, Bottom: 0}, ui.ButtonColor),
				),
			)
			// Flexbox with Bottom alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart, // Bottom
			}.Layout(graphicalContext,
				layout.Rigid(
					ui.DrawButtonWithMargins(state, &quitButton, "Quit", 15, ui.Margins{Right: 200, Left: 200, Top: 0, Bottom: 35}, ui.ButtonColor),
				),
			)
		}
		// Event handling:
		if studentsButton.Clicked() {
			return Students(state, s), layout
		}
		if classesButton.Clicked() {
			return Classes(state, s), layout
		}
		if quitButton.Clicked() {
			state.ShouldQuit = true
		}
		return nil, layout
	}

}
