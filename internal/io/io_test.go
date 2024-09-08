package io

import (
	"bytes"
	"github.com/gilcu2/covertable/internal/golang"
	"gotest.tools/v3/assert"
	"testing"
)

func TestMakeTableFromFile(t *testing.T) {
	// Given golang output
	var coverageFilename = "testdata/output.cov"
	var moduleFilename = "testdata/go.mod"

	// And expected table
	var expected = golang.CoverTable{
		Filename:     "cmd/topdiffxml.go",
		TotalLines:   20,
		CoveredLines: 9,
		UncoveredBlocks: []golang.LineBlock{
			{
				Begin: 20,
				End:   22,
			},
			{
				Begin: 33,
				End:   39,
			},
			{
				Begin: 39,
				End:   42,
			},
		},
	}

	// When create table
	var table, err = MakeTableFromFile(coverageFilename, moduleFilename)

	// Then is expected
	assert.Equal(t, err, nil)
	assert.Equal(t, len(table), 2)
	assert.DeepEqual(t, table[0], expected)
}

func TestPrintTable(t *testing.T) {
	// Given golang table
	var fileCoverage1 = golang.CoverTable{
		Filename:     "cmd/topdiffxml.go",
		TotalLines:   20,
		CoveredLines: 9,
		UncoveredBlocks: []golang.LineBlock{
			{
				Begin: 20,
				End:   22,
			},
			{
				Begin: 33,
				End:   39,
			},
			{
				Begin: 39,
				End:   42,
			},
		},
	}
	var fileCoverage2 = golang.CoverTable{
		Filename:        "cmd/topdiffxml2.go",
		TotalLines:      20,
		CoveredLines:    20,
		UncoveredBlocks: []golang.LineBlock{},
	}
	var coverTable = []golang.CoverTable{fileCoverage1, fileCoverage2}

	// And expected out
	var expected = "File\tCoverage\tUncovered lines\n" +
		"cmd/topdiffxml.go\t45.00%\t20-22,33-39,39-42,\n" +
		"cmd/topdiffxml2.go\t100.00%\t\n" +
		"Total coverage\t72.50%\n"

	// When print table
	var buffer = new(bytes.Buffer)
	var err = PrintTable(coverTable, buffer)

	// Then is expected
	assert.Equal(t, err, nil)
	assert.Equal(t, buffer.String(), expected)
}
