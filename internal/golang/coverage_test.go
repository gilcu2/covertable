package golang

import (
	"gotest.tools/v3/assert"
	"strings"
	"testing"
)

func TestMakeTable(t *testing.T) {
	// Given golang output
	var coverage = `mode: set
github.com/gilcu2/topdiffxml/cmd/topdiffxml.go:20.13,22.2 1 0
github.com/gilcu2/topdiffxml/cmd/topdiffxml.go:24.21,29.21 4 1
github.com/gilcu2/topdiffxml/cmd/topdiffxml.go:29.21,33.3 3 1
github.com/gilcu2/topdiffxml/cmd/topdiffxml.go:33.8,39.17 4 0
github.com/gilcu2/topdiffxml/cmd/topdiffxml.go:39.17,42.4 2 0
github.com/gilcu2/topdiffxml/cmd/topdiffxml.go:45.2,45.17 1 1
`
	var coverageReader = strings.NewReader(coverage)

	// And expected table
	var expected = CoverTable{
		Filename:     "cmd/topdiffxml.go",
		TotalLines:   20,
		CoveredLines: 9,
		UncoveredBlocks: []LineBlock{
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
	var table, err = MakeTableFromReader(coverageReader, len("github.com/gilcu2/topdiffxml/"))

	// Then is expected
	assert.Equal(t, err, nil)
	assert.Equal(t, len(table), 1)
	assert.DeepEqual(t, table[0], expected)
}

func TestMakeTableWhenError(t *testing.T) {
	// Given golang output
	var coverage = `1mode: set
github.com/gilcu2/topdiffxml/cmd/topdiffxml.go:20.13,22.2 1 0
`
	var coverageReader = strings.NewReader(coverage)

	var expected = "error parsing coverage output: bad mode line: 1mode: set"

	// When create table
	var table, err = MakeTableFromReader(coverageReader, len("github.com/gilcu2/topdiffxml/"))

	// Then is expected
	assert.Error(t, err, expected)
	assert.Assert(t, table == nil)
}
