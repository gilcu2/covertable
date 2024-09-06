package coverage

import (
	"gotest.tools/v3/assert"
	"strings"
	"testing"
)

func TestMakeTable(t *testing.T) {
	// Given coverage output
	var coverage = `mode: set
github.com/gilcu2/topdiffxml/cmd/topdiffxml.go:20.13,22.2 1 0
github.com/gilcu2/topdiffxml/cmd/topdiffxml.go:24.21,29.21 4 1
github.com/gilcu2/topdiffxml/cmd/topdiffxml.go:29.21,33.3 3 1
github.com/gilcu2/topdiffxml/cmd/topdiffxml.go:33.8,39.17 4 0
github.com/gilcu2/topdiffxml/cmd/topdiffxml.go:39.17,42.4 2 0
github.com/gilcu2/topdiffxml/cmd/topdiffxml.go:45.2,45.17 1 1
`
	var coverageReader = strings.NewReader(coverage)

	// When create table
	var table, err = MakeTableFromReader(coverageReader)

	// Then is expected
	assert.Equal(t, err, nil)
	assert.Equal(t, len(table), 2)
}
