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

		classButton     widget.Clickable
		classButtonText string
		selectedClass   string //string = "Class"
		drawClassList   bool
		classList       widget.List = widget.List{List: layout.List{Axis: layout.Vertical}}

		saveButton  widget.Clickable
		closeButton widget.Clickable
	)

	classButtonText = "Class" //for button to be with class text from the start

	//Creating a widget.Clickable slice of all classes in DB
	//classes := storage.Singleton.GetAllClasses()
	var widgetList []widget.Clickable
	for range *storage.Singleton.GetAllClasses() {
		var widget widget.Clickable
		widgetList = append(widgetList, widget)
	}

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
				layout.Flexed(1, ui.DrawInputWithMargins(&nameWidget, "Name", 25, ui.Rect{Right: 0, Left: 50, Top: 150, Bottom: 0})),
				layout.Flexed(1, ui.DrawInputWithMargins(&surnameWidget, "Surname", 25, ui.Rect{Right: 50, Left: 25, Top: 150, Bottom: 0})),
			)

			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceEnd, //space sides
			}.Layout(graphicalContext,
				layout.Rigid(
					ui.DrawButtonWithMargins(&classButton, classButtonText, 20, ui.Rect{Right: 175, Left: 175, Top: 230, Bottom: 0}, ui.ClassButtonColor),
				),
			)

			// Flexbox with Bottom alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart, // Bottom
			}.Layout(graphicalContext,

				// Save button:
				layout.Rigid(
					ui.InputCheck(
						ui.DrawButtonWithMargins(&saveButton, "Save", 15, ui.Rect{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
						nameWidget, surnameWidget, /* selectedClass.Name,*/
					),
				),
				// Close button:
				layout.Rigid(
					ui.DrawButtonWithMargins(&closeButton, "Close", 15, ui.Rect{Right: 200, Left: 200, Top: 0, Bottom: 35}, ui.ButtonColor),
				),
			)

			if drawClassList {
				layout.Flex{
					Axis:    layout.Vertical,
					Spacing: layout.SpaceEnd, //around evently sides
				}.Layout(graphicalContext,
					layout.Rigid(
						ui.DrawClassListWithMargins(graphicalContext, &widgetList, storage.Singleton.GetAllClasses(), &classList, ui.Rect{Right: 200, Left: 210, Top: 270, Bottom: 0}),
					),
				)
			}

		}

		// Event handling:
		if closeButton.Clicked() {
			return Students(), layout
		}
		if saveButton.Clicked() {
			storage.Singleton.AddStudent(
				strings.TrimSpace(nameWidget.Text()),
				strings.TrimSpace(surnameWidget.Text()),
				selectedClass,
			)
			return Students(), layout
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
				selectedClass = (*storage.Singleton.GetAllClasses())[index].Name //change the text of the classButton
				classButtonText = selectedClass
			}
		}

		return nil, layout

	}

}
