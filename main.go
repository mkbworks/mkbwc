package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	BytesFlag := flag.Bool("-c", false, "Output the number of bytes in the given file")
	LinesFlag := flag.Bool("-l", false, "Output the number of lines in the given file")
	WordsFlag := flag.Bool("-w", false, "Output the number of words in the given file")
	CharsFlag := flag.Bool("-m", false, "Output the number of characters in the given file")
	flag.Parse()
	FileNames := flag.Args()
	if len(FileNames) == 0 {
		fmt.Println("No filenames were provided in the command-line arguments.")
		os.Exit(1)
	}

	for _, FileName := range FileNames {
		fm := newFileMetrics(FileName, *LinesFlag, *WordsFlag, *CharsFlag, *BytesFlag)
		err := fm.Compute()
		if err != nil {
			fmt.Printf("Error while computing metrics for file [%s] :: %s", FileName, err.Error())
			continue
		}

		fmt.Println(fm.Print())
	}
}