package main

import (
	"strings"
	"bufio"
	"testing"
)

// Helper function to create a test instance of FileMetrics with the given flags setup for the instance.
func newTestFileMetrics(tb testing.TB, InPath string, LineFlag bool, WordFlag bool, CharFlag bool, ByteFlag bool) *FileMetrics {
	tb.Helper()
	fm := new(FileMetrics)
	fm.CompleteFilePath = strings.TrimSpace(InPath)
	fm.ByteFlag = ByteFlag
	fm.CharFlag = CharFlag
	fm.LineFlag = LineFlag
	fm.WordFlag = WordFlag
	return fm
}

// Test case to validate the working of the ProcessMetrics function of FileMetrics.
func Test_FileMetrics_Process(t *testing.T) {
	testCases := []struct {
		Name string
		InFilePath string
		InValue string
		InLineFlag bool
		InWordFlag bool
		InCharFlag bool
		InByteFlag bool
		OutLineCount int
		OutWordCount int
		OutCharCount int
		OutByteCount int
		OutError string
	} {
		{ "A single line input with all flags set and ending with newline", "/Users/maheshkumaarbalaji/Projects/mkbwc/TestFiles/file-one.txt", "I am Mahesh Kumaar Balaji.\n", true, true, true, true, 1, 5, 26, 26, "" },
		{ "A single line input with all flags set but not ending with newline", "/Users/maheshkumaarbalaji/Projects/mkbwc/TestFiles/file-one.txt", "I am a Software Engineer.", true, true, true, true, 1, 5, 25, 25, "" },
		{ "A multi-line input with line flag set", "/Users/maheshkumaarbalaji/Projects/mkbwc/TestFiles/file-one.txt", "I am Mahesh Kumaar Balaji.\nI am a Software Engineer.\n", true, false, false, false, 2, 0, 0, 0, "" },
		{ "A multi-line input with word flag set", "/Users/maheshkumaarbalaji/Projects/mkbwc/TestFiles/file-one.txt", "I am Mahesh Kumaar Balaji.\nI am a Software Engineer.\n", false, true, false, false, 0, 10, 0, 0, "" },
		{ "A multi-line input with character flag set", "/Users/maheshkumaarbalaji/Projects/mkbwc/TestFiles/file-one.txt", "I am Mahesh Kumaar Balaji.\nI am a Software Engineer.\n", false, false, true, false, 0, 0, 51, 0, "" },
		{ "A multi-line input with byte flag set", "/Users/maheshkumaarbalaji/Projects/mkbwc/TestFiles/file-one.txt", "I am Mahesh Kumaar Balaji.\nI am a Software Engineer.\n", false, false, false, true, 0, 0, 0, 51, "" },
		{ "A single line input with non-ASCII characters", "/Users/maheshkumaarbalaji/Projects/mkbwc/TestFiles/file-one.txt", "This is so @்லஉஎஉ்க்எஙுஇக\n", true, true, true, true, 1, 4, 25, 51, ""  },
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(tt *testing.T) {
			fm := newTestFileMetrics(tt, testCase.InFilePath, testCase.InLineFlag, testCase.InWordFlag, testCase.InCharFlag, testCase.InByteFlag)
			strReader := strings.NewReader(testCase.InValue)
			reader := bufio.NewReader(strReader)
			errOne := fm.ProcessMetrics(reader)
			if strings.EqualFold(testCase.OutError, "") {
				if errOne != nil {
					tt.Errorf("Was not expecting an error, and yet got this instead - %v", errOne)
					return
				}
			}

			if strings.EqualFold(testCase.OutError, "FileAccessError") {
				fae, ok := errOne.(*FileAccessError)
				if ok {
					tt.Logf("Was expecting a file access error and received one - %v", fae)
				} else {
					tt.Errorf("Was expecting a file access error, but got this instead - %v", errOne)
				}
				return
			}

			if testCase.OutLineCount == fm.LineCount {
				tt.Logf("The expected line count [%d] matches the processed line count [%d]", testCase.OutLineCount, fm.LineCount)
			} else {
				tt.Errorf("The expected line count [%d] does not match the processed line count [%d]", testCase.OutLineCount, fm.LineCount)
			}

			if testCase.OutWordCount == fm.WordCount {
				tt.Logf("The expected word count [%d] matches the processed word count [%d]", testCase.OutWordCount, fm.WordCount)
			} else {
				tt.Errorf("The expected word count [%d] does not match the processed word count [%d]", testCase.OutWordCount, fm.WordCount)
			}

			if testCase.OutCharCount == fm.CharCount {
				tt.Logf("The expected character count [%d] matches the processed character count [%d]", testCase.OutCharCount, fm.CharCount)
			} else {
				tt.Errorf("The expected character count [%d] does not match the processed character count [%d]", testCase.OutCharCount, fm.CharCount)
			}

			if testCase.OutByteCount == fm.ByteCount {
				tt.Logf("The expected byte count [%d] matches the processed byte count [%d]", testCase.OutByteCount, fm.ByteCount)
			} else {
				tt.Errorf("The expected byte count [%d] does not match the processed byte count [%d]", testCase.OutByteCount, fm.ByteCount)
			}
		})
	}
}

