package io

import (
	"covertable/internal/coverage"
	"fmt"
	"io"
	"os"
	"strings"
)

func MakeTableFromFile(fileName string) ([]coverage.CoverTable, error) {
	var reader, readErr = os.Open(fileName)
	if readErr != nil {
		return nil, readErr
	}
	defer reader.Close()

	var coverages, parseErr = coverage.MakeTableFromReader(reader)
	if parseErr != nil {
		return nil, parseErr
	}

	return coverages, nil
}

func PrintTable(coverages []coverage.CoverTable, writter io.Writer) float32 {
	var totalLines = 0
	var totalCovered = 0
	fmt.Fprint(writter, "File\tCoverage\tUncovered lines\n")
	for _, fileCover := range coverages {
		totalLines += fileCover.TotalLines
		totalCovered += fileCover.CoveredLines
		var fileCoverage = fileCover.CoveredLines / fileCover.TotalLines
		var builder strings.Builder
		for _, block := range fileCover.UncoveredBlocks {
			builder.WriteString(fmt.Sprintf("%d-%d,", block.Begin, block.End))
		}
		fmt.Fprint(writter, "%s\t%f\t%s\n", fileCover.Filename, fileCoverage, builder.String())
	}

}
