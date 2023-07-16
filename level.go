package log

import (
	"fmt"

	"github.com/fatih/color"
)

type Level struct {
	_level     int
	name       string
	_errOutput bool
	_color     color.Attribute
}

func (l Level) String() string {
	return l.name
}

func (l Level) level() int {
	return l._level
}

func (l Level) errOutput() bool {
	return l._errOutput
}

func (l Level) color() color.Attribute {
	return l._color
}

func (l Level) Compare(l2 Level) bool {
	return l._level == l2._level && l.name == l2.name && l._errOutput == l2._errOutput && l._color == l2._color
}

var levelMap syncMap[int, Level]

var (
	LevelDebug = Level{_level: 0, name: "Debug", _errOutput: false, _color: color.FgBlue}
	LevelInfo  = Level{_level: 1, name: "Info", _errOutput: false, _color: color.FgGreen}
	LevelWarn  = Level{_level: 2, name: "Warn", _errOutput: false, _color: color.FgYellow}
	LevelError = Level{_level: 3, name: "Error", _errOutput: true, _color: color.FgRed}
	LevelFatal = Level{_level: 4, name: "Fatal", _errOutput: true, _color: color.FgHiRed}
	LevelPanic = Level{_level: 5, name: "Panic", _errOutput: true, _color: color.FgCyan}
)

func init() {
	levelMap.Store(0, LevelDebug)
	levelMap.Store(1, LevelInfo)
	levelMap.Store(2, LevelWarn)
	levelMap.Store(3, LevelError)
	levelMap.Store(4, LevelFatal)
	levelMap.Store(5, LevelPanic)
}

func RegisterLevel(_level int, name string, _errOutput bool) (Level, error) {
	l := Level{_level: _level, name: name, _errOutput: _errOutput}
	if oldLevel, ok := levelMap.LoadOrStore(l._level, l); ok {
		return oldLevel, fmt.Errorf("level already exists")
	}
	return l, nil
}
