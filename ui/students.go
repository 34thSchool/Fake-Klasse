//the file with student list layout drawing
package ui

import (
	"fake-klasse/storage"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func Students(theme *material.Theme, operations *op.Ops, storage *storage.Storage) func(graphicalContext layout.Context) {

	// Widget declaration:
	var (
		addStudentButton widget.Clickable
		close            widget.Clickable
		list             widget.List = widget.List{List: layout.List{Axis: layout.Vertical}}
	)

	//Creating a widget.Clickable slice of all students in DB
	students := storage.GetStudents()
	var widgetList []widget.Clickable
	for range(*students){
	 	var widget widget.Clickable 
	 	widgetList = append(widgetList, widget)
	}


	// Widget drawing:
	return func(graphicalContext layout.Context) {

		// Drawing background:
		DrawBackground(operations, backgroundColor)

		// Flexbox with Top alignment:
		layout.Flex{
			Axis:    layout.Vertical,
			Spacing: layout.SpaceEnd, // Top
		}.Layout(graphicalContext,
			// Title:
			layout.Rigid(
				DrawTitle(theme, operations, 70, "Students", titleColor, Rect{0, 0, 0, 0}),
			),
			// List:
			layout.Rigid(
				DrawListWithMargins(graphicalContext, theme, &widgetList, storage.GetStudents(), &list, Rect{0, 0, 0, 175}),
			),
		)

		// Flexbox with Bottom alignment:
		layout.Flex{
			Axis:    layout.Vertical,
			Spacing: layout.SpaceStart, // Bottom
		}.Layout(graphicalContext,

			// Add Student button:
			layout.Rigid(
				DrawButtonWithMargins(theme, &addStudentButton, "Add Student", 15, Rect{175, 175, 0, 25}, buttonColor),
			),
			// Close button:
			layout.Rigid(
				DrawButtonWithMargins(theme, &close, "Close", 15, Rect{200, 200, 0, 35}, buttonColor),
			),
		)
	}

}
