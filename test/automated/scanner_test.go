package ioline_test

import (
	"testing"

	. "github.com/thehungry-dev/asserting"
	"github.com/thehungry-dev/ioline"
	"github.com/thehungry-dev/ioline/controls"
)

func TestScanner(t *testing.T) {
	t.Run("Scanner", func(t *testing.T) {
		t.Run("Reader doesn't encounter errors", func(t *testing.T) {
			t.Run("Last line has line ending", func(t *testing.T) {
				reader := controls.ThreeLinesReader()
				scanner := ioline.NewScanner(reader)
				var scannerErr error
				var lines []string

				for scanner.Next() {
					line, err := scanner.Get()
					lines = append(lines, line)
					if err != nil {
						scannerErr = err
					}
				}

				t.Run("Successful", func(t *testing.T) {
					Assert(t, scannerErr == nil)
				})

				linesCount := len(lines)

				t.Run("Gets all the lines in a stream", func(t *testing.T) {
					Assert(t, lines[0] == controls.Line1)
					Assert(t, lines[1] == controls.Line2)
					Assert(t, lines[2] == controls.Line3)
					Assert(t, linesCount == 3)
				})
			})

			t.Run("Last line doesn't include line-ending", func(t *testing.T) {
				reader := controls.TwoLinesAndEOFReader()
				scanner := ioline.NewScanner(reader)
				var scannerErr error
				var lines []string

				for scanner.Next() {
					line, err := scanner.Get()
					lines = append(lines, line)
					if err != nil {
						scannerErr = err
					}
				}

				t.Run("Successful", func(t *testing.T) {
					Assert(t, scannerErr == nil)
				})

				linesCount := len(lines)

				t.Run("Gets all the lines in a stream", func(t *testing.T) {
					Assert(t, lines[0] == controls.Line1)
					Assert(t, lines[1] == controls.Line2)
					Assert(t, linesCount == 2)
				})
			})
		})

		t.Run("Reader encounters error", func(t *testing.T) {
			reader := controls.ErrorReader()
			scanner := ioline.NewScanner(reader)
			var iterations uint

			for scanner.Next() {
				iterations++
			}

			t.Run("Doesn't iterate", func(t *testing.T) {
				Assert(t, iterations == 0)
			})

			t.Run("Errors", func(t *testing.T) {
				Assert(t, scanner.HasError())
				Assert(t, scanner.Error() == controls.ErrErrorReader)
			})
		})
	})
}
