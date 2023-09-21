package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"solaris-gen/utility"
)

func main() {
	buildFileFlag := flag.String("build", "REQUIRED", "path to the ngc model file")
	targetFileFlag := flag.String("target", "REQUIRED", "path to the generated object-config")
	versionFlag := flag.Bool("version", false, "display the commit version")
	flag.Parse()

	if *versionFlag {
		fmt.Println("ngc-cross version: " + utility.CommitVersion())
		return
	}

	if *buildFileFlag == "REQUIRED" || *targetFileFlag == "REQUIRED" {
		flag.PrintDefaults()
		return
	}

	fileHandle, err := os.Open(*buildFileFlag)
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

	if fileHandle, err = os.Create(*targetFileFlag); err != nil {
		fmt.Println(err)
		return
	}

	fileWriter := bufio.NewWriter(fileHandle)
	if _, err := fileWriter.Write(configData); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("New config written to %s\n", *targetFileFlag)
}
