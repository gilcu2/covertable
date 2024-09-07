package io

import (
	"bytes"
	"github.com/gilcu2/covertable/internal/coverage"
	"gotest.tools/v3/assert"
	"testing"
)

func TestMakeTableFromFile(t *testing.T) {
	// Given coverage output
	var filename = "testdata/output.cov"

	// And expected table
	var expected = coverage.CoverTable{
		Filename:     "github.com/gilcu2/topdiffxml/cmd/topdiffxml.go",
		TotalLines:   20,
		CoveredLines: 9,
		UncoveredBlocks: []coverage.LineBlock{
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
	var table, err = MakeTableFromFile(filename)

	// Then is expected
	assert.Equal(t, err, nil)
	assert.Equal(t, len(table), 2)
	assert.DeepEqual(t, table[0], expected)
}

func TestPrintTable(t *testing.T) {
	// Given coverage table
	var fileCoverage = coverage.CoverTable{
		Filename:     "github.com/gilcu2/topdiffxml/cmd/topdiffxml.go",
		TotalLines:   20,
		CoveredLines: 9,
		UncoveredBlocks: []coverage.LineBlock{
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
	var coverTable = []coverage.CoverTable{fileCoverage}

	// And expected out
	var expected = `File	Coverage	Uncovered lines
github.com/gilcu2/topdiffxml/cmd/topdiffxml.go	0.45	20-22,33-39,39-42,
`

	// When print table
	var buffer = new(bytes.Buffer)
	var err = PrintTable(coverTable, buffer)

	// Then is expected
	assert.Equal(t, err, nil)
	assert.Equal(t, buffer.String(), expected)
}
