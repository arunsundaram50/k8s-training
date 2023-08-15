package main

import (
	"io"
)

type TeeWriter struct {
	writers []io.Writer
}

func NewTeeWriter(writers ...io.Writer) *TeeWriter {
	return &TeeWriter{writers: writers}
}

func (t *TeeWriter) Write(p []byte) (n int, err error) {
	for _, writer := range t.writers {
		n, writeErr := writer.Write(p)
		if writeErr != nil && err == nil {
			err = writeErr
		}
		if n != len(p) && err == nil {
			return n, io.ErrShortWrite
		}
	}
	return len(p), err
}
