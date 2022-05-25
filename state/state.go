package state

import (
	"gioui.org/font/gofont"     // Special gioui font.
	"gioui.org/widget/material" // Theme
)

var Theme *material.Theme = material.NewTheme(gofont.Collection())

var ShouldQuit bool = false
