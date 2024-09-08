package main

import (
	"flag"
	"github.com/gilcu2/covertable/internal/io"
	"log"
	"os"
)

func main() {
	os.Exit(realMain(os.Args))
}

func realMain(args []string) int {
	var coverPath string
	var exitCode int
	var modulePath string
	var minimumCoverage float64

	var commandLine = flag.NewFlagSet(args[0], flag.ExitOnError)
	commandLine.StringVar(&modulePath, "module", "go.mod", "Module file path, default go.mod")
	commandLine.Float64Var(&minimumCoverage, "minimum", 0.0, "Minimum coverage, default 0.0")
	var err = commandLine.Parse(args[1:])
	if err != nil || commandLine.NArg() < 1 {
		log.Println("Print golang table per coverPath from golang output coverPath. Usage:")
		log.Println("covertable [-module <path>] [-minimumCoverage <0.0-100.0>] <coverPath>")
		exitCode = 1
	} else {
		coverPath = commandLine.Arg(0)

		var coveTable, err = io.MakeTableFromFile(coverPath, modulePath)
		if err != nil {
			log.Println(err.Error())
			exitCode = 1
		} else {
			var err = io.PrintTable(coveTable, minimumCoverage, os.Stdout)
			if err != nil {
				log.Println(err.Error())
				exitCode = 1
			}
		}
	}
	return exitCode
}
