package metrics

import (
	"os"
	"strings"
	"path/filepath"
)

// Creates and returns a pointer to a new instance of FileMetrics after setting all the appropriate metric flags.
func newFileMetrics(CompleteFilePath string, LineFlag bool, WordFlag bool, CharFlag bool, ByteFlag bool) (*FileMetrics, error) {
	CompleteFilePath = strings.TrimSpace(CompleteFilePath)
	fileStat, err := os.Stat(CompleteFilePath)
	if err != nil {
		fae := new(FileAccessError)
		fae.CompleteFilePath = CompleteFilePath
		fae.Action = "File Access"
		fae.Message = err.Error()
		return nil, fae
	}

	fileMode := fileStat.Mode()
	if !fileMode.IsRegular() {
		fae := new(FileAccessError)
		fae.CompleteFilePath = CompleteFilePath
		fae.Action = "File Mode"
		fae.Message = "The given path does not point to a file"
		return nil, fae
	}

	fileExtension := filepath.Ext(CompleteFilePath)
	fileExtension = strings.TrimSpace(fileExtension)
	fileExtension = strings.TrimPrefix(fileExtension, ".")
	if !strings.EqualFold(fileExtension, "txt") {
		fae := new(FileAccessError)
		fae.CompleteFilePath = CompleteFilePath
		fae.Action = "File Extension"
		fae.Message = "The given file path does not point to a text file."
		return nil, fae
	}

	fm := new(FileMetrics)
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

	return fm, nil
}