package main

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
	var coverPath string
	var exitCode int
	var modulePath string

	flag.StringVar(&modulePath, "module", "go.mod", "Module file path, default go.mod")
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("Print golang table per coverPath from golang output coverPath. Usage:")
		fmt.Println("covertable [-module <path>] <coverPath>")
		exitCode = 1
	} else {
		coverPath = flag.Arg(0)

		var coveTable, err = io.MakeTableFromFile(coverPath, modulePath)
		if err != nil {
			exitCode = 1
		} else {
			io.PrintTable(coveTable, os.Stdout)
		}
	}
	return exitCode
}
