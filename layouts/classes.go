//the file with a list of classes drawing
package layouts

import (
	"fake-klasse/storage"
	"fake-klasse/ui"

	"gioui.org/layout"
	"gioui.org/widget"
)

func Classes() ui.Screen {

	// Widget declaration:
	var (
		addClassButton widget.Clickable
		closeButton      widget.Clickable
		list             widget.List = widget.List{List: layout.List{Axis: layout.Vertical}}
	)
	//Creating a widget.Clickable slice of all classes in DB
	classes := storage.Singleton.GetAllClasses()
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
					ui.DrawTitle(70, "Classes", ui.TitleColor, ui.Rect{Right: 0, Left: 0, Top: 0, Bottom: 0}),
				),
				// List:
				layout.Rigid(
					ui.DrawClassListWithMargins(graphicalContext, &widgetList, storage.Singleton.GetAllClasses(), &list, ui.Rect{Right: 0, Left: 0, Top: 0, Bottom: 175}),//Change on GetClasses() 
				),
			)

			// Flexbox with Bottom alignment:
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart, // Bottom
			}.Layout(graphicalContext,

				// Add Student button:
				layout.Rigid(
					ui.DrawButtonWithMargins(&addClassButton, "Add Class", 15, ui.Rect{Right: 175, Left: 175, Top: 0, Bottom: 25}, ui.ButtonColor),
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
		if addClassButton.Clicked() {
			return Add_Class(), layout
		}
		for index := range widgetList { //iterating through all list buttons and checking if they are clicked
			if widgetList[index].Clicked() {
				return Class_Students(index), layout
			}
		}
		return nil, layout
	}

}
