package cmd

import (
	"gotest.tools/v3/assert"
	"testing"
)

func Test_realMain(t *testing.T) {
	// When run without parameter
	var r = realMain()

	// Then result is error
	assert.Equal(t, r, 1)
}
