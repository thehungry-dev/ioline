package ioline_test

import (
	"testing"

	. "github.com/thehungry-dev/asserting"
	"github.com/thehungry-dev/ioline"
	"github.com/thehungry-dev/ioline/controls"
)

func TestReadFileWithBuffers(t *testing.T) {
	t.Run("Read file with buffers", func(t *testing.T) {
		t.Run("File is missing", func(t *testing.T) {
			_, err := ioline.ReadFileWithBuffers(controls.MissingFilePath, 0, 0, 0)

			t.Run("Errors", func(t *testing.T) {
				Assert(t, err != nil)
			})
		})

		t.Run("File is present", func(t *testing.T) {
			file := controls.CreateTempFile()
			defer controls.RemoveTempFile(file)

			lines, err := ioline.ReadFileWithBuffers(file.Name(), 2, 1, 1)

			t.Run("Successful", func(t *testing.T) {
				Assert(t, err == nil)
			})

			t.Run("Reads 3 lines", func(t *testing.T) {
				Assert(t, lines.Count() == 3)
			})

			t.Run("Reads from first line in file", func(t *testing.T) {
				Assert(t, lines.FirstLine == 1)
			})

			t.Run("Reads one line before the target line", func(t *testing.T) {
				Assert(t, lines.Before[0] == controls.Line1)
			})

			t.Run("Reads the expected requested line", func(t *testing.T) {
				Assert(t, lines.Exact == controls.Line2)
			})

			t.Run("Reads one line after the target line", func(t *testing.T) {
				Assert(t, lines.After[0] == controls.Line3)
			})
		})
	})
}
