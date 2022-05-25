//The file which contains a widget drawing functions and everything related to ui
package ui

import (
	"fmt"
	"image"
	"image/color"
	"strings"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"fake-klasse/state"
	"fake-klasse/storage"
)

// Strange (recursive?) data type macro:
type Screen func(gtx layout.Context) (Screen, func(graphicalContext layout.Context))

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
type Rect struct { // To define margins more conviently.
	Left, Right, Top, Bottom float32
}

//Buttons:
func DrawButtonWithMargins(button *widget.Clickable, text string, textSize float32, r Rect, c color.NRGBA) func(graphicalContext layout.Context) layout.Dimensions {

	return func(graphicalContext layout.Context) layout.Dimensions {
		margins := layout.Inset{
			Top:    unit.Dp(r.Top),
			Bottom: unit.Dp(r.Bottom),
			Left:   unit.Dp(r.Left),
			Right:  unit.Dp(r.Right),
		}
		return margins.Layout(graphicalContext,
			DrawButton(button, text, textSize, c),
		)
	}
}
func DrawButton(button *widget.Clickable, text string, textSize float32, c color.NRGBA) func(graphicalContext layout.Context) layout.Dimensions {

	return func(graphicalContext layout.Context) layout.Dimensions {
		button := material.Button(state.Theme, button, text)
		button.Background = c
		button.TextSize = unit.Dp(textSize)
		return button.Layout(graphicalContext) // returns dimensions
	}
}

