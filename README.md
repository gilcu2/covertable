# covertable 
CLI that generate coverage table by files from the Golang coverage output. 
The output includes non-covered lines per file.
Also check minimum coverage.

covertable [-module <path>] [-minimumCoverage <0.0-100.0>] <coverPath>

## Example

```shell
> go test -coverprofile=coverage.out ./...
> covertable coverage.out
File    Coverage        Uncovered lines
cmd/covertable.go       92.86%  10-12,
internal/golang/coverage.go     100.00%
internal/golang/go_module_file.go       100.00%
internal/io/io.go       100.00%
Total coverage  98.00%
```

## Install

### From Sources
go install ./cmd/covertable.go

### From repo

go get github.com/gilcu2/covertable

### From binary

Download the package from https://github.com/gilcu2/covertable/releases corresponding to your OS and architecture.
Decompress and run the executable.

### With Brew

1.  brew tap gilcu2/packages
2.  brew install covertable

## Usage

covertable \[-module <path>] \[-minimumCoverage <0.0-100.0>] \<coverage file>

## Contributions

Contributions are welcome

- Ensure 100% test coverage
- Add changes to CHANGELOG Unreleased


