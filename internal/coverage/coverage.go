package coverage

import (
	"golang.org/x/tools/cover"
	"io"
)

type LineBlock struct {
	begin int
	end   int
}

type CoverTable struct {
	filename        string
	totalLines      int
	coveredLines    int
	uncoveredBlocks []LineBlock
}

func MakeTableFromReader(reader io.Reader) ([]CoverTable, error) {
	var profiles, err = cover.ParseProfilesFromReader(reader)
	if err != nil {
		return nil, err
	}

	var coverMap = make(map[string]CoverTable)
	for _, profile := range profiles {
		var coverFile, ok = coverMap[profile.FileName]
		if !ok {
			coverFile = CoverTable{
				filename:        profile.FileName,
				totalLines:      0,
				coveredLines:    0,
				uncoveredBlocks: []LineBlock{},
			}
		}
		for _, block := range profile.Blocks {
			var blockLines = block.EndLine - block.StartLine
			coverFile.totalLines += blockLines
			if block.Count == 0 {
				coverFile.uncoveredBlocks = append(coverFile.uncoveredBlocks,
					LineBlock{begin: block.StartLine, end: block.EndLine},
				)
			} else {
				coverFile.coveredLines += blockLines
			}

		}
	}

	var coverTables = make([]CoverTable, 0, len(coverMap))

	for _, value := range coverMap {
		coverTables = append(coverTables, value)
	}
	return coverTables, nil
}
