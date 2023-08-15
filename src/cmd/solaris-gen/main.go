package main

import (
	"fmt"
	"os"
	"solaris-gen/gen"
	"solaris-gen/model"
)

const usage = `
solaris-gen code generator
usage: %s 
	[-v | --version] [-h | --help]
	[--build <config-path>]
`

func main() {
	var buildFile *string
	for index, item := range os.Args {
		if item == "--build" && index < len(os.Args) {
			buildFile = &os.Args[index+1]
		}
	}

	// BuildFile is mandatory
	if buildFile == nil {
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

	var source, header string
	writer := gen.NewCodeWriter()
	source, header, err = writer.GenerateCode(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("header:")
	fmt.Println(header)
	fmt.Println("source:")
	fmt.Println(source)
}
