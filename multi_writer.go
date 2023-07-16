package log

import (
	"io"
)

type MultiWriter struct {
	m syncMap[io.Writer, any]
}

func NewMultiWriter() *MultiWriter {
	m := &MultiWriter{}
	return m
}

func (w *MultiWriter) RegisterWriter(writer io.Writer) {
	w.m.Store(writer, nil)
}

func (w *MultiWriter) UnregisterWriter(writer io.Writer) {
	w.m.Delete(writer)
}

func (w *MultiWriter) Len() int {
	return w.m.Len()
}

func (w *MultiWriter) Write(p []byte) (n int, err error) {
	w.m.GoRange(func(writer io.Writer, _ any) {
		writer.Write(p)
	})
	return len(p), nil
}
