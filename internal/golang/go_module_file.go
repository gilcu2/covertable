package golang

import (
	"bytes"
	"fmt"
	"golang.org/x/mod/modfile"
	"io"
)

func GetModulePath(reader io.Reader) (string, error) {
	var buf = new(bytes.Buffer)
	var _, _ = buf.ReadFrom(reader)
	var moduleName = modfile.ModulePath(buf.Bytes())
	if moduleName == "" {
		return moduleName, fmt.Errorf("module name not found in module file")
	}
	return moduleName, nil
}
