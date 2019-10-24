package main

import (
	"math/rand"
	"os"
	"teamCMP/fileReaders"
	"time"
)

/// Main program
/// Reads a source in order to read all files in feed-exports and then "save the exports"
func main() {
	// Initializing seed for mockUp
	rand.Seed(time.Now().Unix())

	// Read arguments
	argsWithoutProg := os.Args[1:]

	importAllSourceFiles := false
	source := ""

	for _, argument := range argsWithoutProg {
		if argument == "-a" {
			importAllSourceFiles = true
		} else {
			source = argument
		}
	}

	if len(argsWithoutProg) != 1 {
		println("Wrong arguments.")
		println("Options:")
		println("\t -a : import all source files.")
		println("\t <source_name> : import source files from given source.")
		return
	}

	var feedImporter fileReaders.FeedImporter

	err := feedImporter.GetVideosData("", source, "", "", importAllSourceFiles)
	if err != nil {
		println(err.Error())
		return
	}
}
