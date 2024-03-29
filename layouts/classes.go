//the file with a list of classes drawing
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

func Classes(state *state.State, theme *material.Theme, s *storage.Storage) ui.Screen {

	// Widget declaration:
	var (
		addClassButton widget.Clickable
		closeButton    widget.Clickable
		list           widget.List = widget.List{List: layout.List{Axis: layout.Vertical}}
	)
	//Creating a widget.Clickable slice of all classes in DB
	classes, err := s.GetAllClasses()
	if err != nil{
		log.Println("unable to get classes: ", err)
		return nil
	}
	var widgetList []widget.Clickable

	for range classes {
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
					ui.DrawTitle(state, theme, 70, "Classes", ui.TitleColor, ui.Margins{Right: 0, Left: 0, Top: 0, Bottom: 0}),
				),
				// List:
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
						&list, ui.Margins{Right: 0, Left: 0, Top: 0, Bottom: 175}), //Change on GetClasses()
				),
			)

			// Flexbox with Bottom alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart, // Bottom
			}.Layout(gtx,

				// Add Student button:
				layout.Rigid(
					ui.DrawButtonWithMargins(state, theme, &addClassButton, "Add Class", 15, ui.Margins{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
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
		if addClassButton.Clicked() {
			return Add_Class(state, theme, s), layout
		}
		for index := range widgetList { //iterating through all list buttons and checking if they are clicked
			if widgetList[index].Clicked() {
				return Class_Students(state, theme, s, index), layout
			}
		}
		return nil, layout
	}

}
