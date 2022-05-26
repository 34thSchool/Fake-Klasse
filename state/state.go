package state

import (
	"gioui.org/widget/material" // Theme
)

type State struct{
	Theme *material.Theme
	ShouldQuit bool
}