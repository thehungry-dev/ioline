package ioline

import (
	"os"
)

// Line represents a single file line with buffers for the lines before and after it
type Line struct {
	FirstLine int
	Before    []string
	Exact     string
	After     []string
}

// Count the total number of lines
func (line Line) Count() int {
	return 1 + len(line.Before) + len(line.After)
}

// ReadFile gets the specified line from the file at the provided path
func ReadFile(path string, currentLine int) (string, error) {
	line, err := ReadFileWithBuffers(path, currentLine, 0, 0)
	return line.Exact, err
}

// ReadFileWithBuffers returns the specified line with a buffer of lines before and after it
func ReadFileWithBuffers(path string, currentLine int, beforeLinesAmount int, afterLinesAmount int) (Line, error) {
	lastLine := currentLine + afterLinesAmount
	firstLine := currentLine - beforeLinesAmount
	if firstLine < 0 {
		firstLine = 1
	}

	lines := Line{FirstLine: firstLine}
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return lines, err
	}

	scanner := NewScanner(file)
	lineIndex := 1

	for scanner.Next() {
		line, err := scanner.Get()

		if err != nil {
			return lines, err
		}

		if lineIndex >= firstLine && lineIndex < currentLine {
			lines.Before = append(lines.Before, line)
		}
		if lineIndex == currentLine {
			lines.Exact = line
		}
		if lineIndex > currentLine && lineIndex <= lastLine {
			lines.After = append(lines.After, line)
		}

		lineIndex++
	}

	return lines, nil
}
