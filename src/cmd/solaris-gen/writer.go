package main

import (
	"bufio"
	"github.com/agnivade/levenshtein"
	"os"
	"solaris-gen/model"
	"text/template"
)

// EnumEntry represents an enum entry of generated code
type EnumEntry struct {
	Text       string
	Definition string
}

// NewEnumEntry creates a new EnumEntry instance
func NewEnumEntry(text string, definition string) EnumEntry {
	return EnumEntry{
		Text:       text,
		Definition: definition,
	}
}

// Enum represents an enum type of generated code
type Enum struct {
	Entries []EnumEntry
}

// NewEnum creates a new SolarisEnum instance
func NewEnum() Enum {
	return Enum{
		Entries: []EnumEntry{},
	}
}

// Append tries to append the enum entry to the enum
func (e *Enum) Append(text string, definition string) {
	var closeCandidate *EnumEntry
	for _, enumEntry := range e.Entries {
		if levenshtein.ComputeDistance(enumEntry.Text, text) < 2 {
			closeCandidate = &enumEntry
			break
		}
	}
	if closeCandidate != nil {
		return
	}
	e.Entries = append(e.Entries, NewEnumEntry(text, definition))
}

// CodeWriter generates code
type CodeWriter struct {
	ObjectTypes    Enum
	Catalogs       Enum
	Constellations Enum
	Objects        *[]model.SolarisObject
	ObjectCount    int
	HeaderFile     *os.File
	SourceFile     *os.File
}

// NewCodeWriter creates a SolarisWriter instance
func NewCodeWriter(headerFile *os.File, sourceFile *os.File) *CodeWriter {
	return &CodeWriter{
		Objects:        nil,
		ObjectTypes:    NewEnum(),
		ObjectCount:    0,
		Catalogs:       NewEnum(),
		Constellations: NewEnum(),
		HeaderFile:     headerFile,
		SourceFile:     sourceFile,
	}
}

// buildConfig builds the writer config to the enum symbols
func (writer *CodeWriter) buildConfig(config *model.ObjectConfig) {
	for _, entry := range config.Objects {
		writer.ObjectTypes.Append(entry.ObjectTypeText(), entry.ObjectTypeDefinition())
		writer.Catalogs.Append(entry.Designation.CatalogText(), entry.Designation.CatalogDefinition())
		writer.Constellations.Append(entry.ConstellationText(), entry.ConstellationDefinition())
	}
	writer.Objects = &config.Objects
	writer.ObjectCount = len(config.Objects)
}

// GenerateCode produces the source and header files that are used by solaris
func (writer *CodeWriter) GenerateCode(config *model.ObjectConfig) (err error) {
	writer.buildConfig(config)

	var headerTemplate *template.Template
	if headerTemplate, err = template.New(writer.HeaderFile.Name()).Parse(HeaderTemplateDefinition); err != nil {
		return err
	}

	var sourceTemplate *template.Template
	if sourceTemplate, err = template.New(writer.SourceFile.Name()).Parse(SourceTemplateDefinition); err != nil {
		return err
	}

	headerFileWriter := bufio.NewWriter(writer.HeaderFile)
	sourceFileWriter := bufio.NewWriter(writer.SourceFile)

	if err = headerTemplate.Execute(headerFileWriter, writer); err != nil {
		return err
	}
	if err = sourceTemplate.Execute(sourceFileWriter, writer); err != nil {
		return err
	}

	defer func() {
		_ = headerFileWriter.Flush()
	}()
	defer func() {
		_ = sourceFileWriter.Flush()
	}()
	return nil
}
