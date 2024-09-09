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

func TestMakeTableFromFileWhenWrongCoveragePath(t *testing.T) {
	// Given golang output
	var coverageFilename = "testdata/output2.cov"
	var moduleFilename = "testdata/go.mod"

	var expected = "error reading coverage file testdata/go.mod: open testdata/output2.cov: no such file or directory"

	// When create table
	var table, err = MakeTableFromFile(coverageFilename, moduleFilename)

	// Then is expected
	assert.Error(t, err, expected)
	assert.Assert(t, table == nil)
}

func TestMakeTableFromFileWhenWrongModulePath(t *testing.T) {
	// Given golang output
	var coverageFilename = "testdata/output.cov"
	var moduleFilename = "testdata/go2.mod"

	var expected = "error reading module file testdata/go2.mod: open testdata/go2.mod: no such file or directory"

	// When create table
	var table, err = MakeTableFromFile(coverageFilename, moduleFilename)

	// Then is expected
	assert.Error(t, err, expected)
	assert.Assert(t, table == nil)
}

func TestMakeTableFromFileWhenWrongCoverageData(t *testing.T) {
	// Given golang output
	var coverageFilename = "testdata/output1.cov"
	var moduleFilename = "testdata/go.mod"

	var expected = "error parsing coverage output: bad mode line: 1mode: set"
	// When create table
	var table, err = MakeTableFromFile(coverageFilename, moduleFilename)

	// Then is expected
	assert.Error(t, err, expected)
	assert.Assert(t, table == nil)
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
	var err = PrintTable(coverTable, 0.0, buffer)

	// Then is expected
	assert.Equal(t, err, nil)
	assert.Equal(t, buffer.String(), expected)
}

func TestPrintTableWhenLowerCoverage(t *testing.T) {
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
	var coverTable = []golang.CoverTable{fileCoverage1}

	// And expected output
	var expectedOutput = "File\tCoverage\tUncovered lines\n" +
		"cmd/topdiffxml.go\t45.00%\t20-22,33-39,39-42,\n" +
		"Total coverage\t45.00%\n"
	var expectedError = "fail: coverage 45.00% < minimun coverage 80.00%"

	// When print table
	var buffer = new(bytes.Buffer)
	var err = PrintTable(coverTable, 80.0, buffer)

	// Then is expected
	assert.Error(t, err, expectedError)
	assert.Equal(t, buffer.String(), expectedOutput)
}
