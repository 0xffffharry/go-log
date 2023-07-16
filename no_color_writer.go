package log

import "io"

type NoColorWriter struct {
	io.Writer
}

func (w *NoColorWriter) Write(p []byte) (n int, err error) {
	newP := CleanColor(p)
	_, err = w.Writer.Write(newP)
	if err != nil {
		return 0, err
	}
	n = len(p)
	return
}
