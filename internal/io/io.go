package io

import (
	"fmt"
	"github.com/gilcu2/covertable/internal/golang"
	"io"
	"os"
	"strings"
)

func MakeTableFromFile(coverageFileName string, moduleFileName string) ([]golang.CoverTable, error) {

	var moduleReader, moduleError = os.Open(moduleFileName)
	if moduleError != nil {
		return nil, fmt.Errorf("error reading module file %s: %s", moduleFileName, moduleError.Error())
	}
	defer moduleReader.Close()

	var coverageReader, coverageError = os.Open(coverageFileName)
	if coverageError != nil {
		return nil, fmt.Errorf("error reading coverage file %s: %s", moduleFileName, coverageError.Error())
	}
	defer coverageReader.Close()

	var moduleName, _ = golang.GetModulePath(moduleReader)

	var coverages, parseErr = golang.MakeTableFromReader(coverageReader, len(moduleName)+1)
	if parseErr != nil {
		return nil, parseErr
	}

	return coverages, nil
}

func PrintTable(coverages []golang.CoverTable, minimunCoverage float64, writer io.Writer) error {
	var totalLines = 0
	var totalCovered = 0
	fmt.Fprintf(writer, "File\tCoverage\tUncovered lines\n")
	for _, fileCover := range coverages {
		totalLines += fileCover.TotalLines
		totalCovered += fileCover.CoveredLines
		var fileCoverage = float32(fileCover.CoveredLines) * 100.0 / float32(fileCover.TotalLines)
		var builder strings.Builder
		for _, block := range fileCover.UncoveredBlocks {
			builder.WriteString(fmt.Sprintf("%d-%d,", block.Begin, block.End))
		}
		fmt.Fprintf(writer, "%s\t%.2f%%\t%s\n", fileCover.Filename, fileCoverage, builder.String())
	}

	var totalCoverage = float64(totalCovered) * 100.0 / float64(totalLines)
	fmt.Fprintf(writer, "Total coverage\t%.2f%%\n", totalCoverage)

	if totalCoverage < minimunCoverage {
		return fmt.Errorf("fail: coverage %.2f%% < minimun coverage %.2f%%", totalCoverage, minimunCoverage)
	}

	return nil
}
