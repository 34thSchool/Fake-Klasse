//the file with student list layout drawing
package layouts

import (
	"fake-klasse/storage"
	"fake-klasse/ui"

	"gioui.org/layout"
	"gioui.org/widget"
)

func Students() ui.Screen {

	// Widget declaration:
	var (
		addStudentButton widget.Clickable
		closeButton      widget.Clickable
		list             widget.List = widget.List{List: layout.List{Axis: layout.Vertical}}
	)
	//Creating a widget.Clickable slice of all students in DB
	students := storage.Singleton.GetStudents()
	var widgetList []widget.Clickable
	for range *students {
		var widget widget.Clickable
		widgetList = append(widgetList, widget)
	}

	return func(graphicalContext layout.Context) (ui.Screen, func(graphicalContext layout.Context)) {
		// Rendering:
		// Widget drawing:
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
					ui.DrawTitle(70, "Students", ui.TitleColor, ui.Rect{Right: 0, Left: 0, Top: 0, Bottom: 0}),
				),
				// List:
				layout.Rigid(
					ui.DrawListWithMargins(graphicalContext, &widgetList, storage.Singleton.GetStudents(), &list, ui.Rect{Right: 0, Left: 0, Top: 0, Bottom: 175}),
				),
			)

			// Flexbox with Bottom alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart, // Bottom
			}.Layout(graphicalContext,

				// Add Student button:
				layout.Rigid(
					ui.DrawButtonWithMargins(&addStudentButton, "Add Student", 15, ui.Rect{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
				),
				// Close button:
				layout.Rigid(
					ui.DrawButtonWithMargins(&closeButton, "Close", 15, ui.Rect{Right: 200, Left: 200, Top: 0, Bottom: 35}, ui.ButtonColor),
				),
			)
		}
		// Event handling:
		if closeButton.Clicked() {
			return MainMenu(), layout
		}
		if addStudentButton.Clicked() {
			return Add_Student(), layout
		}
		for index := range widgetList {
			if widgetList[index].Clicked() {
				return Edit_Student(index), layout
			}
		}
		return nil, layout
	}

}
