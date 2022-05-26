package layouts

import (
	"fake-klasse/storage"
	"fake-klasse/ui"
	"strings"

	"gioui.org/layout"
	"gioui.org/widget"
)

func Edit_Class(index int) ui.Screen {

	// Widget declaration:
	var (
		classWidget widget.Editor
		className   string

		saveButton        widget.Clickable
		closeButton       widget.Clickable
		deleteClassButton widget.Clickable
	)

	className = (*storage.Singleton.GetAllClasses())[index].Name

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
					ui.DrawTitle(70, "Edit Class", ui.TitleColor, ui.Rect{Right: 0, Left: 0, Top: 0, Bottom: 0}),
				),
				//Input
				layout.Flexed(1,
					ui.DrawInputWithMargins(&classWidget, className, 25, ui.Rect{Right: 300, Left: 300, Top: 150, Bottom: 0}),
				),
			)

			// Flexbox with Bottom alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart, // Bottom
			}.Layout(graphicalContext,
				// Save button:
				layout.Rigid(
					//ui.ClassInputCheck(
					ui.DrawButtonWithMargins(&saveButton, "Save", 15, ui.Rect{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
					//classWidget,
					//),
				),
				// Delete Student button:
				layout.Rigid(
					ui.DrawButtonWithMargins(&deleteClassButton, "Delete Class", 15, ui.Rect{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
				),
				// Close button:
				layout.Rigid(
					ui.DrawButtonWithMargins(&closeButton, "Close", 15, ui.Rect{Right: 200, Left: 200, Top: 0, Bottom: 35}, ui.ButtonColor),
				),
			)
		}

		// Event handling:
		if closeButton.Clicked() {
			return Class_Students(index), layout
		}
		if saveButton.Clicked() {

			storage.Singleton.UpdateAllClassStudentsClass((*storage.Singleton.GetAllClasses())[index].Name, strings.TrimSpace(classWidget.Text()))

			storage.Singleton.DeleteClass((*storage.Singleton.GetAllClasses())[index])

			storage.Singleton.AddClass(
				strings.TrimSpace(classWidget.Text()),
			)

			return Class_Students(index), layout
		}
		if deleteClassButton.Clicked() {

			storage.Singleton.UpdateAllClassStudentsClass((*storage.Singleton.GetAllClasses())[index].Name, "")
			storage.Singleton.DeleteClass((*storage.Singleton.GetAllClasses())[index])
			className = ""

			return Classes(), layout
		}

		return nil, layout

	}
}
