package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"github.com/whuang8/saigo/exercise-001-corpus/corpus"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		printUsageNote()
	}
	fileName := os.Args[1]
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		printUsageNote()
	}

	corpus.Analyze(string(fileContent))
	corpus.PrintWords()
}

func printUsageNote() {
	fmt.Printf("Please define a filename as the first argument.\nUsage: ./word_count file_name.txt\n")
	os.Exit(1)
}
