package controls

import (
	"errors"
	"io"
)

type errorReader struct{}

// Read 0 bytes, always returning error `ErrErrorReader`
func (errorReader) Read([]byte) (int, error) {
	return 0, ErrErrorReader
}

// ErrErrorReader represents the error returned by an `ErrorReader`
var ErrErrorReader error = errors.New("Error reader")

// ErrorReader is an `io.Reader` that returns an error every time an attempt to read is performed
func ErrorReader() io.Reader {
	return errorReader{}
}
