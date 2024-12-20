package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
	"path/filepath"
)

// Size of each chunk of data read from the file.
const CHUNK_SIZE = 1024

// Structure to store computed metrics for an individual file in the file system.
type FileMetrics struct {
	// The complete path of the file in the file system.
	CompleteFilePath string
	// Total number of lines present in the file.
	LineCount int
	// Total number of words present in the file.
	WordCount int
	// Total number of bytes present in the file.
	ByteCount int
	// Total number of unicode characters present in the file.
	CharCount int
	// Flag to denote if the total line count needs to be computed for the file.
	LineFlag bool
	// Flag to denote if the total word count needs to be computed for the file.
	WordFlag bool
	// Flag to denote if the total byte count needs to be computed for the file.
	ByteFlag bool
	// Flag to denote if the total character count needs to be computed for the file.
	CharFlag bool
}

// Compute the required metrics for the given file in the file system. 
// The method computes metrics as per the flags set for the file.
func (fm *FileMetrics) Compute() error {
	CompleteFilePath := strings.TrimSpace(fm.CompleteFilePath)
	if strings.EqualFold(CompleteFilePath, "") {
		return errors.New("file path cannot be empty")
	}

	fileHandler, err := os.Open(CompleteFilePath)
	if err != nil {
		return err
	}

	defer fileHandler.Close()
	reader := bufio.NewReader(fileHandler)
	for {
		nextLine, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}

		nextLine = strings.TrimRight(nextLine, "\n")
		nextLine = strings.TrimSpace(nextLine)

		if fm.LineFlag {
			fm.LineCount++
		}

		if fm.WordFlag {
			words := strings.Split(nextLine, " ")
			for _, word := range words {
				word = strings.TrimSpace(word)
				if !strings.EqualFold(word, "") {
					fm.WordCount++
				}
			}
		}

		if fm.CharFlag {
			fm.CharCount += utf8.RuneCountInString(nextLine)
		}

		if fm.ByteFlag {
			chunk := []byte(nextLine)
			fm.ByteCount += len(chunk)
		}
	}

	return nil
}

// Prints the computed metrics for the file as a line.
// Metrics are displayed in the following order - 
// [LineCount]  [WordCount]  [CharacterCount]  [ByteCount]  [FileName].
// If flag for a component is set as false, an empty string will be printed in its place.
func (fm *FileMetrics) Print() string {
	line := ""
	if fm.LineFlag {
		line += fmt.Sprintf("  %d", fm.LineCount)
	}

	if fm.WordFlag {
		line += fmt.Sprintf("  %d", fm.WordCount)
	}

	if fm.CharFlag {
		line += fmt.Sprintf("  %d", fm.CharCount)
	}

	if fm.ByteFlag {
		line += fmt.Sprintf("  %d", fm.ByteCount)
	}

	line += fmt.Sprintf("  %s", filepath.Base(fm.CompleteFilePath))
	line = strings.TrimSpace(line)
	return line
}