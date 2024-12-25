package main

import (
	"sync"
	"fmt"
	"os"
)

// Structure to bulk process files and compute their metrics.
type BulkProcess struct {
	// Instance of waitgroup to synchronize the termination of all goroutines.
	wg sync.WaitGroup
	// List of files to be processed.
	Files []string
	// Flag to determine if line count has to be computed.
	LineFlag bool
	// Flag to determine if word count has to be computed.
	WordFlag bool
	// Flag to determine if character count has to be computed.
	CharFlag bool
	// Flag to determine if byte count has to be computed.
	ByteFlag bool
}

// Compute metrics for files in the slice.
func (bp *BulkProcess) ComputeMetrics() {
	for _, file := range bp.Files {
		go bp.ProcessFile(file)
		bp.wg.Add(1)
	}

	bp.wg.Wait()
}

// Compute and display the metrics for an individual file.
func (bp *BulkProcess) ProcessFile(file string) {
	defer bp.wg.Done()
	fm, err := newFileMetrics(file, bp.LineFlag, bp.WordFlag, bp.CharFlag, bp.ByteFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred while computing file metrics for file [%s] :: %s", file, err.Error())
		return
	}

	err = fm.Compute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while computing metrics for file [%s] :: %s", file, err.Error())
		return
	}

	fmt.Println(fm.Print())
}