//Inputs:
// Greyes out a button:
func InputCheck(w layout.Widget, nameInput widget.Editor, surnameInput widget.Editor /*, className string */) layout.Widget {
	return func(graphicalContext layout.Context) layout.Dimensions {
		nameInput := strings.TrimSpace(nameInput.Text())
		surnameInput := strings.TrimSpace(surnameInput.Text())
		if nameInput == "" || surnameInput == "" /* || className == "Class"*/ {
			graphicalContext = graphicalContext.Disabled() //If one of them is empty than the button is disabled.
		}
		return w(graphicalContext)
	}
}
func ClassInputCheck(w layout.Widget, classInput widget.Editor) layout.Widget {
	return func(graphicalContext layout.Context) layout.Dimensions {
		classInput := strings.TrimSpace(classInput.Text())
		if classInput == "" {
			graphicalContext = graphicalContext.Disabled()
		}
		return w(graphicalContext)
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
func DrawTitle(rectangleHeight float32, titleText string, c color.NRGBA, r Rect) func(graphicalContext layout.Context) layout.Dimensions {

	return func(graphics_context layout.Context) layout.Dimensions {

		graphics_context.Constraints.Min.Y = graphics_context.Px(unit.Dp(rectangleHeight))

		top_rectangle := image.Rectangle{Max: graphics_context.Constraints.Min}
		paint.FillShape(graphics_context.Ops, c, clip.Rect(top_rectangle).Op())

		title := material.Label(state.Theme, unit.Dp(40), titleText)
		title.Alignment = text.Middle
		title.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
		return layout.Inset{Right: unit.Dp(r.Right), Left: unit.Dp(r.Left), Top: unit.Dp(r.Top), Bottom: unit.Dp(r.Bottom)}.Layout(graphics_context, title.Layout)
	}
}
func DrawBackground(operations *op.Ops, c color.NRGBA) {
	paint.Fill(operations, c) //op.Ops{} doesn't work
}

// Student list:
func DrawStudentListWithMargins(graphicalContext layout.Context, buttons *[]widget.Clickable, students *[]storage.Student, list *widget.List, r Rect) func(graphicalContext layout.Context) layout.Dimensions {

	return func(graphicalContext layout.Context) layout.Dimensions {
		margins := layout.Inset{
			Top:    unit.Dp(r.Top),
			Bottom: unit.Dp(r.Bottom),
			Left:   unit.Dp(r.Left),
			Right:  unit.Dp(r.Right),
		}
		return margins.Layout(graphicalContext,
			DrawStudentList(graphicalContext, buttons, students, list),
		)
	}
}
func DrawStudentList(graphicalContext layout.Context, buttons *[]widget.Clickable, students *[]storage.Student, list *widget.List) func(graphicalContext layout.Context) layout.Dimensions {

	return func(graphicalContext layout.Context) layout.Dimensions {

		return material.List(state.Theme, list).Layout(graphicalContext, len(*students), func(graphicalContext layout.Context, index int) layout.Dimensions {
			var button *widget.Clickable = &(*buttons)[index]
			return DrawStudentListElement(graphicalContext, button, students, index)
		})
	}
}
func DrawStudentListElement(graphicalContext layout.Context, button *widget.Clickable, students *[]storage.Student, index int) layout.Dimensions {

	student := (*students)[index]

	var color color.NRGBA
	if index%2 == 0 {
		color = DarkListColor
	} else {
		color = LightListColor
	}

	return DrawButton(button, fmt.Sprintf("%s %s %s", student.Name, student.Surname, student.Class), 15, color)(graphicalContext)
}
func DrawListWidget(graphicalContext layout.Context, index int) func(graphicalContext layout.Context) layout.Dimensions {

	return func(graphicalContext layout.Context) layout.Dimensions {
		var color color.NRGBA
		if index%2 == 0 {
			color = DarkListColor
		} else {
			color = LightListColor
		}
		max := image.Pt(graphicalContext.Constraints.Max.X, graphicalContext.Constraints.Min.Y)
		paint.FillShape(graphicalContext.Ops, color, clip.Rect{Max: max}.Op())
		return layout.Dimensions{Size: graphicalContext.Constraints.Min}
	}
}

// Class list:
func DrawClassListWithMargins(graphicalContext layout.Context, buttons *[]widget.Clickable, classes *[]storage.Class, list *widget.List, r Rect) func(graphicalContext layout.Context) layout.Dimensions {

	return func(graphicalContext layout.Context) layout.Dimensions {
		margins := layout.Inset{
			Top:    unit.Dp(r.Top),
			Bottom: unit.Dp(r.Bottom),
			Left:   unit.Dp(r.Left),
			Right:  unit.Dp(r.Right),
		}
		return margins.Layout(graphicalContext,
			DrawClassList(graphicalContext, buttons, classes, list),
		)
	}
}
func DrawClassList(graphicalContext layout.Context, buttons *[]widget.Clickable, classes *[]storage.Class, list *widget.List) func(graphicalContext layout.Context) layout.Dimensions {

	return func(graphicalContext layout.Context) layout.Dimensions {

		return material.List(state.Theme, list).Layout(graphicalContext, len(*classes), func(graphicalContext layout.Context, index int) layout.Dimensions {
			var button *widget.Clickable = &(*buttons)[index]
			return DrawClassListElement(graphicalContext, button, classes, index)
		})
	}
}
func DrawClassListElement(graphicalContext layout.Context, button *widget.Clickable, classes *[]storage.Class, index int) layout.Dimensions {

	class := (*classes)[index]

	var color color.NRGBA
	if index%2 == 0 {
		color = DarkListColor
	} else {
		color = LightListColor
	}

	return DrawButton(button, class.Name, 15, color)(graphicalContext)
}

//Input:
func DrawInput(input *widget.Editor, hint string, textSize float32) func(graphicalContext layout.Context) layout.Dimensions {

	return func(graphicalContext layout.Context) layout.Dimensions {
		input := material.Editor(state.Theme, input, hint)
		input.TextSize = unit.Dp(textSize)
		input.HintColor = HintColor
		input.Color = TextColor
		input.SelectionColor = LightListColor //very useful thing ._.
		return input.Layout(graphicalContext)
	}
}
func DrawInputWithMargins(input *widget.Editor, hint string, textSize float32, r Rect) func(graphicalContext layout.Context) layout.Dimensions {
	return func(graphicalContext layout.Context) layout.Dimensions {
		margins := layout.Inset{
			Top:    unit.Dp(r.Top),
			Bottom: unit.Dp(r.Bottom),
			Left:   unit.Dp(r.Left),
			Right:  unit.Dp(r.Right),
		}
		return margins.Layout(graphicalContext,
			DrawInput(input, hint, textSize),
		)
	}
}
