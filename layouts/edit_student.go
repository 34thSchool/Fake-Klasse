package layouts

import (
	"fake-klasse/storage"
	"fake-klasse/ui"
	"fmt"
	"strings"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func Edit_Student(theme *material.Theme, operations *op.Ops, storage *storage.Storage, shouldQuit *bool, id int) ui.Screen {

	// Widget declaration:
	var (
		nameWidget    widget.Editor
		surnameWidget widget.Editor

		saveButton          widget.Clickable
		closeButton         widget.Clickable
		deleteStudentButton widget.Clickable
	)
	
	students := storage.GetStudents()

	//Widget drawing:
	return func(graphicalContext layout.Context) (ui.Screen, func(graphicalContext layout.Context)) {

		layout := func(graphicalContext layout.Context) {

			// Drawing background:
			ui.DrawBackground(operations, ui.BackgroundColor)

			// Flexbox with Top alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceEnd, // Top
			}.Layout(graphicalContext,
				// Title:
				layout.Rigid(
					ui.DrawTitle(theme, operations, 70, "Edit Student", ui.TitleColor, ui.Rect{Right: 0, Left: 0, Top: 0, Bottom: 0}),
				),
			)

			//Horizontal Middle Flexbox
			layout.Flex{
				Axis:    layout.Horizontal,
				Spacing: layout.SpaceAround,
			}.Layout(graphicalContext,
				layout.Flexed(1, ui.DrawInputWithMargins(theme, &nameWidget, (*students)[id].Name, 25, ui.Rect{Right: 0, Left: 50, Top: 150, Bottom: 0})),
				layout.Flexed(1, ui.DrawInputWithMargins(theme, &surnameWidget, (*students)[id].Surname, 25, ui.Rect{Right: 50, Left: 25, Top: 150, Bottom: 0})),
			)

			// Flexbox with Bottom alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart, // Bottom
			}.Layout(graphicalContext,
				// Save button:
				layout.Rigid(
					ui.InputCheck(
						ui.DrawButtonWithMargins(theme, &saveButton, "Save", 15, ui.Rect{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
						nameWidget, surnameWidget,
					),
				),
				// Delete Student button:
				layout.Rigid(
					ui.DrawButtonWithMargins(theme, &deleteStudentButton, "Delete Student", 15, ui.Rect{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
				),
				// Close button:
				layout.Rigid(
					ui.DrawButtonWithMargins(theme, &closeButton, "Close", 15, ui.Rect{Right: 200, Left: 200, Top: 0, Bottom: 35}, ui.ButtonColor),
				),
			)
		}

		// Event handling:
		if closeButton.Clicked() {
			return Students(theme, operations, storage, shouldQuit), layout
		}
		if saveButton.Clicked() {
			storage.DeleteStudent((*students)[id].Rowid)
			storage.AddStudent(
				strings.TrimSpace(nameWidget.Text()),
				strings.TrimSpace(surnameWidget.Text()),
			)
			fmt.Println("name widget: ", strings.TrimSpace(nameWidget.Text()))
			fmt.Println("surname widget: ", strings.TrimSpace(surnameWidget.Text()))
			return Students(theme, operations, storage, shouldQuit), layout
		}
		if deleteStudentButton.Clicked() {
			storage.DeleteStudent((*students)[id].Rowid)
			return Students(theme, operations, storage, shouldQuit), layout
		}

		return nil, layout

	}
}
