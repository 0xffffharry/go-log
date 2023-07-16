package log

import "github.com/fatih/color"

type ColorLayer interface {
	ColorEnabled() bool
	EnableColor()
	DisableColor()
	SetColor(c color.Attribute)
	GetColor() color.Attribute
	Print(str string) string
	PrintWithColor(color color.Attribute, str string) string
}

type ColorLogger interface {
	InitColor()
	ColorLayer() ColorLayer
}
