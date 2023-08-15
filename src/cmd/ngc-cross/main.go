package main

import (
	"bufio"
	"fmt"
	"os"
)

const usage = `
ngc-cross transpiler
usage: %s 
	[-v | --version] [-h | --help]
	[--build <ngc-path> --target <config-path>]
`

func main() {
	var buildFile *string
	var targetFile *string
	for index, item := range os.Args {
		if item == "--build" && index < len(os.Args) {
			buildFile = &os.Args[index+1]
		} else if item == "--target" && index < len(os.Args) {
			targetFile = &os.Args[index+1]
		}
	}

	// Both the build and target file must be present
	if buildFile == nil || targetFile == nil {
		fmt.Printf(usage, os.Args[0])
		return
	}

	fileHandle, err := os.Open(*buildFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(fileHandle)
	fileScanner.Split(bufio.ScanLines)

	// Eat all lines into object
	eater := NewNgcEater()
	for fileScanner.Scan() {
		eater.EatLine(fileScanner.Text())
	}

	// Create config from eater
	config := eater.AsNewConfig()

	var configData []byte
	if configData, err = config.MarshalJson(); err != nil {
		fmt.Println(err)
		return
	}

	if fileHandle, err = os.Create(*targetFile); err != nil {
		fmt.Println(err)
		return
	}

	fileWriter := bufio.NewWriter(fileHandle)
	if _, err := fileWriter.Write(configData); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("New config written to %s\n", *targetFile)
}
