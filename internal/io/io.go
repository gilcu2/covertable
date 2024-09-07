package io

import (
	"fmt"
	"github.com/gilcu2/covertable/internal/coverage"
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

func PrintTable(coverages []coverage.CoverTable, writer io.Writer) error {
	var totalLines = 0
	var totalCovered = 0
	fmt.Fprintf(writer, "File\tCoverage\tUncovered lines\n")
	for _, fileCover := range coverages {
		totalLines += fileCover.TotalLines
		totalCovered += fileCover.CoveredLines
		var fileCoverage = float32(fileCover.CoveredLines) / float32(fileCover.TotalLines)
		var builder strings.Builder
		for _, block := range fileCover.UncoveredBlocks {
			builder.WriteString(fmt.Sprintf("%d-%d,", block.Begin, block.End))
		}
		fmt.Fprintf(writer, "%s\t%.2g\t%s\n", fileCover.Filename, fileCoverage, builder.String())
	}
	return nil
}
