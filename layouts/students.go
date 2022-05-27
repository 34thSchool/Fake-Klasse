//the file with student list layout drawing
package layouts

import (
	"fake-klasse/state"
	"fake-klasse/storage"
	"fake-klasse/ui"
	"log"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func Students(state *state.State, theme *material.Theme, s *storage.Storage) ui.Screen {

	// Widget declaration:
	var (
		addStudentButton widget.Clickable
		closeButton      widget.Clickable
		list             widget.List = widget.List{List: layout.List{Axis: layout.Vertical}}
	)
	//Creating a widget.Clickable slice of all students in DB
	students,err := s.GetAllStudents()
	if err != nil{
		log.Println("failed to get students:", err)
		return nil
	}
	var widgetList []widget.Clickable
	for range students {
		var widget widget.Clickable
		widgetList = append(widgetList, widget)
	}

	return func(gtx layout.Context) (ui.Screen, func(gtx layout.Context)) {
		// Rendering:
		// Widget drawing:
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
					ui.DrawTitle(state, theme, 70, "Students", ui.TitleColor, ui.Margins{Right: 0, Left: 0, Top: 0, Bottom: 0}),
				),
				// List:
				layout.Rigid(
					ui.DrawStudentListWithMargins(state, theme, gtx, &widgetList, students, &list, ui.Margins{Right: 0, Left: 0, Top: 0, Bottom: 175}),
				),
			)

			// Flexbox with Bottom alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart, // Bottom
			}.Layout(gtx,

				// Add Student button:
				layout.Rigid(
					ui.DrawButtonWithMargins(state, theme, &addStudentButton, "Add Student", 15, ui.Margins{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
				),
				// Close button:
				layout.Rigid(
					ui.DrawButtonWithMargins(state, theme, &closeButton, "Close", 15, ui.Margins{Right: 200, Left: 200, Top: 0, Bottom: 35}, ui.ButtonColor),
				),
			)
		}
		// Event handling:
		if closeButton.Clicked() {
			return MainMenu(state, theme, s), layout
		}
		if addStudentButton.Clicked() {
			return Add_Student(state, theme, s), layout
		}
		for index := range widgetList {
			if widgetList[index].Clicked() {
				return Edit_Student(state, theme, s, index, Students(state, theme, s)), layout
			}
		}
		return nil, layout
	}

}
