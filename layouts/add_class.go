package layouts

import (
	"fake-klasse/state"
	"fake-klasse/storage"
	"fake-klasse/ui"
	"strings"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func Add_Class(state *state.State, theme *material.Theme, s *storage.Storage) ui.Screen {

	// Widget declaration:
	var (
		classWidget widget.Editor

		saveButton  widget.Clickable
		closeButton widget.Clickable
	)

	//Widget drawing:
	return func(gtx layout.Context) (ui.Screen, func(gtx layout.Context)) {

		layout := func(gtx layout.Context) {
			// Drawing background:
			ui.DrawBackground(gtx.Ops, ui.BackgroundColor)

			// Flexbox with Top alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceEnd, // Top
			}.Layout(gtx,
				// Title:
				layout.Rigid(
					ui.DrawTitle(state, theme, 70, "Add Class", ui.TitleColor, ui.Margins{Right: 0, Left: 0, Top: 0, Bottom: 0}),
				),

				layout.Flexed(1, ui.DrawInputWithMargins(state, theme, &classWidget, "Class", 25, ui.Margins{Right: 300, Left: 300, Top: 150, Bottom: 0})),
			)

			//Vertical  Flexbox
			// layout.Flex{
			// 	Axis:    layout.Vertical,//Horizontal
			// 	Spacing: layout.SpaceEnd,//spacearound
			// }.Layout(gtx,
			// 	layout.Flexed(1,ui.DrawInputWithMargins(&classWidget, "Class", 25, ui.Rect{Right: 50, Left: 0, Top: 150, Bottom: 0})),
			// )

			// Flexbox with Bottom alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart, // Bottom
			}.Layout(gtx,

				// Save button:
				layout.Rigid(
					ui.ClassInputCheck(
						ui.DrawButtonWithMargins(state, theme, &saveButton, "Save", 15, ui.Margins{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
						classWidget,
					),
				),
				// Close button:
				layout.Rigid(
					ui.DrawButtonWithMargins(state, theme, &closeButton, "Close", 15, ui.Margins{Right: 200, Left: 200, Top: 0, Bottom: 35}, ui.ButtonColor),
				),
			)
		}

		// Event handling:
		if closeButton.Clicked() {
			return Classes(state, theme, s), layout
		}
		if saveButton.Clicked() {
			//Calling add class function from Storage file

			s.AddClass(
				strings.TrimSpace(classWidget.Text()),
			)
			return Classes(state, theme, s), layout
		}
		return nil, layout

	}

}
