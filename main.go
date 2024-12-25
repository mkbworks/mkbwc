package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: ./mkbwc.out [options] filename(s)")
		fmt.Println("Options available:")
		flag.PrintDefaults()
	}

	BytesFlag := flag.Bool("c", false, "Output the number of bytes in the given file")
	LinesFlag := flag.Bool("l", false, "Output the number of lines in the given file")
	WordsFlag := flag.Bool("w", false, "Output the number of words in the given file")
	CharsFlag := flag.Bool("m", false, "Output the number of characters in the given file")
	HelpFlag := flag.Bool("h", false, "Show the help message")
	flag.Parse()

	if *HelpFlag {
		flag.Usage()
		os.Exit(0)
	}

	FileNames := flag.Args()
	if len(FileNames) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			file := scanner.Text()
			FileNames = append(FileNames, file)
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error occurred while reading piped stream input: %s\n", err.Error())
			os.Exit(1)
		}
	}

	bp := new(BulkProcess)
	bp.LineFlag = *LinesFlag
	bp.WordFlag = *WordsFlag
	bp.CharFlag = *CharsFlag
	bp.ByteFlag = *BytesFlag
	bp.Files = make([]string, 0)
	bp.Files = append(bp.Files, FileNames...)
	bp.ComputeMetrics()
}