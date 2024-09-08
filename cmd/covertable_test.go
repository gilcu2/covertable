package main

import (
	"gotest.tools/v3/assert"
	"testing"
)

func Test_realMainNoArguments(t *testing.T) {
	// When run without parameter
	var r = realMain([]string{"covertable"})

	// Then result is error
	assert.Equal(t, r, 1)
}

func Test_realMainOK(t *testing.T) {
	// When run without parameter
	var r = realMain([]string{
		"covertable",
		"-module",
		"../internal/io/testdata/go.mod",
		"../internal/io/testdata/output.cov",
	})

	// Then result is error
	assert.Equal(t, r, 0)
}

func Test_realMainWhenLowCoverage(t *testing.T) {
	// When run without parameter
	var r = realMain([]string{
		"covertable",
		"-module", "../internal/io/testdata/go.mod",
		"-minimum", "98",
		"../internal/io/testdata/output.cov",
	})

	// Then result is error
	assert.Equal(t, r, 1)
}

func Test_realMainWrongCoverageData(t *testing.T) {
	// When run without parameter
	var r = realMain([]string{
		"covertable",
		"-module",
		"internal/io/testdata/go.mod",
		"internal/io/testdata/output1.cov",
	})

	// Then result is error
	assert.Equal(t, r, 1)
}