// Test case to validate the working of the Print function of FileMetrics.
func Test_FileMetrics_Print(t *testing.T) {
	testCases := []struct {
		Name string
		InFilePath string
		InValue string
		InLineFlag bool
		InWordFlag bool
		InCharFlag bool
		InByteFlag bool
		OutLine string
		OutError string
	} {
		{ "A single line input with all flags set and ending with newline", "/Users/maheshkumaarbalaji/Projects/mkbwc/TestFiles/file-one.txt", "I am Mahesh Kumaar Balaji.\n", true, true, true, true, "1  5  26  26  file-one.txt", "" },
		{ "A single line input with all flags set but not ending with newline", "/Users/maheshkumaarbalaji/Projects/mkbwc/TestFiles/file-one.txt", "I am a Software Engineer.", true, true, true, true, "1  5  25  25  file-one.txt", "" },
		{ "A multi-line input with line flag set", "/Users/maheshkumaarbalaji/Projects/mkbwc/TestFiles/file-one.txt", "I am Mahesh Kumaar Balaji.\nI am a Software Engineer.\n", true, false, false, false, "2  file-one.txt", "" },
		{ "A multi-line input with word flag set", "/Users/maheshkumaarbalaji/Projects/mkbwc/TestFiles/file-one.txt", "I am Mahesh Kumaar Balaji.\nI am a Software Engineer.\n", false, true, false, false, "10  file-one.txt", "" },
		{ "A multi-line input with character flag set", "/Users/maheshkumaarbalaji/Projects/mkbwc/TestFiles/file-one.txt", "I am Mahesh Kumaar Balaji.\nI am a Software Engineer.\n", false, false, true, false, "51  file-one.txt", "" },
		{ "A multi-line input with byte flag set", "/Users/maheshkumaarbalaji/Projects/mkbwc/TestFiles/file-one.txt", "I am Mahesh Kumaar Balaji.\nI am a Software Engineer.\n", false, false, false, true, "51  file-one.txt", "" },
		{ "A single line input with non-ASCII characters","/Users/maheshkumaarbalaji/Projects/mkbwc/TestFiles/file-one.txt", "This is so @்லஉஎஉ்க்எஙுஇக\n", true, true, true, true, "1  4  25  51  file-one.txt", "" },
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(tt *testing.T) {
			fm := newTestFileMetrics(tt, testCase.InFilePath, testCase.InLineFlag, testCase.InWordFlag, testCase.InCharFlag, testCase.InByteFlag)
			strReader := strings.NewReader(testCase.InValue)
			reader := bufio.NewReader(strReader)
			errOne := fm.ProcessMetrics(reader)
			if strings.EqualFold(testCase.OutError, "") {
				if errOne != nil {
					tt.Errorf("Was not expecting an error, and yet got this instead - %v", errOne)
					return
				}
			}

			if strings.EqualFold(testCase.OutError, "FileAccessError") {
				fae, ok := errOne.(*FileAccessError)
				if ok {
					tt.Logf("Was expecting a file access error and received one - %v", fae)
				} else {
					tt.Errorf("Was expecting a file access error, but got this instead - %v", errOne)
				}
				return
			}

			outValue := fm.Print()
			if strings.EqualFold(outValue, testCase.OutLine) {
				tt.Logf("The line returned by print [%s] matches the expected value [%s].", outValue, testCase.OutLine)
			} else {
				tt.Errorf("The line returned by print [%s] does not match the expected value [%s].", outValue, testCase.OutLine)
			}
		})
	}
}