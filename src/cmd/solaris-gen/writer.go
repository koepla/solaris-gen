package main

import (
	"bufio"
	"os"
	"solaris-gen/model"
	"text/template"
)

// CodeWriter generates code
type CodeWriter struct {
	Objects *[]model.SolarisObject
	File    *os.File
}

// NewCodeWriter creates a SolarisWriter instance
func NewCodeWriter(file *os.File) *CodeWriter {
	return &CodeWriter{
		Objects: nil,
		File:    file,
	}
}

// GenerateCode produces the source and header files that are used by solaris
func (writer *CodeWriter) GenerateCode(config *model.ObjectConfig) (err error) {
	writer.Objects = &config.Objects

	var headerTemplate *template.Template
	if headerTemplate, err = template.New(writer.File.Name()).Parse(HeaderTemplateDefinition); err != nil {
		return err
	}
	headerFileWriter := bufio.NewWriter(writer.File)
	if err = headerTemplate.Execute(headerFileWriter, writer); err != nil {
		return err
	}
	defer func() {
		_ = headerFileWriter.Flush()
	}()
	return nil
}
