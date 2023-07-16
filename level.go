package log

import "fmt"

type Level struct {
	_level     int
	name       string
	_errOutput bool
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

func (l Level) Compare(l2 Level) bool {
	return l._level == l2._level && l.name == l2.name && l._errOutput == l2._errOutput
}

var levelMap syncMap[int, Level]

var (
	LevelDebug = Level{_level: 0, name: "Debug", _errOutput: false}
	LevelInfo  = Level{_level: 1, name: "Info", _errOutput: false}
	LevelWarn  = Level{_level: 2, name: "Warn", _errOutput: false}
	LevelError = Level{_level: 3, name: "Error", _errOutput: true}
	LevelFatal = Level{_level: 4, name: "Fatal", _errOutput: true}
	LevelPanic = Level{_level: 5, name: "Panic", _errOutput: true}
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
