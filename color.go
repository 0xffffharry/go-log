package log

import (
	"bytes"

	"github.com/fatih/color"
)

var colorMap syncMap[color.Attribute, *color.Color]

func NewColor(c color.Attribute) *color.Color {
	var cl *color.Color
	cl, exist := colorMap.Load(c)
	if !exist || cl == nil {
		cl = color.New(c)
		colorMap.Store(c, cl)
	}
	return cl
}

var bytesBuffer syncPool[*bytes.Buffer]

func init() {
	bytesBuffer.New(func() *bytes.Buffer {
		return bytes.NewBuffer(nil)
	})
}

func CleanColor(b []byte) []byte {
	buffer := bytesBuffer.Get()
	defer bytesBuffer.Put(buffer)
	defer buffer.Reset()
	var head bool
	for _, c := range b {
		if c == '\x1b' && head {
			buffer.WriteByte(c)
			continue
		}
		if c == '\x1b' && !head {
			head = true
			continue
		}
		if head {
			if c == '[' {
				continue
			}
			if c >= '0' && c <= '9' {
				continue
			}
			if c == ';' {
				continue
			}
			if c == 'm' {
				head = false
				continue
			}
			head = false
		}
		buffer.WriteByte(c)
	}
	return buffer.Bytes()
}
