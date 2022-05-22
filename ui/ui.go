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

	"fake-klasse/storage"
)

// Strange (recursive?) data type macro:
type Screen func(gtx layout.Context) (Screen, func(graphicalContext layout.Context))

// Color literals:
var(
	ButtonColor = color.NRGBA{64, 64, 64, 255}//light grey
	TitleColor = color.NRGBA{55, 55, 55, 255}//dark grey
	BackgroundColor = color.NRGBA{25, 25, 25, 255}//black
	TextColor = color.NRGBA{255,255,255,255}//white
	LightListColor = color.NRGBA{75, 75, 75, 255}// light grey for list widget elements
	DarkListColor = color.NRGBA{65, 65, 65, 255}// dark grey for list widget elements
	HintColor = color.NRGBA{150, 150, 150, 255}// grey
)

//Margin dimentions
type Rect struct { // To define margins more conviently.
	Left, Right, Top, Bottom float32
}


//Buttons:
func DrawButtonWithMargins(theme *material.Theme, button *widget.Clickable, text string, textSize float32, r Rect, c color.NRGBA) func(graphicalContext layout.Context) layout.Dimensions {

	return func(graphicalContext layout.Context) layout.Dimensions {
		margins := layout.Inset{
			Top:    unit.Dp(r.Top),
			Bottom: unit.Dp(r.Bottom),
			Left:   unit.Dp(r.Left),
			Right:  unit.Dp(r.Right),
		}
		return margins.Layout(graphicalContext,
			DrawButton(theme, button, text, textSize, c),
		)
	}
}
func DrawButton(theme *material.Theme, button *widget.Clickable, text string, textSize float32, c color.NRGBA) func(graphicalContext layout.Context) layout.Dimensions {

	return func(graphicalContext layout.Context) layout.Dimensions {
		button := material.Button(theme, button, text)
		button.Background = c
		button.TextSize = unit.Dp(textSize)
		return button.Layout(graphicalContext) // returns dimensions
	}
}

// Greyes out a button:
func InputCheck (w layout.Widget, firstInput, secondInput widget.Editor) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		firstInput := strings.TrimSpace(firstInput.Text())
		secondInput := strings.TrimSpace(secondInput.Text())
		if firstInput == "" || secondInput == "" { 
			gtx = gtx.Disabled()//If one of them is empty than the button is disabled.
		}
		return w(gtx)
	}
}

// Draws title in the rectangle:
func DrawTitle(theme *material.Theme, operations *op.Ops, rectangleHeight float32, titleText string, c color.NRGBA, r Rect) func(graphicalContext layout.Context) layout.Dimensions{

	return func(graphics_context layout.Context) layout.Dimensions {

		graphics_context.Constraints.Min.Y = graphics_context.Px(unit.Dp(rectangleHeight))

		top_rectangle := image.Rectangle{Max: graphics_context.Constraints.Min}
		paint.FillShape(operations, c, clip.Rect(top_rectangle).Op())
		

		title := material.Label(theme, unit.Dp(40), titleText)
		title.Alignment = text.Middle
		title.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
		return layout.Inset{Right: unit.Dp(r.Right), Left: unit.Dp(r.Left), Top: unit.Dp(r.Top), Bottom: unit.Dp(r.Bottom)}.Layout(graphics_context, title.Layout)
	}
}
func DrawBackground(operations *op.Ops, c color.NRGBA) func(graphicalContext layout.Context) layout.Dimensions{
	 paint.Fill(operations, c)
	 return nil
}


// List:
func DrawListWithMargins(graphicalContext layout.Context, theme *material.Theme, buttons *[]widget.Clickable, students *[]storage.Student, list *widget.List, r Rect) func(graphicalContext layout.Context) layout.Dimensions {

	return func(graphicalContext layout.Context) layout.Dimensions {
		margins := layout.Inset{
			Top:    unit.Dp(r.Top),
			Bottom: unit.Dp(r.Bottom),
			Left:   unit.Dp(r.Left),
			Right:  unit.Dp(r.Right),
		}
		return margins.Layout(graphicalContext,
			DrawList(graphicalContext, theme, buttons, students, list),
		)
	}
}
func DrawList(graphicalContext layout.Context, theme *material.Theme, buttons *[]widget.Clickable, students *[]storage.Student, list *widget.List) func(graphicalContext layout.Context) layout.Dimensions {

	return func(graphicalContext layout.Context) layout.Dimensions {

		return material.List(theme, list).Layout(graphicalContext, len(*students), func(graphicalContext layout.Context, index int) layout.Dimensions {
			var button *widget.Clickable = &(*buttons)[index]
			return DrawListElement(graphicalContext, theme, button, students, index)
		})
	}
}
func DrawListElement(graphicalContext layout.Context, theme *material.Theme, button *widget.Clickable, students *[]storage.Student, index int) layout.Dimensions{
	student := (*students)[index]

	var color color.NRGBA
	if index%2 == 0 {
		color = DarkListColor
	} else {
		color = LightListColor
	}

	return DrawButton(theme, button, fmt.Sprintf("%s %s", student.Name, student.Surname), 15, color)(graphicalContext)
}
func DrawListWidget(graphicalContext layout.Context, index int) func (graphicalContext layout.Context) layout.Dimensions{

	return func (graphicalContext layout.Context) layout.Dimensions{
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


//Input:
func DrawInput(theme *material.Theme, input *widget.Editor, hint string, textSize float32) func(graphicalContext layout.Context) layout.Dimensions {

	return func(graphicalContext layout.Context) layout.Dimensions {
		input := material.Editor(theme, input, hint)
		input.TextSize = unit.Dp(textSize)
		input.HintColor = HintColor
		input.Color = TextColor
		input.SelectionColor = LightListColor//very useful thing ._.
		return input.Layout(graphicalContext)
	}
}
func DrawInputWithMargins(theme *material.Theme, input *widget.Editor, hint string, textSize float32, r Rect) func(graphicalContext layout.Context) layout.Dimensions{
	return func(graphicalContext layout.Context) layout.Dimensions{
		margins := layout.Inset{
			Top:    unit.Dp(r.Top),
			Bottom: unit.Dp(r.Bottom),
			Left:   unit.Dp(r.Left),
			Right:  unit.Dp(r.Right),
		}
	return margins.Layout(graphicalContext,
		DrawInput(theme, input, hint, textSize),
	)
	}
}