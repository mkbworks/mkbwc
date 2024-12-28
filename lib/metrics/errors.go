package metrics

import (
	"fmt"
)

// A custom error to keep track of events raised when trying to access a file or its properties.
type FileAccessError struct {
	// The complete path of the file being accessed.
	CompleteFilePath string
	// Action taken that led to the error being raised.
	Action string
	// The actual message associated with the error that occurred.
	Message string
}

// Error message associated with the instance of FileAccessError.
func (fae *FileAccessError) Error() string {
	return fmt.Sprintf("File Access Error :: File - %s :: Action - %s :: %s", fae.CompleteFilePath, fae.Action, fae.Message)
}