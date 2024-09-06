package cmd

import (
	"covertable/internal/coverage"
	"flag"
	"fmt"
	"os"
)

func main() {
	os.Exit(realMain())

}

func realMain() int {
	var file string
	var exitCode int

	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("Print coverage table per file from coverage output file. Usage:")
		fmt.Println("covertable <file>")
		exitCode = 1
	} else {
		file = flag.Arg(0)

		var err = coverage.MakeTable(file)
		if err != nil {
			exitCode = 1
		}
	}
	return exitCode
}
