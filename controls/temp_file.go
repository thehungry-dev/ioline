package controls

import (
	"fmt"
	"io/ioutil"
	"os"
)

// MissingFilePath represents a non-existing file path
const MissingFilePath string = "/_invalid_/path"

// TempFileContent represents the content of the temporary file
var TempFileContent string = fmt.Sprintf("%s\n%s\n%s", Line1, Line2, Line3)

// CreateTempFile writes a new temporary file
func CreateTempFile() *os.File {
	file, err := ioutil.TempFile("", "")
	if err != nil {
		panic(err)
	}
	_, err = file.WriteString(TempFileContent)
	if err != nil {
		panic(err)
	}

	return file
}

// RemoveTempFile deletes the temporary file
func RemoveTempFile(file *os.File) {
	os.Remove(file.Name())
}
