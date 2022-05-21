package ui

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Rect struct {
	left, right, top, bottom float32
}
type Color struct {
	R, G, B, A uint8
}

func MainMenu(theme *material.Theme, operations *op.Ops) func(graphicalContext layout.Context) {

	var studentsButton widget.Clickable
	var quit widget.Clickable

	// Drawing widgets:
	return func(graphicalContext layout.Context) {

		// Drawing background:
		paint.Fill(operations, color.NRGBA{R: 25, G: 25, B: 25, A: 255})

		layout.Flex{
			Axis:    layout.Vertical,
			Spacing: layout.SpaceStart,
		}.Layout(graphicalContext,
			layout.Rigid(
				DrawTitle(theme, operations, 80, "Nach der Schule", 45, Color{55, 55, 55, 255}, Rect{0, 0, 0, 220}), //rect 220
			),
			layout.Rigid(
				DrawButtonWithMargins(theme, &studentsButton, "Students", 15, Rect{150, 150, 0, 250}, Color{100, 200, 150, 255}),
			),
			layout.Rigid(
				DrawButtonWithMargins(theme, &quit, "Quit", 15, Rect{200, 200, 0, 175}, Color{197, 95, 95, 255}),
			),
		)
	}
}

func DrawButtonWithMargins(theme *material.Theme, button *widget.Clickable, text string, textSize float32, r Rect, c Color) func(graphicalContext layout.Context) layout.Dimensions {

	return func(graphicalContext layout.Context) layout.Dimensions {
		margins := layout.Inset{
			// UniformInset() sets all the sides to one value.
			Top:    unit.Dp(r.top),
			Bottom: unit.Dp(r.bottom),
			Left:   unit.Dp(r.left),
			Right:  unit.Dp(r.right),
		}
		return margins.Layout(graphicalContext,
			DrawButton(theme, button, text, textSize, c),
		)
	}
}
func DrawButton(theme *material.Theme, button *widget.Clickable, text string, textSize float32, c Color) func(graphicalContext layout.Context) layout.Dimensions {

	return func(graphicalContext layout.Context) layout.Dimensions {
		button := material.Button(theme, button, text)
		button.Background = color.NRGBA{R: c.A, G: c.G, B: c.B, A: c.A}
		button.TextSize = unit.Dp(textSize)
		return button.Layout(graphicalContext) // returns dimensions
	}
}

func DrawTitle(theme *material.Theme, operations *op.Ops, rectangleHeight float32, titleText string, textSize float32, c Color, r Rect) func(graphicalContext layout.Context) layout.Dimensions {

	return func(graphics_context layout.Context) layout.Dimensions {

		graphics_context.Constraints.Min.Y = graphics_context.Px(unit.Dp(rectangleHeight))

		top_rectangle := image.Rectangle{Max: graphics_context.Constraints.Min}
		paint.FillShape(operations, color.NRGBA{R: c.R, G: c.G, B: c.B, A: c.A}, clip.Rect(top_rectangle).Op())

		title := material.Label(theme, unit.Dp(textSize), titleText)
		title.Alignment = text.Middle
		title.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
		return layout.Inset{Right: unit.Dp(r.right), Left: unit.Dp(r.left), Top: unit.Dp(r.top), Bottom: unit.Dp(r.bottom)}.Layout(graphics_context, title.Layout)
	}
}
