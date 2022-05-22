//the file for main menu layout drawing and event handling
package ui

import (
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func MainMenu(theme *material.Theme, operations *op.Ops) func(graphicalContext layout.Context) {

	// Widget declaration:
	var (
		studentsButton widget.Clickable
		quit           widget.Clickable
	)

	// Widget drawing:
	return func(graphicalContext layout.Context) {

		// Drawing background:
		DrawBackground(operations, backgroundColor)

		// Flexbox with Top alignment:
		layout.Flex{
			Axis:    layout.Vertical,
			Spacing: layout.SpaceEnd, // Top
		}.Layout(graphicalContext,
			layout.Rigid(
				DrawTitle(theme, operations, 70, "Fake-Klasse", titleColor, Rect{0, 0, 5, 0}),
			),
		)

		// Flexbox with Middle alignment:
		layout.Flex{
			Axis:    layout.Vertical,
			Spacing: layout.SpaceAround, // Middle
		}.Layout(graphicalContext,
			layout.Rigid(
				DrawButtonWithMargins(theme, &studentsButton, "Students", 15, Rect{175, 175, 0, 0}, buttonColor),
			),
		)

		// Flexbox with Bottom alignment:
		layout.Flex{
			Axis:    layout.Vertical,
			Spacing: layout.SpaceStart, // Bottom
		}.Layout(graphicalContext,
			layout.Rigid(
				DrawButtonWithMargins(theme, &quit, "Quit", 15, Rect{200, 200, 0, 75}, buttonColor),
			),
		)

	}
}
