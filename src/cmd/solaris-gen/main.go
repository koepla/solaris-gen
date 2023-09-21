package main

import (
	"flag"
	"fmt"
	"os"
	"solaris-gen/model"
	"solaris-gen/utility"
)

func main() {
	buildFileFlag := flag.String("build", "REQUIRED", "path to the object-config.json file")
	targetFileFlag := flag.String("target", "REQUIRED", "path to the generated header file")
	versionFlag := flag.Bool("version", false, "display the commit version")
	flag.Parse()

	if *versionFlag {
		fmt.Println("solaris-gen version: " + utility.CommitVersion())
		return
	}

	if *buildFileFlag == "REQUIRED" || *targetFileFlag == "REQUIRED" {
		flag.PrintDefaults()
		return
	}

	configData, err := os.ReadFile(*buildFileFlag)
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

	var file *os.File
	if file, err = os.Create(*targetFileFlag); err != nil {
		fmt.Println(err)
		return
	}

	writer := NewCodeWriter(file)
	if err = writer.GenerateCode(config); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Generation finished...")
}
