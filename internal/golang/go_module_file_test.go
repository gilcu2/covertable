package golang

import (
	"gotest.tools/v3/assert"
	"strings"
	"testing"
)

func Test_GetModulePath(t *testing.T) {
	// Given module file
	var module = `module github.com/gilcu2/covertable

go 1.23.0

require (
	github.com/gilcu2/topdiffxml v0.5.6 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	golang.org/x/tools v0.24.0 // indirect
	gotest.tools/v3 v3.5.1 // indirect
)
`
	var reader = strings.NewReader(module)

	// When get path
	var path, err = GetModulePath(reader)

	// Then is expected
	assert.Equal(t, err, nil)
	assert.Equal(t, path, "github.com/gilcu2/covertable")
}
