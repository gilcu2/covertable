package coverage

import (
	"fmt"
	"golang.org/x/tools/cover"
	"io"
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

func MakeTableFromReader(reader io.Reader) ([]CoverTable, error) {
	var profiles, err = cover.ParseProfilesFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("error parsing coverage output: %s", err)
	}

	var coverMap = make(map[string]CoverTable)
	for _, profile := range profiles {
		var coverFile, ok = coverMap[profile.FileName]
		if !ok {
			coverFile = CoverTable{
				Filename:        profile.FileName,
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

	var coverTables = make([]CoverTable, 0, len(coverMap))

	for _, value := range coverMap {
		coverTables = append(coverTables, value)
	}
	return coverTables, nil
}
