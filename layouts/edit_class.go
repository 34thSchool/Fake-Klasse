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

func Edit_Class(state *state.State, theme *material.Theme, s *storage.Storage, index int) ui.Screen {

	// Widget declaration:
	var (
		classWidget widget.Editor
		className   string

		saveButton        widget.Clickable
		closeButton       widget.Clickable
		deleteClassButton widget.Clickable
	)

	classWidget.Focus()// Places cursor in class name field by default.

	classes, err := s.GetAllClasses()
	if err != nil{
		log.Println("unable to get classes: ", err)
		return nil
	}
	className = classes[index].Name

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
					ui.DrawTitle(state, theme, 70, "Edit Class", ui.TitleColor, ui.Margins{Right: 0, Left: 0, Top: 0, Bottom: 0}),
				),
				//Input
				layout.Flexed(1,
					ui.DrawInputWithMargins(state, theme, &classWidget, className, 25, ui.Margins{Right: 300, Left: 300, Top: 150, Bottom: 0}),
				),
			)

			// Flexbox with Bottom alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart, // Bottom
			}.Layout(gtx,
				// Save button:
				layout.Rigid(
					//ui.ClassInputCheck(
					ui.DrawButtonWithMargins(state, theme, &saveButton, "Save", 15, ui.Margins{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
					//classWidget,
					//),
				),
				// Delete Student button:
				layout.Rigid(
					ui.DrawButtonWithMargins(state, theme, &deleteClassButton, "Delete Class", 15, ui.Margins{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
				),
				// Close button:
				layout.Rigid(
					ui.DrawButtonWithMargins(state, theme, &closeButton, "Close", 15, ui.Margins{Right: 200, Left: 200, Top: 0, Bottom: 35}, ui.ButtonColor),
				),
			)
		}

		// Event handling:
		if closeButton.Clicked() {
			return Class_Students(state, theme, s, index), layout
		}
		if saveButton.Clicked() {

			className := strings.TrimSpace(classWidget.Text())

			if className != ""{ // If user entered smth
				s.UpdateAllClassStudentsClass(classes[index].Name, className)
				
				s.DeleteClass(classes[index])
	
				s.AddClass(
					strings.TrimSpace(classWidget.Text()),
				)
			}

			return Class_Students(state, theme, s, index), layout
		}
		if deleteClassButton.Clicked() {

			s.UpdateAllClassStudentsClass(classes[index].Name, "")
			s.DeleteClass(classes[index])
			className = ""

			return Classes(state, theme, s), layout
		}

		return nil, layout

	}
}
