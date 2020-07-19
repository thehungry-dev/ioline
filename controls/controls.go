// Package controls provides test controls for the asserting package
package controls

import (
	"fmt"
	"io"
	"strings"
)

// Line representation
const (
	Line1 string = "Line 1"
	Line2 string = "Line 2"
	Line3 string = "Line 3"
)

// ThreeLines represents 3 lines ending with new line characters
var ThreeLines string = fmt.Sprintf("%s\n%s\n%s\n", Line1, Line2, Line3)

// TwoLinesAndEOF represents 2 lines where the last one doesn't include line-ending character
var TwoLinesAndEOF string = fmt.Sprintf("%s\n%s", Line1, Line2)

// ThreeLinesReader is an `io.Reader` for `ThreeLines`
func ThreeLinesReader() io.Reader {
	return strings.NewReader(ThreeLines)
}

// TwoLinesAndEOFReader is an `io.Reader` for `TwoLinesAndEOF`
func TwoLinesAndEOFReader() io.Reader {
	return strings.NewReader(TwoLinesAndEOF)
}
