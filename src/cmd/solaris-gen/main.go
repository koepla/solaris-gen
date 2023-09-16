package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	"solaris-gen/model"
)

func commitVersion() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, settingsEntry := range info.Settings {
			if settingsEntry.Key == "vcs.revision" {
				return settingsEntry.Value
			}
		}
	}
	return "unknown"
}

func main() {
	buildFileFlag := flag.String("build", "REQUIRED", "path to the object-config.json file")
	targetFileFlag := flag.String("target", "REQUIRED", "path to the generated header file")
	versionFlag := flag.Bool("version", false, "display the commit version")
	flag.Parse()

	if *versionFlag {
		fmt.Println("solaris-gen version: " + commitVersion())
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
