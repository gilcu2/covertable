package io

import (
	"covertable/internal/coverage"
	"os"
)

func MakeTableFromFile(fileName string) error {
	var reader, err = os.Open(fileName)
	if err != nil {
		return err
	}
	defer reader.Close()

	var r, e = coverage.MakeTableFromReader(reader)
	println(r, e)
	return nil
}
