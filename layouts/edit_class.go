package layouts

import (
	"fake-klasse/state"
	"fake-klasse/storage"
	"fake-klasse/ui"
	"strings"

	"gioui.org/layout"
	"gioui.org/widget"
)

func Edit_Class(state *state.State, s *storage.Storage, index int) ui.Screen {

	// Widget declaration:
	var (
		classWidget widget.Editor
		className   string

		saveButton        widget.Clickable
		closeButton       widget.Clickable
		deleteClassButton widget.Clickable
	)

	className = (*s.GetAllClasses())[index].Name

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
					ui.DrawTitle(state, 70, "Edit Class", ui.TitleColor, ui.Margins{Right: 0, Left: 0, Top: 0, Bottom: 0}),
				),
				//Input
				layout.Flexed(1,
					ui.DrawInputWithMargins(state, &classWidget, className, 25, ui.Margins{Right: 300, Left: 300, Top: 150, Bottom: 0}),
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
					ui.DrawButtonWithMargins(state, &saveButton, "Save", 15, ui.Margins{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
					//classWidget,
					//),
				),
				// Delete Student button:
				layout.Rigid(
					ui.DrawButtonWithMargins(state, &deleteClassButton, "Delete Class", 15, ui.Margins{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
				),
				// Close button:
				layout.Rigid(
					ui.DrawButtonWithMargins(state, &closeButton, "Close", 15, ui.Margins{Right: 200, Left: 200, Top: 0, Bottom: 35}, ui.ButtonColor),
				),
			)
		}

		// Event handling:
		if closeButton.Clicked() {
			return Class_Students(state, s, index), layout
		}
		if saveButton.Clicked() {

			className := strings.TrimSpace(classWidget.Text())

			if className != ""{ // If user entered smth
				s.UpdateAllClassStudentsClass((*s.GetAllClasses())[index].Name, className)
				
				s.DeleteClass((*s.GetAllClasses())[index])
	
				s.AddClass(
					strings.TrimSpace(classWidget.Text()),
				)
			}

			return Class_Students(state, s, index), layout
		}
		if deleteClassButton.Clicked() {

			s.UpdateAllClassStudentsClass((*s.GetAllClasses())[index].Name, "")
			s.DeleteClass((*s.GetAllClasses())[index])
			className = ""

			return Classes(state, s), layout
		}

		return nil, layout

	}
}
