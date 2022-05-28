//The file which contains a widget drawing functions and everything related to ui
package ui

import (
	"fmt"
	"image"
	"image/color"
	//"log"
	"strings"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	//"gioui.org/x/component"

	"fake-klasse/state"
	"fake-klasse/storage"
)

// Strange (recursive?) data type macro:
type Screen func(gtx layout.Context) (Screen, func(gtx layout.Context))

// Color literals:
var (
	ButtonColor      = color.NRGBA{64, 64, 64, 255}    //light grey
	TitleColor       = color.NRGBA{55, 55, 55, 255}    //dark grey
	BackgroundColor  = color.NRGBA{25, 25, 25, 255}    //black
	TextColor        = color.NRGBA{255, 255, 255, 255} //white
	LightListColor   = color.NRGBA{75, 75, 75, 255}    // light grey for list widget elements
	DarkListColor    = color.NRGBA{65, 65, 65, 255}    // dark grey for list widget elements
	HintColor        = color.NRGBA{150, 150, 150, 255} // grey
	ClassButtonColor = color.NRGBA{35, 35, 35, 255}    //very dark grey (almost background)
)

//Margin dimentions
type Margins struct { // To define margins more conviently.
	Left, Right, Top, Bottom float32
}

//Buttons:
func DrawButtonWithMargins(state *state.State, theme *material.Theme, button *widget.Clickable, text string, textSize float32, m Margins, c color.NRGBA) func(gtx layout.Context) layout.Dimensions {

	return func(gtx layout.Context) layout.Dimensions {
		margins := layout.Inset{
			Top:    unit.Dp(m.Top),
			Bottom: unit.Dp(m.Bottom),
			Left:   unit.Dp(m.Left),
			Right:  unit.Dp(m.Right),
		}
		return margins.Layout(gtx,
			DrawButton(state, theme, button, text, textSize, c),
		)
	}
}
func DrawButton(state *state.State, theme *material.Theme, button *widget.Clickable, text string, textSize float32, c color.NRGBA) func(gtx layout.Context) layout.Dimensions {

	return func(gtx layout.Context) layout.Dimensions {
		button := material.Button(theme, button, text)
		button.Background = c
		button.TextSize = unit.Dp(textSize)
		return button.Layout(gtx) // returns dimensions
	}
}

//Inputs:
// Greyes out a button:
func InputCheck(w layout.Widget, nameInput widget.Editor, surnameInput widget.Editor /*, className string */) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		nameInput := strings.TrimSpace(nameInput.Text())
		surnameInput := strings.TrimSpace(surnameInput.Text())
		if nameInput == "" || surnameInput == "" /* || className == "Class"*/ {
			gtx = gtx.Disabled() //If one of them is empty than the button is disabled.
		}
		return w(gtx)
	}
}
func ClassInputCheck(w layout.Widget, classInput widget.Editor) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		classInput := strings.TrimSpace(classInput.Text())
		if classInput == "" {
			gtx = gtx.Disabled()
		}
		return w(gtx)
	}
}
func DataCheck(oldText, newText string) string {
	if newText == "" {
		return oldText
	} else {
		return newText
	}
}

// Draws title in the rectangle:
func DrawTitle(state *state.State, theme *material.Theme, rectangleHeight float32, titleText string, c color.NRGBA, m Margins) func(gtx layout.Context) layout.Dimensions {

	return func(graphics_context layout.Context) layout.Dimensions {

		graphics_context.Constraints.Min.Y = graphics_context.Px(unit.Dp(rectangleHeight))

		top_rectangle := image.Rectangle{Max: graphics_context.Constraints.Min}
		paint.FillShape(graphics_context.Ops, c, clip.Rect(top_rectangle).Op())

		title := material.Label(theme, unit.Dp(40), titleText)
		title.Alignment = text.Middle
		title.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
		return layout.Inset{Right: unit.Dp(m.Right), Left: unit.Dp(m.Left), Top: unit.Dp(m.Top), Bottom: unit.Dp(m.Bottom)}.Layout(graphics_context, title.Layout)
	}
}
func DrawBackground(operations *op.Ops, c color.NRGBA) {
	paint.Fill(operations, c) //op.Ops{} doesn't work
}

//Input:
func DrawInput(state *state.State, theme *material.Theme, input *widget.Editor, hint string, textSize float32) func(gtx layout.Context) layout.Dimensions {

	return func(gtx layout.Context) layout.Dimensions {
		inputWidget := material.Editor(theme, input, hint)
		input.SingleLine = true // Forces the box to always be one line high. Without it, the box will grow when user presses enter key.
		//input.Focus()
		inputWidget.TextSize = unit.Dp(textSize)
		inputWidget.HintColor = HintColor
		inputWidget.Color = TextColor
		inputWidget.SelectionColor = LightListColor //very useful thing ._.
		return inputWidget.Layout(gtx)
	}
}
func DrawInputWithMargins(state *state.State, theme *material.Theme, input *widget.Editor, hint string, textSize float32, m Margins) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		margins := layout.Inset{
			Top:    unit.Dp(m.Top),
			Bottom: unit.Dp(m.Bottom),
			Left:   unit.Dp(m.Left),
			Right:  unit.Dp(m.Right),
		}
		return margins.Layout(gtx,
			DrawInput(state, theme, input, hint, textSize),
		)
	}
}

