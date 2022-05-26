//the file with a list of classes drawing
package layouts

import (
	"fake-klasse/state"
	"fake-klasse/storage"
	"fake-klasse/ui"

	"gioui.org/layout"
	"gioui.org/widget"
)

func Classes(state *state.State, s *storage.Storage) ui.Screen {

	// Widget declaration:
	var (
		addClassButton widget.Clickable
		closeButton    widget.Clickable
		list           widget.List = widget.List{List: layout.List{Axis: layout.Vertical}}
	)
	//Creating a widget.Clickable slice of all classes in DB
	classes := s.GetAllClasses()
	var widgetList []widget.Clickable

	for range *classes {
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
					ui.DrawTitle(state, 70, "Classes", ui.TitleColor, ui.Margins{Right: 0, Left: 0, Top: 0, Bottom: 0}),
				),
				// List:
				layout.Rigid(
					ui.DrawClassListWithMargins(state, graphicalContext, &widgetList, s.GetAllClasses(), &list, ui.Margins{Right: 0, Left: 0, Top: 0, Bottom: 175}), //Change on GetClasses()
				),
			)

			// Flexbox with Bottom alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart, // Bottom
			}.Layout(graphicalContext,

				// Add Student button:
				layout.Rigid(
					ui.DrawButtonWithMargins(state, &addClassButton, "Add Class", 15, ui.Margins{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
				),
				// Close button:
				layout.Rigid(
					ui.DrawButtonWithMargins(state, &closeButton, "Close", 15, ui.Margins{Right: 200, Left: 200, Top: 0, Bottom: 35}, ui.ButtonColor),
				),
			)
		}
		// Event handling:
		if closeButton.Clicked() {
			return MainMenu(state, s), layout
		}
		if addClassButton.Clicked() {
			return Add_Class(state, s), layout
		}
		for index := range widgetList { //iterating through all list buttons and checking if they are clicked
			if widgetList[index].Clicked() {
				return Class_Students(state, s, index), layout
			}
		}
		return nil, layout
	}

}
