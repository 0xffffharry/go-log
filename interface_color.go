package log

import "github.com/fatih/color"

type ColorLayer interface {
	ColorEnabled() bool
	EnableColor()
	DisableColor()
	SetColor(c color.Attribute)
	GetColor() color.Attribute
	Print(message string) string
}

type ColorLogger interface {
	InitColor()
	ColorLayer() ColorLayer
}
