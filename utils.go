package main

import (
	"strings"
)

// Creates and returns a pointer to a new instance of FileMetrics after setting all the appropriate metric flags.
func newFileMetrics(CompleteFilePath string, LineFlag bool, WordFlag bool, CharFlag bool, ByteFlag bool) *FileMetrics {
	fm := new(FileMetrics)
	CompleteFilePath = strings.TrimSpace(CompleteFilePath)
	fm.CompleteFilePath = CompleteFilePath
	if !LineFlag && !WordFlag && !CharFlag && !ByteFlag {
		// default option where none of the command-line flags are given
		fm.LineFlag = true
		fm.WordFlag = true
		fm.ByteFlag = true
	} else {
		if LineFlag {
			fm.LineFlag = true
		}

		if WordFlag {
			fm.WordFlag = true
		}

		if CharFlag {
			fm.CharFlag = true
		}

		if ByteFlag {
			fm.ByteFlag = true
		}
	}

	return fm
}