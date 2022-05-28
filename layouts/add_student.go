package layouts

import (
	"fake-klasse/state"
	"fake-klasse/storage"
	"fake-klasse/ui"
	"log"
	"strings"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func Add_Student(state *state.State, theme *material.Theme, s *storage.Storage) ui.Screen {

	// Widget declaration:
	var (
		nameWidget    widget.Editor
		surnameWidget widget.Editor

		classButton     widget.Clickable
		classButtonText string
		selectedClass   string //string = "Class"
		drawClassList   bool
		classList       widget.List = widget.List{List: layout.List{Axis: layout.Vertical}}

		saveButton  widget.Clickable
		closeButton widget.Clickable
	)
	
	nameWidget.Focus()// Places cursor in name field by default.

	classButtonText = "Class" //for button to be with class text from the start

	//Creating a widget.Clickable slice of all classes in DB
	//classes := storage.Singleton.GetAllClasses()
	var widgetList []widget.Clickable
	classes, err := s.GetAllClasses()
	if err != nil {
		log.Println("unable to get classes: ", err)
		return nil
	}
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
					ui.DrawTitle(state, theme, 70, "Add Student", ui.TitleColor, ui.Margins{Right: 0, Left: 0, Top: 0, Bottom: 0}),
				),
			)

			//Horizontal Middle Flexbox
			layout.Flex{
				Axis:    layout.Horizontal,
				Spacing: layout.SpaceAround,
			}.Layout(gtx,
				layout.Flexed(1, ui.DrawInputWithMargins(state, theme, &nameWidget, "Name", 25, ui.Margins{Right: 0, Left: 50, Top: 150, Bottom: 0})),
				layout.Flexed(1, ui.DrawInputWithMargins(state, theme, &surnameWidget, "Surname", 25, ui.Margins{Right: 50, Left: 25, Top: 150, Bottom: 0})),
			)

			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceEnd, //space sides
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
					ui.InputCheck(
						ui.DrawButtonWithMargins(state, theme, &saveButton, "Save", 15, ui.Margins{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
						nameWidget, surnameWidget, /* selectedClass.Name,*/
					),
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
						ui.DrawClassListWithMargins(state, gtx, theme, widgetList, 
							func ()[]storage.Class{
								classes, err := s.GetAllClasses()
								if err != nil {
									log.Println("unable to get classes: ", err)
									return nil
								}
								return classes
							}(),
							&classList, ui.Margins{Right: 200, Left: 210, Top: 270, Bottom: 0},
						),
					),
				)
			}

		}

		// Event handling:
		if closeButton.Clicked() {
			return Students(state, theme, s), layout
		}
		if saveButton.Clicked() {
			s.AddStudent(
				strings.TrimSpace(nameWidget.Text()),
				strings.TrimSpace(surnameWidget.Text()),
				selectedClass,
			)
			return Students(state, theme, s), layout
		}
		if classButton.Clicked() {
			if drawClassList {
				drawClassList = false
			} else {
				drawClassList = true
			}
		}
		for index := range widgetList {
			if widgetList[index].Clicked() {
				drawClassList = false

				classes, err = s.GetAllClasses()
				if err != nil {
					log.Println("unable to get classes: ", err)
					return nil, layout
				}
				selectedClass = classes[index].Name //change the text of the classButton
				classButtonText = selectedClass
			}
		}
		if nameWidget.Focused() || surnameWidget.Focused(){
			drawClassList = false
		}

		return nil, layout

	}

}
