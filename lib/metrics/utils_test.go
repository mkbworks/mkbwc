package metrics

import (
	"testing"
	"strings"
)

// Test case to validate the working of creating a new FileMetrics instance and setting up the appropriate flags.
func Test_NewFileMetrics(t *testing.T) {
	testCases := []struct {
		Name string
		InFilePath string
		InLineFlag bool
		InWordFlag bool
		InCharFlag bool
		InByteFlag bool
		OutLineFlag bool
		OutWordFlag bool
		OutCharFlag bool
		OutByteFlag bool
		OutError string
	} {
		{ "File metrics with line flag only", "/Users/maheshkumaarbalaji/Projects/mkbwc/TestFiles/file-one.txt", true, false, false, false, true, false, false, false, "" },
		{ "File metrics with word flag only", "/Users/maheshkumaarbalaji/Projects/mkbwc/TestFiles/file-one.txt", false, true, false, false, false, true, false, false, "" },
		{ "File metrics with character flag only", "/Users/maheshkumaarbalaji/Projects/mkbwc/TestFiles/file-one.txt", false, false, true, false, false, false, true, false, "" },
		{ "File metrics with byte flag only", "/Users/maheshkumaarbalaji/Projects/mkbwc/TestFiles/file-one.txt", false, false, false, true, false, false, false, true, "" },
		{ "File metrics with default option",  "/Users/maheshkumaarbalaji/Projects/mkbwc/TestFiles/file-one.txt", false, false, false, false, true, true, false, true, "" },
		{ "Invalid file path",  "/Users/maheshkumaarbalaji/Projects/mkbwc/TestFiles/file-two.txt", false, false, false, false, false, true, false, false, "FileAccessError" },
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(tt *testing.T) {
			fm, err := newFileMetrics(testCase.InFilePath, testCase.InLineFlag, testCase.InWordFlag, testCase.InCharFlag, testCase.InByteFlag)
			if strings.EqualFold(testCase.OutError, "") {
				if err != nil {
					tt.Errorf("Was not expecting an error and yet got this - %v", err)
					return
				}
			}

			if strings.EqualFold(testCase.OutError, "FileAccessError") {
				fae, ok := err.(*FileAccessError)
				if !ok {
					tt.Errorf("Was expecting a FileAccessError, but got this instead - %v", err)
				} else {
					tt.Logf("Was expecting a FileAccessError and got the same error - %v", fae)
				}
				return
			}

			if fm.LineFlag == testCase.OutLineFlag {
				tt.Log("The expected line flag matches the line flag setup for FileMetrics instance.")
			} else {
				tt.Error("The expected line flag does not match the line flag setup for FileMetrics instance.")
			}

			if fm.WordFlag == testCase.OutWordFlag {
				tt.Log("The expected word flag matches the word flag setup for FileMetrics instance.")
			} else {
				tt.Error("The expected word flag does not match the word flag setup for FileMetrics instance.")
			}

			if fm.CharFlag == testCase.OutCharFlag {
				tt.Log("The expected character flag matches the character flag setup for FileMetrics instance.")
			} else {
				tt.Error("The expected character flag does not match the character flag setup for FileMetrics instance.")
			}

			if fm.ByteFlag == testCase.OutByteFlag {
				tt.Log("The expected byte flag matches the byte flag setup for FileMetrics instance.")
			} else {
				tt.Error("The expected byte flag does not match the byte flag setup for FileMetrics instance.")
			}
		})
	}
}