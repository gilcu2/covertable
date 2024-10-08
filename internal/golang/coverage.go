package golang

import (
	"fmt"
	"golang.org/x/tools/cover"
	"io"
	"sort"
)

type LineBlock struct {
	Begin int
	End   int
}

type CoverTable struct {
	Filename        string
	TotalLines      int
	CoveredLines    int
	UncoveredBlocks []LineBlock
}

func MakeTableFromReader(reader io.Reader, modulePathLen int) ([]CoverTable, error) {
	var profiles, err = cover.ParseProfilesFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("error parsing coverage output: %s", err)
	}

	var coverMap = make(map[string]CoverTable)
	for _, profile := range profiles {
		var coverFile, ok = coverMap[profile.FileName]
		if !ok {
			coverFile = CoverTable{
				Filename:        profile.FileName[modulePathLen:],
				TotalLines:      0,
				CoveredLines:    0,
				UncoveredBlocks: []LineBlock{},
			}
		}
		for _, block := range profile.Blocks {
			var blockLines = block.EndLine - block.StartLine
			coverFile.TotalLines += blockLines
			if block.Count == 0 {
				coverFile.UncoveredBlocks = append(coverFile.UncoveredBlocks,
					LineBlock{Begin: block.StartLine, End: block.EndLine},
				)
			} else {
				coverFile.CoveredLines += blockLines
			}

		}
		coverMap[profile.FileName] = coverFile
	}

	coverTables := getValuesSortedByKeys(coverMap)
	return coverTables, nil
}

func getValuesSortedByKeys(coverMap map[string]CoverTable) []CoverTable {
	var files = make([]string, 0)
	for file := range coverMap {
		files = append(files, file)
	}
	sort.Strings(files)

	var coverTables = make([]CoverTable, 0, len(coverMap))
	for _, file := range files {
		coverTables = append(coverTables, coverMap[file])
	}
	return coverTables
}
