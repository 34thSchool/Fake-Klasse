package layouts

import (
	"fake-klasse/state"
	"fake-klasse/storage"
	"fake-klasse/ui"
	"strings"

	"gioui.org/layout"
	"gioui.org/widget"
)

func Add_Class(state *state.State, s *storage.Storage) ui.Screen {

	// Widget declaration:
	var (
		classWidget widget.Editor

		saveButton  widget.Clickable
		closeButton widget.Clickable
	)

	//Widget drawing:
	return func(graphicalContext layout.Context) (ui.Screen, func(graphicalContext layout.Context)) {

		layout := func(graphicalContext layout.Context) {
			// Drawing background:
			ui.DrawBackground(graphicalContext.Ops, ui.BackgroundColor)

			// Flexbox with Top alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceEnd, // Top
			}.Layout(graphicalContext,
				// Title:
				layout.Rigid(
					ui.DrawTitle(state, 70, "Add Class", ui.TitleColor, ui.Margins{Right: 0, Left: 0, Top: 0, Bottom: 0}),
				),

				layout.Flexed(1, ui.DrawInputWithMargins(state, &classWidget, "Class", 25, ui.Margins{Right: 300, Left: 300, Top: 150, Bottom: 0})),
			)

			//Vertical  Flexbox
			// layout.Flex{
			// 	Axis:    layout.Vertical,//Horizontal
			// 	Spacing: layout.SpaceEnd,//spacearound
			// }.Layout(graphicalContext,
			// 	layout.Flexed(1,ui.DrawInputWithMargins(&classWidget, "Class", 25, ui.Rect{Right: 50, Left: 0, Top: 150, Bottom: 0})),
			// )

			// Flexbox with Bottom alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart, // Bottom
			}.Layout(graphicalContext,

				// Save button:
				layout.Rigid(
					ui.ClassInputCheck(
						ui.DrawButtonWithMargins(state, &saveButton, "Save", 15, ui.Margins{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
						classWidget,
					),
				),
				// Close button:
				layout.Rigid(
					ui.DrawButtonWithMargins(state, &closeButton, "Close", 15, ui.Margins{Right: 200, Left: 200, Top: 0, Bottom: 35}, ui.ButtonColor),
				),
			)
		}

		// Event handling:
		if closeButton.Clicked() {
			return Classes(state, s), layout
		}
		if saveButton.Clicked() {
			//Calling add class function from Storage file

			s.AddClass(
				strings.TrimSpace(classWidget.Text()),
			)
			return Classes(state, s), layout
		}
		return nil, layout

	}

}
