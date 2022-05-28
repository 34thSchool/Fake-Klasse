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

func Class_Students(state *state.State, theme *material.Theme, s *storage.Storage, id int) ui.Screen {

	// Widget declaration:
	var (
		editClassButton widget.Clickable
		closeButton     widget.Clickable
		list            widget.List = widget.List{List: layout.List{Axis: layout.Vertical}}
	)

	//Creating a widget.Clickable slice of all students in DB
	var widgetList []widget.Clickable
	classes, err := s.GetAllClasses()
	if err != nil {
		log.Println("unable to get classes:", err)
		return nil
	}
	students, err := s.GetClassStudents(classes[id])
	if err != nil{
		log.Println("unable to get class students:", err)
		return nil
	}
	for range students{
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
					ui.DrawTitle(state, theme, 70, classes[id].Name, ui.TitleColor, ui.Margins{Right: 0, Left: 0, Top: 0, Bottom: 0},
					),
				),
			
				// List:
				layout.Rigid(
					ui.DrawStudentListWithMargins(state, theme, gtx, widgetList, students, &list, ui.Margins{Right: 0, Left: 0, Top: 0, Bottom: 175}),
				),
			)

			// Flexbox with Bottom alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart, // Bottom
			}.Layout(gtx,

				// Add Student button:
				layout.Rigid(
					ui.DrawButtonWithMargins(state, theme, &editClassButton, "Edit Class", 15, ui.Margins{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
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
		if editClassButton.Clicked() {
			return Edit_Class(state, theme, s, id), layout
		}
		for index := range widgetList {
			if widgetList[index].Clicked() {
				return Edit_Student(state, theme, s, index, Class_Students(state, theme, s, id)), layout
			}
		}
		return nil, layout
	}

}
