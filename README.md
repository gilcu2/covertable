# covertable 
CLI that generate coverage table by files from the Golang coverage output. 
Include non covered lines per file.
Also check minimum coverage.

covertable [-module <path>] [-minimumCoverage <0.0-100.0>] <coverPath>

## Example

> go test -coverprofile=coverage.out ./...
> covertable coverage.out
File    Coverage        Uncovered lines
cmd/covertable.go       92.86%  10-12,
internal/golang/coverage.go     100.00%
internal/golang/go_module_file.go       100.00%
internal/io/io.go       100.00%
Total coverage  98.00%

## Install

### From Sources
go install ./cmd/covertable.go

### From repo

go install github.com/gilcu2/covertable

### From binary

Download the package from https://github.com/gilcu2/covertable/releases corresponding to your OS and architecture.
Decompress and run the executable.

### With Brew

1.  brew tap gilcu2/packages
2.  brew install covertable

## Usage

topdiffxml \<file1.xml> \<file2.xml>

## Contributions

Happy to get PR 

- Ensure 100% test coverage
- Add changes to CHANGELOG Unreleased


