package main

import (
	"fmt"
	"os"
	"solaris-gen/model"
)

const usage = `
solaris-gen code generator
usage: %s 
	[-v | --version] [-h | --help]
	[--build <config-path> --target <output-file>]
`

func main() {
	var buildFile *string
	var targetFile *string
	for index, item := range os.Args {
		if item == "--build" && index < len(os.Args) {
			buildFile = &os.Args[index+1]
		}
		if item == "--target" && index < len(os.Args) {
			targetFile = &os.Args[index+1]
		}
	}

	// BuildFile is mandatory
	if buildFile == nil || targetFile == nil {
		fmt.Printf(usage, os.Args[0])
		return
	}

	configData, err := os.ReadFile(*buildFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Retrieve config from json
	config := model.NewSolarisConfig()
	if err := config.UnmarshalJson(configData); err != nil {
		fmt.Println(err)
		return
	}

	var headerFile *os.File
	if headerFile, err = os.Create(fmt.Sprintf("%s.h", *targetFile)); err != nil {
		fmt.Println(err)
		return
	}

	var sourceFile *os.File
	if sourceFile, err = os.Create(fmt.Sprintf("%s.c", *targetFile)); err != nil {
		fmt.Println(err)
		return
	}

	writer := NewCodeWriter(headerFile, sourceFile)
	if err = writer.GenerateCode(config); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Generation finished...")
}
