package layouts

import (
	"fake-klasse/storage"
	"fake-klasse/ui"
	"strings"

	"gioui.org/layout"
	"gioui.org/widget"
)

func Add_Student() ui.Screen {

	// Widget declaration:
	var (
		nameWidget    widget.Editor
		surnameWidget widget.Editor

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
					ui.DrawTitle(70, "Add Student", ui.TitleColor, ui.Rect{Right: 0, Left: 0, Top: 0, Bottom: 0}),
				),
			)
		
			//Horizontal Middle Flexbox
			layout.Flex{
				Axis:    layout.Horizontal,
				Spacing: layout.SpaceAround,
			}.Layout(graphicalContext,
				layout.Flexed(1,ui.DrawInputWithMargins(&nameWidget, "Name", 25, ui.Rect{Right: 0, Left: 50, Top: 150, Bottom: 0})),
				layout.Flexed(1,ui.DrawInputWithMargins(&surnameWidget, "Surname", 25, ui.Rect{Right: 50, Left: 25, Top: 150, Bottom: 0})),
			)
		
			// Flexbox with Bottom alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart, // Bottom
			}.Layout(graphicalContext,
			
				// Save button:
				layout.Rigid(
					ui.InputCheck(
						ui.DrawButtonWithMargins(&saveButton, "Save", 15, ui.Rect{Right: 175,Left: 175,Top: 0, Bottom: 25}, ui.ButtonColor),
						nameWidget, surnameWidget,
					),
				),
				// Close button:
				layout.Rigid(
					ui.DrawButtonWithMargins(&closeButton, "Close", 15, ui.Rect{Right: 200,Left: 200,Top: 0, Bottom: 35}, ui.ButtonColor),
				),
			)
		}

		// Event handling:
		if closeButton.Clicked(){
			return Students(), layout
		}
		if saveButton.Clicked(){
			storage.Singleton.AddStudent(
				strings.TrimSpace(nameWidget.Text()),
				strings.TrimSpace(surnameWidget.Text()),
			)
			return Students(), layout
		}

		return nil,layout

	}

}


