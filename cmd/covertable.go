package cmd

import (
	"flag"
	"fmt"
	"github.com/gilcu2/covertable/internal/io"
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

		var coveTable, err = io.MakeTableFromFile(file)
		if err != nil {
			exitCode = 1
		} else {
			io.PrintTable(coveTable, os.Stdout)
		}
	}
	return exitCode
}