// Student list:
func DrawStudentListWithMargins(state *state.State, theme *material.Theme, gtx layout.Context, buttons []widget.Clickable, students []storage.Student, list *widget.List, m Margins) func(gtx layout.Context) layout.Dimensions {

	return func(gtx layout.Context) layout.Dimensions {
		margins := layout.Inset{
			Top:    unit.Dp(m.Top),
			Bottom: unit.Dp(m.Bottom),
			Left:   unit.Dp(m.Left),
			Right:  unit.Dp(m.Right),
		}
		return margins.Layout(gtx,
			DrawStudentList(state, theme, gtx, buttons, students, list),
		)
	}
}
func DrawStudentList(state *state.State, theme *material.Theme, gtx layout.Context, buttons []widget.Clickable, students []storage.Student, list *widget.List) func(gtx layout.Context) layout.Dimensions {

	return func(gtx layout.Context) layout.Dimensions {

		return material.List(theme, list).Layout(gtx, len(students), func(gtx layout.Context, index int) layout.Dimensions {
			var button *widget.Clickable = &buttons[index]
			return DrawStudentListElement(state, theme, gtx, button, students, index)
		})
	}
}
func DrawStudentListElement(state *state.State, theme *material.Theme, gtx layout.Context, button *widget.Clickable, students []storage.Student, index int) layout.Dimensions {

	student := (students)[index]

	var color color.NRGBA
	if index%2 == 0 {
		color = DarkListColor
	} else {
		color = LightListColor
	}

	return DrawButton(state, theme, button, fmt.Sprintf("%s %s %s", student.Name, student.Surname, student.Class), 15, color)(gtx)
}
func DrawListWidget(gtx layout.Context, index int) func(gtx layout.Context) layout.Dimensions {

	return func(gtx layout.Context) layout.Dimensions {
		var color color.NRGBA
		if index%2 == 0 {
			color = DarkListColor
		} else {
			color = LightListColor
		}
		max := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Min.Y)
		paint.FillShape(gtx.Ops, color, clip.Rect{Max: max}.Op())
		return layout.Dimensions{Size: gtx.Constraints.Min}
	}
}

// Class list:
func DrawClassListWithMargins(state *state.State, gtx layout.Context, theme *material.Theme, buttons []widget.Clickable, classes []storage.Class, list *widget.List, m Margins) func(gtx layout.Context) layout.Dimensions {

	return func(gtx layout.Context) layout.Dimensions {
		margins := layout.Inset{
			Top:    unit.Dp(m.Top),
			Bottom: unit.Dp(m.Bottom),
			Left:   unit.Dp(m.Left),
			Right:  unit.Dp(m.Right),
		}
		return margins.Layout(gtx,
			DrawClassList(state, theme, gtx, buttons, classes, list),
		)
	}
}
func DrawClassList(state *state.State, theme *material.Theme, gtx layout.Context, buttons []widget.Clickable, classes []storage.Class, list *widget.List) func(gtx layout.Context) layout.Dimensions {

	return func(gtx layout.Context) layout.Dimensions {
		return material.List(theme, list).Layout(gtx, len(classes), 
			func(gtx layout.Context, index int) layout.Dimensions {
				var button *widget.Clickable = &buttons[index]
				return DrawClassListElement(state, gtx, theme, button, classes, index)
			},
		)
	}
}
func DrawClassListElement(state *state.State, gtx layout.Context, theme *material.Theme, button *widget.Clickable, classes []storage.Class, index int) layout.Dimensions {

	class := (classes)[index]

	var color color.NRGBA
	if index % 2 == 0 {
		color = DarkListColor
	} else {
		color = LightListColor
	}

	return DrawButton(state, theme, button, class.Name, 15, color)(gtx)
}


// // Menu:
// func DrawClassesPopupWithMargins(theme *material.Theme, gtx layout.Context, s *storage.Storage, buttons []widget.Clickable, classes []storage.Class, m Margins) func(gtx layout.Context) layout.Dimensions {
// 	return func(gtx layout.Context) layout.Dimensions {
// 		margins := layout.Inset{
// 			Top:    unit.Dp(m.Top),
// 			Bottom: unit.Dp(m.Bottom),
// 			Left:   unit.Dp(m.Left),
// 			Right:  unit.Dp(m.Right),
// 		}
// 		return margins.Layout(gtx,
// 			DrawClassesPopup(theme, gtx, s, buttons, classes),
// 		)
// 	}
// }
// func DrawClassesPopup(theme *material.Theme, gtx layout.Context, s *storage.Storage, buttons []widget.Clickable, classes []storage.Class) func(gtx layout.Context) layout.Dimensions {
	
// 	return func(gtx layout.Context) layout.Dimensions {
// 		// Option list:
// 		list := layout.List{
// 			Axis: layout.Horizontal,
// 			ScrollToEnd: false,
// 			Alignment: layout.Middle,
// 		}

// 		// Options:
// 		var items []func(gtx layout.Context) layout.Dimensions

		

// 		for index := range classes{
// 			var button *widget.Clickable = &buttons[index]
// 			items = append(items, ClassPopupElement(gtx, index, theme, s, button))
// 		}


// 		menuState := component.MenuState{
// 			OptionList: list,
// 			Options: items,
// 		}

// 		menu := component.Menu(theme, &menuState)
		

// 		return menu.Layout(gtx)
// 	}
	
// }

// func ClassPopupElement(gtx layout.Context, index int, theme *material.Theme, s *storage.Storage, button *widget.Clickable) func(gtx layout.Context) layout.Dimensions{
// 	// Getting classes:
// 	classes, err := s.GetAllClasses()
// 	if err != nil{log.Fatal("unable to get classes ", err)}
// 	class := classes[index]

// 	// Coloring element:
// 	if index % 2 == 0 {
// 		theme.Bg = LightListColor
// 	} else {
// 		theme.Bg = DarkListColor
// 	}
// 	//style := material.ButtonStyle{Color: color.NRGBA{255,255,255,255}}
	
// 	//style.Layout(gtx)

// 	item := component.MenuItem(theme, button, class.Name)
// 	item.HoverColor = ButtonColor

// 	return item.Layout
// }