package log

import "github.com/fatih/color"

type colorLayer struct {
	status    bool
	attribute color.Attribute
	color     *color.Color
}

func newColorLayer() *colorLayer {
	return &colorLayer{
		status:    false,
		attribute: color.FgBlack,
		color:     NewColor(color.FgBlack),
	}
}

func (l *colorLayer) ColorEnabled() bool {
	return l.status
}

func (l *colorLayer) EnableColor() {
	l.status = true
}

func (l *colorLayer) DisableColor() {
	l.status = false
}

func (l *colorLayer) SetColor(c color.Attribute) {
	l.attribute = c
	l.color = NewColor(c)
}

func (l *colorLayer) GetColor() color.Attribute {
	return l.attribute
}

func (l *colorLayer) Print(str string) string {
	if l.status {
		if l.color != nil {
			return l.color.Sprint(str)
		} else {
			return str
		}
	}
	return str
}

func (l *colorLayer) PrintWithColor(color color.Attribute, str string) string {
	if l.status {
		if color > 0 {
			return NewColor(color).Sprint(str)
		}
		if l.color != nil {
			return l.color.Sprint(str)
		} else {
			return str
		}
	}
	return str
}
