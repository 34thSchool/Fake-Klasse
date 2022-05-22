package layouts

import (
	"fake-klasse/storage"
	"fake-klasse/ui"
	"strings"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func Add_Student(theme *material.Theme, operations *op.Ops, storage *storage.Storage, shouldQuit *bool) ui.Screen {

	// Widget declaration:
	var (
<<<<<<< HEAD
		nameWidget    widget.Editor
		surnameWidget widget.Editor
=======
		nameInput    widget.Editor
		surnameInput widget.Editor
>>>>>>> 7b9a823e9c36526339b028d2ba351d712a94bdad

		saveButton  widget.Clickable
		closeButton widget.Clickable
	)

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
					ui.DrawTitle(theme, operations, 70, "Add Student", ui.TitleColor, ui.Rect{Right: 0, Left: 0, Top: 0, Bottom: 0}),
				),
			)
		
			//Horizontal Middle Flexbox
			layout.Flex{
				Axis:    layout.Horizontal,
				Spacing: layout.SpaceAround,
			}.Layout(graphicalContext,
<<<<<<< HEAD
				layout.Flexed(1,ui.DrawInputWithMargins(theme, &nameWidget, "Name", 25, ui.Rect{Right: 0, Left: 50, Top: 150, Bottom: 0})),
				layout.Flexed(1,ui.DrawInputWithMargins(theme, &surnameWidget, "Surname", 25, ui.Rect{Right: 50, Left: 25, Top: 150, Bottom: 0})),
=======
				layout.Flexed(1,ui.DrawInputWithMargins(theme, &nameInput, "Name", 25, ui.Rect{Right: 0, Left: 50, Top: 150, Bottom: 0})),
				layout.Flexed(1,ui.DrawInputWithMargins(theme, &surnameInput, "Surname", 25, ui.Rect{Right: 50, Left: 25, Top: 150, Bottom: 0})),
>>>>>>> 7b9a823e9c36526339b028d2ba351d712a94bdad
			)
		
			// Flexbox with Bottom alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart, // Bottom
			}.Layout(graphicalContext,
			
<<<<<<< HEAD
				// Save button:
				layout.Rigid(
					ui.InputCheck(
						ui.DrawButtonWithMargins(theme, &saveButton, "Save", 15, ui.Rect{Right: 175,Left: 175,Top: 0, Bottom: 25}, ui.ButtonColor),
						nameWidget, surnameWidget,
=======
				// Add Student button:
				layout.Rigid(
					ui.InputCheck(
						ui.DrawButtonWithMargins(theme, &saveButton, "Save", 15, ui.Rect{Right: 175,Left: 175,Top: 0, Bottom: 25}, ui.ButtonColor),
						nameInput, surnameInput,
>>>>>>> 7b9a823e9c36526339b028d2ba351d712a94bdad
					),
				),
				// Close button:
				layout.Rigid(
					ui.DrawButtonWithMargins(theme, &closeButton, "Close", 15, ui.Rect{Right: 200,Left: 200,Top: 0, Bottom: 35}, ui.ButtonColor),
				),
			)
		}

		// Event handling:
		if closeButton.Clicked(){
			return Students(theme, operations, storage, shouldQuit), layout
		}
		if saveButton.Clicked(){
			storage.AddStudent(
<<<<<<< HEAD
				strings.TrimSpace(nameWidget.Text()),
				strings.TrimSpace(surnameWidget.Text()),
=======
				strings.TrimSpace(nameInput.Text()),
				strings.TrimSpace(surnameInput.Text()),
>>>>>>> 7b9a823e9c36526339b028d2ba351d712a94bdad
			)
			return Students(theme, operations, storage, shouldQuit), layout
		}

		return nil,layout

	}

}


