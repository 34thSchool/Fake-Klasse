package layouts

import (
	"fake-klasse/state"
	"fake-klasse/storage"
	"fake-klasse/ui"
	"log"
	"strings"

	//"gioui.org/gesture"
	//"gioui.org/key"
	//"gioui.org/io/key"
	//"gioui.org/io/pointer"
	"gioui.org/layout"

	//"gioui.org/op"
	//"gioui.org/key"
	//"gioui.org/event"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func Edit_Student(state *state.State, theme *material.Theme, s *storage.Storage, id int, returnLayout ui.Screen) ui.Screen {

	// Widget declaration:
	var (
		nameWidget    widget.Editor
		surnameWidget widget.Editor

		classButton   widget.Clickable
		selectedClass string //string = "Class"
		drawClassList bool
		classList     widget.List = widget.List{List: layout.List{Axis: layout.Vertical}}

		saveButton          widget.Clickable
		closeButton         widget.Clickable
		deleteStudentButton widget.Clickable
	)

	nameWidget.Focus()// Places cursor in name field by default.

	//var tag = new(bool) // We could use &pressed for this instead.
	//var pressed = false

	//students table
	students,err := s.GetAllStudents()
	if err != nil{
		log.Println("failed to get students:", err)
		return nil
	}

	selectedClass = students[id].Class
	classButtonText := selectedClass
	if selectedClass == "" {
		classButtonText = "Class"
	}

	// Classes
	classes, err := s.GetAllClasses()
	if err != nil{
		log.Println("unable to get classes: ", err)
		return nil
	}

	//Creating a widget.Clickable slice of all classes in classes table
	var widgetList []widget.Clickable
	for range classes {

		var widget widget.Clickable
		widgetList = append(widgetList, widget)
	}

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
					ui.DrawTitle(state, theme, 70, "Edit Student", ui.TitleColor, ui.Margins{Right: 0, Left: 0, Top: 0, Bottom: 0}),
				),
			)

			//Horizontal Middle Flexbox
			layout.Flex{
				Axis:    layout.Horizontal,
				Spacing: layout.SpaceAround,
			}.Layout(gtx,
				layout.Flexed(1, ui.DrawInputWithMargins(state, theme, &nameWidget, (students)[id].Name, 25, ui.Margins{Right: 0, Left: 50, Top: 150, Bottom: 0})),
				layout.Flexed(1, ui.DrawInputWithMargins(state, theme, &surnameWidget, (students)[id].Surname, 25, ui.Margins{Right: 50, Left: 25, Top: 150, Bottom: 0})),
			)

			//Vertical Middle Flexbox for classes button
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceEnd,
			}.Layout(gtx,
				layout.Rigid(
					ui.DrawButtonWithMargins(state, theme, &classButton, classButtonText, 20, ui.Margins{Right: 175, Left: 175, Top: 230, Bottom: 0}, ui.ClassButtonColor),
				),
			)

			// Flexbox with Bottom alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart, // Bottom
			}.Layout(gtx,
				// Save button:
				layout.Rigid(
					//ui.InputCheck(
					ui.DrawButtonWithMargins(state, theme, &saveButton, "Save", 15, ui.Margins{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
					//nameWidget, surnameWidget,/* selectedClass.Name,*/
					//),
				),
				// Delete Student button:
				layout.Rigid(
					ui.DrawButtonWithMargins(state, theme, &deleteStudentButton, "Delete Student", 15, ui.Margins{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
				),
				// Close button:
				layout.Rigid(
					ui.DrawButtonWithMargins(state, theme, &closeButton, "Close", 15, ui.Margins{Right: 200, Left: 200, Top: 0, Bottom: 35}, ui.ButtonColor),
				),
			)

			if drawClassList {
				layout.Flex{
					Axis:    layout.Vertical,
					Spacing: layout.SpaceEnd, //around evently sides
				}.Layout(gtx,
					layout.Rigid(
						ui.DrawClassListWithMargins(state, gtx, theme, widgetList, classes, &classList, ui.Margins{Right: 200, Left: 210, Top: 270, Bottom: 0}),
						//ui.DrawClassesPopupWithMargins(theme, gtx, s, widgetList, classes, ui.Margins{Right: 200, Left: 210, Top: 270, Bottom: 0}),
					),
				)
			}

		}

		// Event handling:
		if closeButton.Clicked() {
			return returnLayout, layout
		}
		if saveButton.Clicked() {
			
			s.UpdateStudent(students[id], storage.Student{Name: strings.TrimSpace(nameWidget.Text()), Surname: strings.TrimSpace(surnameWidget.Text()), Class: selectedClass})

			return returnLayout, layout
		}
		if deleteStudentButton.Clicked() {
			s.DeleteStudent(students[id].Rowid)
			return returnLayout, layout
		}
		if classButton.Clicked() {
			if drawClassList {
				drawClassList = false
			} else {
				drawClassList = true
			}
		}
		for index := range widgetList {
			// Whether or not any of classes are clicked:
			if widgetList[index].Clicked() {
				drawClassList = false
				selectedClass = classes[index].Name //change the text of the classButton
				classButtonText = selectedClass
			}
		}
		// When we point in other place but our class list, we need class list to close.
		// When we point in other place we select name or surname fields. Yah, some crutchy crutchezz.
		if nameWidget.Focused() || surnameWidget.Focused(){
			drawClassList = false
		}

		return nil, layout

	}
}
