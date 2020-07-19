package ioline

import (
	"bufio"
	"io"
	"strings"
)

// Scanner provides methods to read the entire file line-by-line, omitting newline character
type Scanner interface {
	Next() bool
	Get() (string, error)
	Error() error
	HasError() bool
}

type linesIterator struct {
	rd         io.Reader
	reader     *bufio.Reader
	err        error
	line       string
	isFinished bool
}

// Next moves the iteration to the next line
func (iterator *linesIterator) Next() bool {
	if iterator.ended() {
		return false
	}

	iterator.line = ""
	line, err := iterator.ReadEntireLine()

	switch err {
	case io.EOF:
		iterator.isFinished = true
	case nil:
		iterator.line = line
	default:
		iterator.err = err
	}

	return iterator.needsIteration()
}

// Error gets the last error
func (iterator *linesIterator) Error() error {
	return iterator.err
}

// HasError is true when an error is triggered
func (iterator *linesIterator) HasError() bool {
	return iterator.err != nil
}

// Get the current line the iterator is positioned at
func (iterator *linesIterator) Get() (string, error) {
	return iterator.line, iterator.err
}

// NewScanner starts a new `Scanner` for the provided `io.Reader`
func NewScanner(rd io.Reader) Scanner {
	iterator := &linesIterator{}
	iterator.rd = rd
	iterator.reader = bufio.NewReader(iterator.rd)

	return iterator
}

func (iterator *linesIterator) ReadEntireLine() (string, error) {
	reader := iterator.reader
	var entireLine strings.Builder

	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil && err != io.EOF {
			return "", err
		}

		_, bufErr := entireLine.Write(line)
		if bufErr != nil {
			return "", bufErr
		}

		if isPrefix == false && (err == nil || err == io.EOF) {
			return entireLine.String(), err
		}
	}
}

func (iterator *linesIterator) ended() bool {
	return iterator.isFinished || iterator.err != nil
}

func (iterator *linesIterator) needsIteration() bool {
	return !iterator.ended()
}
