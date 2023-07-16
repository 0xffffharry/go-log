package log

import (
	"io"
)

type WriterChan struct {
	chanWriter chan []byte
	closed     bool
}

func NewWriterChan() *WriterChan {
	return NewWriterChanSize(0)
}

func NewWriterChanSize(size int) *WriterChan {
	if size <= 0 {
		size = 2048
	}
	return &WriterChan{
		chanWriter: make(chan []byte, size),
	}
}

func (w *WriterChan) Write(p []byte) (n int, err error) {
	if w.closed {
		return 0, io.ErrClosedPipe
	}
	for {
		select {
		case w.chanWriter <- p:
		default:
			if w.closed {
				return 0, io.ErrClosedPipe
			}
			continue
		}
		break
	}
	return len(p), nil
}

func (w *WriterChan) Chan() <-chan []byte {
	return w.chanWriter
}

func (w *WriterChan) Close() {
	w.closed = true
	close(w.chanWriter)
}
